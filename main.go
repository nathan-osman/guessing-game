//go:generate go run ui/generate.go

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nathan-osman/guessing-game/server"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

func main() {
	app := cli.NewApp()
	app.Name = "i5"
	app.Usage = "reverse proxy for Docker containers"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "addr",
			Value:  ":http",
			EnvVar: "ADDR",
			Usage:  "HTTP address to listen on",
		},
		cli.BoolFlag{
			Name:   "debug",
			EnvVar: "DEBUG",
			Usage:  "enable debug mode",
		},
	}
	app.Action = func(c *cli.Context) error {

		// Initialize zap
		var cfg zap.Config
		if c.Bool("debug") {
			cfg = zap.NewDevelopmentConfig()
		} else {
			cfg = zap.NewProductionConfig()
		}
		logger, err := cfg.Build()
		if err != nil {
			return err
		}

		// Initialize the server
		s, err := server.New(&server.Config{
			Addr:   c.String("addr"),
			Logger: logger,
		})
		if err != nil {
			return err
		}
		defer s.Close()

		// Wait for SIGINT or SIGTERM
		sigChan := make(chan os.Signal)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		return nil
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "fatal: %s\n", err.Error())
	}
}
