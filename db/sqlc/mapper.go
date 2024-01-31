package db

import (
	"github.com/rostekus/ghtm/internal/app/domain"
)

func mapUserDBToUser(userDB User) domain.User {
	return domain.User{
		CreatedAt: userDB.CreatedAt.Time,
		UpdatedAt: userDB.UpdatedAt.Time,
		Username:  userDB.Username,
		Email:     userDB.Email,
		Role:      userDB.Role,
		Password:  userDB.Password,
		Id:        userDB.ID,
	}
}
