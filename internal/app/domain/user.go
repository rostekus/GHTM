package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Email     string
	Role      string
	Password  string
	Id        uuid.UUID
}
