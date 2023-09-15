package userService

import (
	"context"
	"log/slog"

	"github.com/ReidMason/naughts-and-crosses/server/internal/database"
	"github.com/google/uuid"
)

type UserService struct {
	ctx     context.Context
	queries *database.Queries
}

func New(ctx context.Context, queries *database.Queries) *UserService {
	return &UserService{
		ctx:     ctx,
		queries: queries,
	}
}

func (us *UserService) CreateUser(name string) database.User {
	slog.Info("Creating new user", "name", name)

	result, err := us.queries.CreateUser(us.ctx, database.CreateUserParams{
		Name:  name,
		Token: uuid.New().String(),
	})

	if err != nil {
		slog.Info("Failed to create new user")
	}

	return result
}
