package routesResource

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
)

type RoutesResource struct{}

func New() *RoutesResource {
	return &RoutesResource{}
}

func (rr RoutesResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rr.index)

	return r
}

func (rr RoutesResource) index(w http.ResponseWriter, _ *http.Request) {
	templ := template.Must(template.ParseFiles("internal/templates/index.html"))
	templ.Execute(w, nil)
}
