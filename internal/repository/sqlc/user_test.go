package sqlc

import (
	"context"
	"testing"
	"time"

	"github.com/rostekus/ghtm/internal/app/domain"
	"github.com/rostekus/ghtm/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) domain.User {
	ctx := context.Background()
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	email := util.RandomEmail()
	username := util.RandomString(5)

	usr := domain.User{
		Email:    email,
		Username: username,
		Password: hashedPassword,
	}

	user, err := testStore.CreateUser(ctx, usr)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, username, user.Username)
	require.Equal(t, hashedPassword, user.Password)
	require.Equal(t, email, user.Email)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testStore.GetUser(context.Background(), user1)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
