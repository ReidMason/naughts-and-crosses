// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: queries.sql

package database

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  name, token
) VALUES (
  $1, $2
)
RETURNING id, name, token, wins, losses
`

type CreateUserParams struct {
	Name  string
	Token string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.Token)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Token,
		&i.Wins,
		&i.Losses,
	)
	return i, err
}
