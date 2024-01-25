package input

import (
	"github.com/google/uuid"
	"github.com/rostekus/ghtm/internal/app/domain"
)

type Repository interface {
	CreateUser(domain.User) (domain.User, error)
	GetUserById(uuid.UUID) (domain.User, error)
}
