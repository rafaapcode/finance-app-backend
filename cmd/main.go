package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "host=localhost user=root password=root dbname=finances port=5433 sslmode=disable TimeZone=America/Sao_Paulo"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error to connect with database")
	}

	db = database
}

func main() {
	e := echo.New()

	// Repos

	// Groups Routes
	// userRoutes := e.Group("/user")
	authRoutes := e.Group("/auth")

	// User Routes
	// userRoutes.POST("/", func(c echo.Context) error {
	// 	// Pegar Nome e Email
	// 	user := user.UserController{DB: db}
	// 	msg, code, err := user.CreateUser()
	// 	if err != nil {
	// 		return c.String(code, err.Error())
	// 	}
	// 	return c.String(code, msg)
	// })
	// userRoutes.GET("/", func(c echo.Context) error {
	// 	user := user.UserController{DB: db}
	// 	msg, code, err := user.GetUser()
	// 	if err != nil {
	// 		return c.String(code, err.Error())
	// 	}
	// 	return c.String(code, msg)
	// })
	// userRoutes.PUT("/", func(c echo.Context) error {
	// 	user := user.UserController{DB: db}
	// 	msg, code, err := user.UpdateUser("", "")
	// 	if err != nil {
	// 		return c.String(code, err.Error())
	// 	}
	// 	return c.String(code, msg)
	// })
	// userRoutes.DELETE("/", func(c echo.Context) error {
	// 	user := user.UserController{DB: db}
	// 	msg, code, err := user.DeleteUser()
	// 	if err != nil {
	// 		return c.String(code, err.Error())
	// 	}
	// 	return c.String(code, msg)
	// })

	// Auth routes
	authRoutes.POST("/", func(c echo.Context) error {
		return c.String(200, "teste")
	})

	e.Logger.Fatal(e.Start(":3001"))
}
