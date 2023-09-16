-- name: CreateUser :one
INSERT INTO users (
  name, token
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users 
WHERE id = $1;
