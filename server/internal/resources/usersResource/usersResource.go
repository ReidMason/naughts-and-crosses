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

func (rs usersResource) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var newUser NewUserDTO
	err := decoder.Decode(&newUser)
	if err != nil {
		slog.Error("Failed to parse request body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(newUser.Name) == "" {
		slog.Error("Name field missing from request body")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Include a name"))
		return
	}

	user, err := rs.userService.CreateUser(newUser.Name)
	if err != nil {
		slog.Error("Error creating user", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		slog.Error("Error marshalling user data", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(response))
}
