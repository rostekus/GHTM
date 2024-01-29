package mapper

import (
	"github.com/rostekus/ghtm/internal/app/domain"
	"github.com/rostekus/ghtm/internal/app/service/dto"
)

func UserDTOtoUser(dto dto.UserDTO) domain.User {
	return domain.User{
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
		Username:  dto.Username,
		Email:     dto.Email,
		Password:  dto.Password,
		Id:        dto.Id,
	}
}

func UsertoUserDTO(user domain.User) dto.UserDTO {
	return dto.UserDTO{
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Email:     user.Email,
		Password:  "",
		Id:        user.Id,
	}
}
