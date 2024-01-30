package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rostekus/ghtm/config"
	db "github.com/rostekus/ghtm/db/sqlc"
	"github.com/rostekus/ghtm/internal/app/service"
	handler "github.com/rostekus/ghtm/internal/handler/user"
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
	db.RunDBMigration(cfg.MigrationURL, cfg.DBSource)
	connPool, err := pgxpool.New(context.Background(), cfg.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	store := db.NewStore(connPool)
	s, err := service.NewApp(store)
	userHandler := handler.Handler{Service: s}
	srv, err := server.NewServer(userHandler, cfg)
	srv.StartAndListen()
}
