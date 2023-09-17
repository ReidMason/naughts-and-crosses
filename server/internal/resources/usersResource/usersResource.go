package usersResource

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ReidMason/naughts-and-crosses/server/internal/database"
	httpHelper "github.com/ReidMason/naughts-and-crosses/server/internal/helpers"
	"github.com/go-chi/chi"
)

type usersResource struct {
	userService UserService
}

type UserService interface {
	CreateUser(name string) (database.User, error)
	GetUser(id int32) (database.User, error)
}

func New(userService UserService) *usersResource {
	return &usersResource{
		userService: userService,
	}
}

func (rs usersResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", rs.Create)
	r.Route("/{userId}", func(r chi.Router) {
		r.Use(rs.userCtx)
		r.Get("/", rs.Get)
	})

	return r
}

func (ur usersResource) userCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := strconv.ParseInt(chi.URLParam(r, "userId"), 10, 32)
		if err != nil {
			slog.Info("Failed to convert userId", err)
			httpHelper.SendResponse[interface{}](w, nil, false, "User not found", http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "user", int32(userId))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type NewUserDTO struct {
	Name string `json:"name"`
}

func (rs usersResource) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newUser NewUserDTO
	err := decoder.Decode(&newUser)
	if err != nil {
		slog.Error("Failed to parse request body", err)
		httpHelper.SendResponse[interface{}](w, nil, false, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(newUser.Name) == "" {
		slog.Error("Name field missing from request body")
		httpHelper.SendResponse[interface{}](w, nil, false, "The 'name' field is required", http.StatusBadRequest)
		return
	}

	user, err := rs.userService.CreateUser(newUser.Name)
	if err != nil {
		slog.Error("Error creating user", err)
		httpHelper.SendResponse[interface{}](w, nil, false, "Error creating user", http.StatusInternalServerError)
		return
	}

	httpHelper.SendResponse(w, &user, true, "New user created", http.StatusCreated)
}

type UserDTO struct {
	Name        string
	ID          int32
	Wins        int64
	Losses      int64
	DateCreated time.Time
}

func (ur usersResource) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("user").(int32)

	user, err := ur.userService.GetUser(userId)
	if err != nil {
		httpHelper.SendResponse[interface{}](w, nil, false, "User not found", http.StatusNotFound)
		return
	}

	userResponse := UserDTO{
		Name:   user.Name,
		ID:     user.ID,
		Wins:   user.Wins,
		Losses: user.Losses,
	}
	httpHelper.SendResponse(w, &userResponse, true, "User found", http.StatusOK)
}
