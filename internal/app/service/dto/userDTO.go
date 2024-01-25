package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserDTO struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Id        uuid.UUID `json:"id"`
}
