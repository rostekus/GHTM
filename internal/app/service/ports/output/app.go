package output

import (
	"github.com/rostekus/ghtm/internal/app/service/dto"
)

type App interface {
	CreateUser(dto.UserDTO) (dto.UserDTO, error)
	GetUser(dto.UserDTO) (dto.UserDTO, error)
}
