package app

import (
	"github.com/rs/zerolog"
	"os"
)

func newLogger() *zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stderr}

	lg := zerolog.New(output).With().Timestamp().Logger()
	return &lg
}
