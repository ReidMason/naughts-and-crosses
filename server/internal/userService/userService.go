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

func (us *UserService) GetUserByToken(token string) (database.User, error) {
	var newUser database.User
	newUser, err := us.queries.GetUserByToken(us.ctx, token)

	return newUser, err
}

func (us *UserService) CreateUser(name string) (database.User, error) {
	slog.Info("Creating new user", "name", name)

	var newUser database.User
	newUser, err := us.queries.CreateUser(us.ctx, database.CreateUserParams{
		Name:  name,
		Token: uuid.New().String(),
	})

	if err != nil {
		slog.Info("Failed to create new user")
	}

	return newUser, err
}

func (us *UserService) GetUser(userId int32) (database.User, error) {
	slog.Info("Getting user", "UserID", userId)

	var user database.User
	user, err := us.queries.GetUser(us.ctx, userId)

	if err != nil {
		slog.Info("Failed to find user")
	}

	return user, err
}
