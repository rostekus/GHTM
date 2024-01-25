package service

import (
	"context"

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

func (a *app) CreateUser(c context.Context, u dto.UserDTO) (dto.UserDTO, error) {
	user := mapper.UserDTOtoUser(u)
	userDTO, err := a.db.CreateUser(c, user)
	return mapper.UsertoUserDTO(userDTO), err
}

func (a *app) GetUser(c context.Context, u dto.UserDTO) (dto.UserDTO, error) {
	user := mapper.UserDTOtoUser(u)
	userDTO, err := a.db.GetUser(c, user)
	return mapper.UsertoUserDTO(userDTO), err
}
