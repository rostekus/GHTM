package input

import (
	"context"

	"github.com/rostekus/ghtm/internal/app/domain"
)

type Repository interface {
	CreateUser(context.Context, domain.User) (domain.User, error)
	GetUser(context.Context, domain.User) (domain.User, error)
}
