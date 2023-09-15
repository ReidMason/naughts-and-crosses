package usersResource

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type usersResource struct{}

func New() *usersResource {
	return &usersResource{}
}

func (rs usersResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", rs.Create)

	return r
}

func (rs usersResource) Create(w http.ResponseWriter, r *http.Request) {
	token := uuid.New().String()
	w.Write([]byte(token))
}
