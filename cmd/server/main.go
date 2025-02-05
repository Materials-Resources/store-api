package main

import (
	"github.com/materials-resources/customer-api/app"
	"github.com/materials-resources/customer-api/internal/oas"
	"github.com/materials-resources/customer-api/internal/service"
	"github.com/materials-resources/customer-api/internal/session"
	"log"
	"net/http"
)

func main() {
	sm := session.NewManager("https://auth.materials-resources.com/oauth/v2/keys")

	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	h := NewHandler(service.NewService(a), sm)

	srv, err := oas.NewServer(h, NewSecurityHandler(sm), oas.WithMeterProvider(a.Otel.GetMeterProvider()), oas.WithTracerProvider(a.Otel.GetTracerProvider()))
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
