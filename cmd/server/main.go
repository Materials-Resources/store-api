package main

import (
	"context"
	"fmt"
	"github.com/materials-resources/store-api/app"
	"github.com/materials-resources/store-api/config"
	"github.com/materials-resources/store-api/internal/mailer"
	"github.com/materials-resources/store-api/internal/oas"
	"github.com/materials-resources/store-api/internal/service"
	"github.com/materials-resources/store-api/internal/session"
	"github.com/urfave/cli/v3"
	"log"
	"net/http"
	"os"
)

func main() {
	var configFlag string

	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Usage:       "Path to the config file",
				Value:       "config.yaml",
				Destination: &configFlag,
			},
		},
		Action: func(ctx context.Context, command *cli.Command) error {
			cfg, err := config.ReadConfig(configFlag)
			if err != nil {
				log.Fatal(err)
			}
			a, err := app.New(*cfg)
			sm := session.NewManager(a)
			if err != nil {
				log.Fatal(err)
			}

			h := NewHandler(a, service.NewService(a), sm, mailer.New(a))

			srv, err := oas.NewServer(h, NewSecurityHandler(sm), oas.WithMeterProvider(a.Otel.GetMeterProvider()), oas.WithTracerProvider(a.Otel.GetTracerProvider()))
			if err != nil {
				a.Logger.Fatal().Err(err).Msg("could not create server handler")
			}

			a.Logger.Info().Msgf("starting server on port %d", a.Config.Server.Port)
			if err := http.ListenAndServe(fmt.Sprintf(":%d", a.Config.Server.Port), srv); err != nil {
				a.Logger.Fatal().Err(err).Msg("could not start server")
			}
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
