package service

import (
	"github.com/rostekus/ghtm/internal/app/service/dto"
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

func (a *app) CreateUser(dto.UserDTO) (dto.UserDTO, error) {
	return dto.UserDTO{}, nil
}

func (a *app) GetUser(dto.UserDTO) (dto.UserDTO, error) {
	return dto.UserDTO{}, nil
}
