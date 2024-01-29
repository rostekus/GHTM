package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/rostekus/ghtm/config"
	handler "github.com/rostekus/ghtm/internal/handler/user"
	"github.com/rs/zerolog/log"
)

type Server struct {
	userH handler.Handler
	e     *echo.Echo
	cfg   config.Config
}

func NewServer(uH handler.Handler, cfg config.Config) (*Server, error) {
	return &Server{
		userH: uH,
		cfg:   cfg,
	}, nil
}

func (server *Server) StartAndListen() error {
	server.setupRouter()
	log.Info().Msg("listening on port")
	return server.e.Start(fmt.Sprintf(":%d", server.cfg.HttpPort))
}

func (server *Server) setupRouter() {
	server.e = echo.New()
	server.e.Static("/static", "static")
	server.e.POST("/api/v1/register", server.userH.RegisterUserApi)
	server.e.POST("/api/v1/login", server.userH.LoginUserApi)
}
