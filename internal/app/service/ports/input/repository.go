package input

import (
	"github.com/rostekus/ghtm/internal/app/domain"
)

type Repository interface {
	CreateUser(domain.User) (domain.User, error)
	GetUser(domain.User) (domain.User, error)
}
