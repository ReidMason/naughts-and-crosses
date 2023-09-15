-- name: CreateUser :one
INSERT INTO users (
  name, token
) VALUES (
  $1, $2
)
RETURNING *;
