package usersResource

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/ReidMason/naughts-and-crosses/server/internal/database"
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
			sendResponse[interface{}](w, nil, false, "User not found", http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "user", int32(userId))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type NewUserDTO struct {
	Name string `json:"name"`
}

type Response[T any] struct {
	Data    *T
	Message string
	Success bool
}

func (rs usersResource) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newUser NewUserDTO
	err := decoder.Decode(&newUser)
	if err != nil {
		slog.Error("Failed to parse request body", err)
		sendResponse[interface{}](w, nil, false, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(newUser.Name) == "" {
		slog.Error("Name field missing from request body")
		sendResponse[interface{}](w, nil, false, "The 'name' field is required", http.StatusBadRequest)
		return
	}

	user, err := rs.userService.CreateUser(newUser.Name)
	if err != nil {
		slog.Error("Error creating user", err)
		sendResponse[interface{}](w, nil, false, "Error creating user", http.StatusInternalServerError)
		return
	}

	sendResponse(w, &user, true, "New user created", http.StatusCreated)
}

func (ur usersResource) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("user").(int32)

	user, err := ur.userService.GetUser(userId)
	if err != nil {
		sendResponse[interface{}](w, nil, false, "User not found", http.StatusNotFound)
		return
	}

	sendResponse(w, &user, true, "User found", http.StatusOK)
}

func sendResponse[T any](w http.ResponseWriter, data *T, success bool, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := Response[T]{
		Data:    data,
		Success: success,
		Message: message,
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		slog.Error("Failed to serialize response", err)
		sendResponse[interface{}](w, nil, false, "Failed to serialize repsonse", http.StatusInternalServerError)
		return
	}

	w.Write(bytes)
}
