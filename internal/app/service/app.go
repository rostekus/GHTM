package service

import (
	"github.com/rostekus/ghtm/internal/app/service/dto"
	"github.com/rostekus/ghtm/internal/app/service/mapper"
	"github.com/rostekus/ghtm/internal/app/service/ports/input"
)

type app struct {
	db input.Repository
}

func NewApp(db input.Repository) (*app, error) {
	return &app{
		db: db,
	}, nil
}

func (a *app) CreateUser(u dto.UserDTO) (dto.UserDTO, error) {
	user := mapper.UserDTOtoUser(u)
	userDTO, err := a.db.CreateUser(user)
	return mapper.UsertoUserDTO(userDTO), err
}

func (a *app) GetUser(u dto.UserDTO) (dto.UserDTO, error) {
	user := mapper.UserDTOtoUser(u)
	userDTO, err := a.db.GetUser(user)
	return mapper.UsertoUserDTO(userDTO), err
}
