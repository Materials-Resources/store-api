package config

import (
	"fmt"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Server    Server    `koanf:"server"`
	Services  Services  `konanf:"services"`
	Session   Session   `konanf:"session"`
	Telemetry Telemetry `konanf:"telemetry"`
	Mailer    Mailer    `konanf:"mailer"`
	Zitadel   Zitadel   `konanf:"zitadel"`
	Env       string
}

func newConfig() *Config {
	cfg := &Config{}
	cfg.Server.SetDefaults()
	cfg.Services.SetDefaults()
	cfg.Mailer.SetDefaults()
	cfg.Telemetry.SetDefaults()
	return cfg
}

func ReadConfig() (*Config, error) {
	cfg := newConfig()
	var k = koanf.New(".")

	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	if err := k.Unmarshal("", cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return cfg, nil
}
