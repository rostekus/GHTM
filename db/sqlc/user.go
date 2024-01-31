package db

import (
	"context"

	"github.com/rostekus/ghtm/internal/app/domain"
)

func (s *SQLStore) CreateUser(c context.Context, u domain.User) (domain.User, error) {
	arg := CreateUserParams{
		Username: u.Username,
		Email:    u.Email,
		Role:     u.Role,
		Password: u.Password,
	}
	usr, err := s.Queries.CreateUser(c, arg)
	if err != nil {
		return domain.User{}, err
	}

	domainUser := mapUserDBToUser(usr)
	return domainUser, nil
}

func (s *SQLStore) GetUser(c context.Context, u domain.User) (domain.User, error) {
	usr, err := s.GetUserByEmail(c, u.Email)
	if err != nil {
		return domain.User{}, err
	}
	return mapUserDBToUser(usr), nil
}
