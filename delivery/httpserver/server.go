package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.sazito.com/sazito/event_publisher/config"
)

type HTTPServer struct {
	Port int `env:"SMD_SERVE_PORT"`
}

type Server struct {
	config	config.Config
}

func New(cfg config.Config) *Server {
	return &Server{
		config: cfg,
	}
}

func (s Server) Serve() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Start server
	err := e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port))
	e.Logger.Fatal(err)
}

