package config

import (
	"fmt"
	"github.com/caarlos0/env/v8"
	"gitlab.sazito.com/sazito/event_publisher/pkg/richerror"
	"log"
)

func Load() (Config, error) {
	const op = "config.Load"
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Println("Error parsing environment variables:", err)
		return cfg, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	fmt.Printf("Load config: %+v\n", cfg)

	return cfg, nil
}
