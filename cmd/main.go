package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	// Initialize database
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "tudo certo")
	})

	e.Logger.Fatal(e.Start(":3001"))
}
