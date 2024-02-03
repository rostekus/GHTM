package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rostekus/ghtm/internal/app/service/dto"
	"github.com/rostekus/ghtm/internal/app/service/mapper"
	mock_sqlc "github.com/rostekus/ghtm/internal/repository/sqlc/mock"
	"github.com/rostekus/ghtm/internal/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mock_sqlc.NewMockStore(ctrl)

	usr := createRandomUser(t)
	userDB := mapper.UserDTOtoUser(usr)
	mockStore.EXPECT().CreateUser(gomock.Any(), userDB).Return(userDB, nil).Times(1)

	srv, err := NewApp(mockStore)
	require.NoError(t, err)
	usr1, err := srv.CreateUser(context.Background(), usr)

	require.NoError(t, err)

	require.Equal(t, usr1.Username, usr.Username)
	require.Equal(t, usr1.Password, "")
	require.Equal(t, usr1.Email, usr.Email)
}

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mock_sqlc.NewMockStore(ctrl)

	usr := createRandomUser(t)
	userDB := mapper.UserDTOtoUser(usr)
	mockStore.EXPECT().GetUser(gomock.Any(), userDB).Return(userDB, nil).Times(1)

	srv, err := NewApp(mockStore)
	require.NoError(t, err)
	usr1, err := srv.GetUser(context.Background(), usr)

	require.NoError(t, err)

	require.Equal(t, usr1.Username, usr.Username)
	require.Equal(t, usr1.Password, "")
	require.Equal(t, usr1.Email, usr.Email)
}

func createRandomUser(t *testing.T) dto.UserDTO {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	email := util.RandomEmail()
	username := util.RandomString(5)

	usr := dto.UserDTO{
		Email:    email,
		Username: username,
		Password: hashedPassword,
	}

	return usr
}
