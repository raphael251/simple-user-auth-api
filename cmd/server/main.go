package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"

	"github.com/raphael251/simple-user-auth-api/configs"
	"github.com/raphael251/simple-user-auth-api/infra/webserver/handlers"
)

func main() {
	configs, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, configs.DBConnectionString)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", handlers.CreateUserHandler(db))
	})

	http.ListenAndServe(":"+configs.ServerPort, r)
}
