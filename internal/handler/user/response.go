package handler

import "github.com/rostekus/ghtm/internal/app/service/dto"

type UserLoginResponse struct {
	User dto.UserDTO `json:"user"`
}
