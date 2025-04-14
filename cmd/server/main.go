package main

import (
	"fmt"
	"github.com/materials-resources/store-api/app"
	"github.com/materials-resources/store-api/config"
	"github.com/materials-resources/store-api/internal/mailer"
	"github.com/materials-resources/store-api/internal/oas"
	"github.com/materials-resources/store-api/internal/service"
	"github.com/materials-resources/store-api/internal/session"
	"log"
	"net/http"
)

func main() {
	sm := session.NewManager("https://auth.materials-resources.com/oauth/v2/keys")

	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	a, err := app.New(*cfg)
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
}
