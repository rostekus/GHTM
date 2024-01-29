-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    role,
    password
) VALUES (
  $1, $2, $3, $4
) RETURNING *;
