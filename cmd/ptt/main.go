package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/MunifTanjim/go-ptt"
	"github.com/MunifTanjim/go-ptt/cmd/ptt/server"
	"github.com/urfave/cli/v3"
)

func Int32SliceFromInt(input []int) []int32 {
	output := make([]int32, len(input))
	for i := range input {
		output[i] = int32(input[i])
	}
	return output
}

func main() {
	cmd := &cli.Command{
		Name:  "ptt",
		Usage: "parse torrent title",
		Commands: []*cli.Command{
			{
				Name: "parse",
				Flags: []cli.Flag{
					&cli.BoolWithInverseFlag{
						Name:  "normalize",
						Value: true,
					},
					&cli.BoolFlag{
						Name: "pretty",
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					torrent_title := cmd.Args().First()
					r := ptt.Parse(torrent_title)
					if err := r.Error(); err != nil {
						return err
					}
					if cmd.Bool("normalize") {
						r = r.Normalize()
					}
					var blob []byte
					var err error
					if cmd.Bool("pretty") {
						blob, err = json.MarshalIndent(&r, "", "  ")
					} else {
						blob, err = json.Marshal(&r)
					}
					if err != nil {
						return err
					}
					fmt.Printf("%v\n", string(blob))
					return nil
				},
			},
			{
				Name: "server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "network",
						Usage: "network protocol",
						Value: "unix",
						Validator: func(network string) error {
							switch network {
							case "unix", "tcp":
								return nil
							default:
								return fmt.Errorf("unsupported network: %s", network)
							}
						},
					},
					&cli.StringFlag{
						Name:  "address",
						Usage: "network address",
						Value: "ptt.sock",
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) (err error) {
					network := cmd.String("network")
					address := cmd.String("address")
					switch network {
					case "unix":
						address, err = filepath.Abs(address)
						if err != nil {
							return err
						}
					}

					if err := server.Start(&server.PTTServerConfig{
						Network: network,
						Address: address,
					}); err != nil {
						log.Fatalf("server failed to start: %v", err)
					}

					log.Println("server stopped")
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
