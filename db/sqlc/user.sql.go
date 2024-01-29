// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    role,
    password
) VALUES (
  $1, $2, $3, $4
) RETURNING id, created_at, updated_at, username, email, role, password
`

type CreateUserParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.Role,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.Email,
		&i.Role,
		&i.Password,
	)
	return i, err
}
