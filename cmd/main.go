package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/rafaapcode/finance-app-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	err := godotenv.Load()

	if err != nil {
		panic("Error to load the environment variables")
	}
	dsn := config.GetDSN()
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error to connect with database")
	}

	db = database
}

func main() {
	e := echo.New()

	// Health Endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "Your API is HEALTHY!")
	})

	port := config.GetPort()
	e.Logger.Fatal(e.Start(port))
}
