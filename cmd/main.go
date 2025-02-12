package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rafaapcode/finance-app-backend/config"
	"github.com/rafaapcode/finance-app-backend/internal/infra/database"
	"github.com/rafaapcode/finance-app-backend/internal/infra/webservers/handlers"
)

var db *sql.DB

func init() {
	err := godotenv.Load()

	if err != nil {
		panic("Error to load the environment variables")
	}
	dsn := config.GetDSN()
	database, err := sql.Open("postgres", dsn)

	if err != nil {
		panic("Error to connect with database")
	}

	db = database
}

func main() {
	jwtConfig := config.GetJwtSecrets()
	defer db.Close()
	r := chi.NewRouter()
	usersDb := database.NewUserDb(db)
	userHandler := handlers.NewUserHandler(usersDb)

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Middleware with Contexts
	r.Use(middleware.WithValue("token", jwtConfig.TokenAuth))
	r.Use(middleware.WithValue("jwtexp", jwtConfig.Jwtexpires))

	r.Route("/users", func(r chi.Router) {
		r.Post("/token", userHandler.GetJwt)
		r.Post("/", userHandler.Create)
	})

	http.ListenAndServe(config.GetPort(), r)
}
