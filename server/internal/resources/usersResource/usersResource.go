package usersResource

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/ReidMason/naughts-and-crosses/server/internal/database"
	"github.com/go-chi/chi"
)

type usersResource struct {
	userService UserService
}

type UserService interface {
	CreateUser(name string) database.User
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
	decoder := json.NewDecoder(r.Body)
	var newUser NewUserDTO
	err := decoder.Decode(&newUser)
	if err != nil {
		slog.Error("Failed to parse request body", err)
		return
	}

	user := rs.userService.CreateUser(newUser.Name)

	b, err := json.Marshal(user)
	if err != nil {
		slog.Error("Error marshalling user data", err)
		return
	}

	w.Write([]byte(b))
}
