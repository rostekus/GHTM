package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rostekus/ghtm/internal/app/service/ports/input"
)

type Store interface {
	input.Repository
}
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
