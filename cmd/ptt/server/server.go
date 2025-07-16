//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/ptt.proto

package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/MunifTanjim/go-ptt"
	"github.com/MunifTanjim/go-ptt/cmd/ptt/server/proto"
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

func (s *server) Parse(ctx context.Context, req *proto.ParseRequest) (*proto.ParseResponse, error) {
	results := []*proto.ParseResponse_Result{}
	for _, torrent_title := range req.TorrentTitles {
		r := ptt.Parse(torrent_title)
		if req.Normalize {
			r = r.Normalize()
		}
		if err := r.Error(); err != nil {
			results = append(results, &proto.ParseResponse_Result{
				Err: err.Error(),
			})
			continue
		}
		results = append(results, &proto.ParseResponse_Result{
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
		})
	}
	return &proto.ParseResponse{
		Results: results,
	}, nil
}

type PTTServer struct {
	socket     string
	listener   net.Listener
	grpcServer *grpc.Server
	onSignal   func(s os.Signal)
	signalChan chan os.Signal
}

type PTTServerConfig struct {
	SocketPath string
	OnSignal   func(s os.Signal)
}

func NewPTTServer(conf *PTTServerConfig) *PTTServer {
	pttServer := PTTServer{
		socket:     conf.SocketPath,
		grpcServer: grpc.NewServer(),
		onSignal:   conf.OnSignal,
		signalChan: make(chan os.Signal, 1),
	}

	proto.RegisterServiceServer(pttServer.grpcServer, &server{})

	return &pttServer
}

func (s *PTTServer) Listen() error {
	if err := os.RemoveAll(s.socket); err != nil {
		return fmt.Errorf("failed to remove socket: %v", err)
	}

	listener, err := net.Listen("unix", s.socket)
	if err != nil {
		return fmt.Errorf("failed to listen on unix socket: %v", err)
	}
	s.listener = listener

	if err := os.Chmod(s.socket, 0666); err != nil {
		return fmt.Errorf("failed to set socket permissions: %v", err)
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
		os.Remove(s.socket)
	}()

	return s.grpcServer.Serve(s.listener)
}

func Start(conf *PTTServerConfig) error {
	s := NewPTTServer(&PTTServerConfig{
		SocketPath: conf.SocketPath,
		OnSignal: func(s os.Signal) {
			log.Printf("received signal: %s\n", s.String())
		},
	})

	err := s.Listen()
	if err != nil {
		return err
	}
	defer s.Close()

	log.Printf("grpc server listening on unix socket: %s", conf.SocketPath)

	return s.Serve()
}
