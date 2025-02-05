package app

import "github.com/materials-resources/customer-api/config"

type App struct {
	Otel *Otel
}

func New() (*App, error) {
	otel, err := NewOtel(config.Config{
		Env:  "dev",
		Otel: struct{ ServiceName string }{ServiceName: "customer-api"},
	})
	if err != nil {
		return nil, err
	}
	return &App{
		Otel: otel,
	}, nil
}
