package config

import (
	"gitlab.sazito.com/sazito/event_publisher/adapter/redisqueue"
	"gitlab.sazito.com/sazito/event_publisher/pkg/postgresql"
)

type HTTPServer struct {
	Port int `env:"SMD_SERVE_PORT"`
}

type Config struct {
	HTTPServer HTTPServer `env:"SMD_SERVE_HTTP"`
	Redis	redisqueue.Config
	Postgres postgresql.Config
}
