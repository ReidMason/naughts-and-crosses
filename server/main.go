package main

import (
	"database/sql"
	"log"
	"net/http"

	"log/slog"

	"github.com/ReidMason/naughts-and-crosses/server/internal/migrations"
	"github.com/ReidMason/naughts-and-crosses/server/internal/resources/usersResource"
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

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/user", usersResource.New().Routes())
	})

	slog.Info("Http server started")
	http.ListenAndServe("localhost:3000", r)
}
