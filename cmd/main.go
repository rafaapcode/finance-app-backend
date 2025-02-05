package main

import (
	"database/sql"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/rafaapcode/finance-app-backend/config"
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

	// err = db.Ping()

	// if err != nil {
	// 	panic("Error to connect with database")
	// }

	db = database
}

func main() {
	e := echo.New()
	defer db.Close()
	// Health Endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "Your API is HEALTHY!")
	})

	port := config.GetPort()
	e.Logger.Fatal(e.Start(port))
}
