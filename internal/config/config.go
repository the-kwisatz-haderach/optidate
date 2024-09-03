package config

import (
	"github.com/caarlos0/env"
	"github.com/rs/zerolog/log"
)

type config struct {
	Port int `env:"PORT" envDefault:"8000"`
}

func ParseConfig() *config {
	var cfg config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse env variables")
	}
	return &cfg
}
