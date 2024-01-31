package domain

import (
	"time"
)

type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Email     string
	Role      string
	Password  string
	Id        int32
}
