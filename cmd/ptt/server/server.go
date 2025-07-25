//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/ptt.proto

package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"slices"
	"syscall"

	"github.com/MunifTanjim/go-ptt"
	"github.com/MunifTanjim/go-ptt/cmd/ptt/server/proto"
	"github.com/alitto/pond/v2"
	"google.golang.org/grpc"
)

func Int32SliceFromInt(input []int) []int32 {
	output := make([]int32, len(input))
	for i := range input {
		output[i] = int32(input[i])
	}
	return output
}

type server struct {
	proto.UnimplementedServiceServer
}

var pool = pond.NewPool(20)

func (s *server) Parse(ctx context.Context, req *proto.ParseRequest) (*proto.ParseResponse, error) {
	count := len(req.TorrentTitles)
	results := make([]*proto.ParseResponse_Result, count)

	chunk_size := min(count, 500)
	if chunk_size == count && count > 200 {
		chunk_size = 100
	}
	chunk_idx := -1
	g := pool.NewGroup()
	for torrent_titles := range slices.Chunk(req.TorrentTitles, chunk_size) {
		chunk_idx++
		cidx := chunk_idx
		g.Submit(func() {
			for idx, torrent_title := range torrent_titles {
				r := ptt.Parse(torrent_title)
				if req.Normalize {
					r = r.Normalize()
				}
				if err := r.Error(); err != nil {
					results[cidx*chunk_size+idx] = &proto.ParseResponse_Result{
						Err: err.Error(),
					}
					continue
				}
				results[cidx*chunk_size+idx] = &proto.ParseResponse_Result{
					Audio:        r.Audio,
					BitDepth:     r.BitDepth,
					Channels:     r.Channels,
					Codec:        r.Codec,
					Commentary:   r.Commentary,
					Complete:     r.Complete,
					Container:    r.Container,
					Convert:      r.Convert,
					Date:         r.Date,
					Documentary:  r.Documentary,
					Dubbed:       r.Dubbed,
					Edition:      r.Edition,
					EpisodeCode:  r.EpisodeCode,
					Episodes:     Int32SliceFromInt(r.Episodes),
					Extended:     r.Extended,
					Extension:    r.Extension,
					Group:        r.Group,
					Hdr:          r.HDR,
					Hardcoded:    r.Hardcoded,
					Languages:    r.Languages,
					Network:      r.Network,
					Proper:       r.Proper,
					Quality:      r.Quality,
					ReleaseTypes: r.ReleaseTypes,
					Region:       r.Region,
					Remastered:   r.Remastered,
					Repack:       r.Repack,
					Resolution:   r.Resolution,
					Retail:       r.Retail,
					Seasons:      Int32SliceFromInt(r.Seasons),
					Site:         r.Site,
					Size:         r.Size,
					Subbed:       r.Subbed,
					ThreeD:       r.ThreeD,
					Title:        r.Title,
					Uncensored:   r.Uncensored,
					Unrated:      r.Unrated,
					Upscaled:     r.Upscaled,
					Volumes:      Int32SliceFromInt(r.Volumes),
					Year:         r.Year,
				}
			}
		})
	}
	g.Wait()
	return &proto.ParseResponse{
		Results: results,
	}, nil
}

func (s *server) Ping(ctx context.Context, req *proto.PingRequest) (*proto.PingResponse, error) {
	return &proto.PingResponse{
		Message: req.Message,
	}, nil
}

type PTTServer struct {
	network    string
	address    string
	listener   net.Listener
	grpcServer *grpc.Server
	onSignal   func(s os.Signal)
	signalChan chan os.Signal
}

type PTTServerConfig struct {
	Network  string
	Address  string
	OnSignal func(s os.Signal)
}

func NewPTTServer(conf *PTTServerConfig) *PTTServer {
	switch conf.Network {
	case "unix", "tcp":
	default:
		log.Fatalf("unsupported network: %s", conf.Network)
	}

	pttServer := PTTServer{
		network:    conf.Network,
		address:    conf.Address,
		grpcServer: grpc.NewServer(),
		onSignal:   conf.OnSignal,
		signalChan: make(chan os.Signal, 1),
	}

	proto.RegisterServiceServer(pttServer.grpcServer, &server{})

	return &pttServer
}

func (s *PTTServer) Listen() error {
	if s.network == "unix" {
		if err := os.RemoveAll(s.address); err != nil {
			return fmt.Errorf("failed to remove unix socket: %v", err)
		}
	}

	listener, err := net.Listen(s.network, s.address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	s.listener = listener

	if s.network == "unix" {
		if err := os.Chmod(s.address, 0666); err != nil {
			return fmt.Errorf("failed to set unix socket permissions: %v", err)
		}
	}

	return nil
}

func (s *PTTServer) Close() error {
	err := s.listener.Close()
	return err
}

func (s *PTTServer) Serve() error {
	signal.Notify(s.signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-s.signalChan
		s.onSignal(sig)
		s.grpcServer.GracefulStop()
		if s.network == "unix" {
			os.Remove(s.address)
		}
	}()

	return s.grpcServer.Serve(s.listener)
}

func Start(conf *PTTServerConfig) error {
	s := NewPTTServer(&PTTServerConfig{
		Network: conf.Network,
		Address: conf.Address,
		OnSignal: func(s os.Signal) {
			log.Printf("received signal: %s\n", s.String())
		},
	})

	err := s.Listen()
	if err != nil {
		return err
	}
	defer s.Close()

	log.Printf("grpc server listening on: %s://%s", conf.Network, conf.Address)

	return s.Serve()
}
