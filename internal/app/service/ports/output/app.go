package output

import (
	"context"

	"github.com/rostekus/ghtm/internal/app/service/dto"
)

type App interface {
	CreateUser(context.Context, dto.UserDTO) (dto.UserDTO, error)
	GetUser(context.Context, dto.UserDTO) (dto.UserDTO, error)
}
