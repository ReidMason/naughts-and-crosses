package main

import (
	"context"
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"log/slog"

	"github.com/ReidMason/naughts-and-crosses/server/internal/database"
	"github.com/ReidMason/naughts-and-crosses/server/internal/migrations"
	"github.com/ReidMason/naughts-and-crosses/server/internal/resources/routesResource"
	"github.com/ReidMason/naughts-and-crosses/server/internal/resources/usersResource"
	"github.com/ReidMason/naughts-and-crosses/server/internal/userService"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	connectionString := "postgres://user:password@localhost:54321/testdb?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Panic("Failed to connect to database", err)
	}

	migrations.Migrate(db)

	ctx := context.Background()
	queries := database.New(db)

	userService := userService.New(ctx, queries)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/user", usersResource.New(userService).Routes())
	})

	r.Mount("/", routesResource.New().Routes())

	r.Get("/login-successful", func(w http.ResponseWriter, _ *http.Request) {
		templ := template.Must(template.ParseFiles("internal/templates/loginSuccessful.html"))
		templ.Execute(w, nil)
	})

	slog.Info("Http server started")
	http.ListenAndServe("localhost:3000", r)
}
