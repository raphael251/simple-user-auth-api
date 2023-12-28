package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/raphael251/simple-user-auth-api/configs"
	_ "github.com/raphael251/simple-user-auth-api/docs"
	"github.com/raphael251/simple-user-auth-api/infra/webserver/handlers"
)

// @title 				 						 Simple User Auth API
// @version 			 						 1.0
// @termsOfService 						 http://swagger.io/terms/

// @contact.name 	 						 Raphael Passos
// @contact.url 	 						 http://github.com/raphael251
// @contact.email  						 raphael251@hotmail.com

// @host					 						 localhost:3000
// @BasePath			 						 /
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

	r.Get("/docs/*",
		httpSwagger.Handler(
			httpSwagger.URL(
				fmt.Sprintf("%s://%s:%s/docs/doc.json", configs.ServerProtocol, configs.ServerHost, configs.ServerPort),
			),
		),
	)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.ServerPort), r)
}
