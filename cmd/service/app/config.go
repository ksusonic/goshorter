package app

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type config struct {
	DatabaseDSN    string `env:"DATABASE_DSN"`
	Port           int    `env:"PORT" envDefault:"8080"`
	ShortURLPrefix string `env:"SHORT_URL_PREFIX" envDefault:"http://localhost:8080/"`
}

func newConfig() config {
	var cfg config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("parse env: %+v", err)
	}

	return cfg
}
