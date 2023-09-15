package usersResource

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"github.com/ReidMason/naughts-and-crosses/server/internal/database"
	"github.com/go-chi/chi"
)

type usersResource struct {
	userService UserService
}

type UserService interface {
	CreateUser(name string) (database.User, error)
}

func New(userService UserService) *usersResource {
	return &usersResource{
		userService: userService,
	}
}

func (rs usersResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", rs.Create)

	return r
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
