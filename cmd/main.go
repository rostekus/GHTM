package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rostekus/ghtm/config"
	"github.com/rostekus/ghtm/internal/app/service"
	handler "github.com/rostekus/ghtm/internal/handler/user"
	"github.com/rostekus/ghtm/internal/repository"
	"github.com/rostekus/ghtm/internal/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	if cfg.Environment == "prod" {
		// log.Logger = zerolog.New(file).With().Timestamp().Logger()
	} else {
		// fileLogger := zerolog.New(file).With().Timestamp().Logger()
		stdoutLogger := zerolog.ConsoleWriter{Out: os.Stderr}
		// multi := zerolog.MultiLevelWriter(stdoutLogger, fileLogger)
		log.Logger = zerolog.New(stdoutLogger).With().Timestamp().Logger()

	}
	e := echo.New()
	e.Static("/static", "static")

	db := repository.NewInMemoryUserRepository()
	s, err := service.NewApp(db)
	userHandler := handler.Handler{Service: s}
	srv, err := server.NewServer(userHandler, cfg)
	srv.StartAndListen()
}
