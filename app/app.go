package app

import (
	"github.com/materials-resources/store-api/config"
	"github.com/rs/zerolog"
)

type App struct {
	Config *config.Config
	Otel   *Otel
	Logger *zerolog.Logger
}

func New(c config.Config) (*App, error) {
	otel, err := NewOtel(c)
	if err != nil {
		return nil, err
	}
	return &App{
		Config: &c,
		Otel:   otel,
		Logger: newLogger(),
	}, nil
}
