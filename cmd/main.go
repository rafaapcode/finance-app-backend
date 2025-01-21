package main

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rafaapcode/finance-app-backend/api/model"
	"github.com/rafaapcode/finance-app-backend/config"
	"github.com/rafaapcode/finance-app-backend/internal/income"
	"github.com/rafaapcode/finance-app-backend/internal/outcome"
	"github.com/rafaapcode/finance-app-backend/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn, err := config.GetDSN()
	if err != nil {
		panic("Error to read the environment variables")
	}
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error to connect with database")
	}

	db = database
}

func main() {
	e := echo.New()

	// Repos
	userRepo := user.UserRepo{
		DB: db,
	}
	outcomeRepo := outcome.OutcomeRepo{DB: db}
	incomeRepo := income.IncomeRepo{DB: db}

	// Groups Routes
	userRoutes := e.Group("/user")
	authRoutes := e.Group("/auth")
	outcomeRoutes := e.Group("/outcome")
	incomeRoutes := e.Group("/income")

	// User Routes
	userRoutes.POST("/", func(c echo.Context) error {
		// Pegar Nome e Email
		user := user.UserController{Repo: userRepo}
		_, code, err := user.CreateUser()
		if err != nil {
			return c.String(code, err.Error())
		}
		return c.String(code, "msg")
	})
	userRoutes.GET("/", func(c echo.Context) error {
		user := user.UserController{Repo: userRepo}
		_, code, err := user.GetUser()
		if err != nil {
			return c.String(code, err.Error())
		}
		return c.String(code, "msg")
	})
	userRoutes.PUT("/", func(c echo.Context) error {
		user := user.UserController{Repo: userRepo}
		_, code, err := user.UpdateUser("", "")
		if err != nil {
			return c.String(code, err.Error())
		}
		return c.String(code, "msg")
	})
	userRoutes.DELETE("/", func(c echo.Context) error {
		user := user.UserController{Repo: userRepo}
		_, code, err := user.DeleteUser()
		if err != nil {
			return c.String(code, err.Error())
		}
		return c.String(code, "msg")
	})

	// Auth routes
	authRoutes.POST("/", func(c echo.Context) error {
		return c.String(200, "teste")
	})

	// Outcome routes
	outcomeRoutes.POST("/", func(c echo.Context) error {
		outcomeController := outcome.OutComeController{Repo: outcomeRepo}

		// Criando Outcome
		outcomeController.CreateOutcome()
		return c.String(200, "teste")
	})
	outcomeRoutes.GET("/total", func(c echo.Context) error {
		outcomeController := outcome.OutComeController{Repo: outcomeRepo}

		// Retornando o total de gastos no mes
		outcomeController.GetTotalOfOutcomeOfMonth()
		return c.String(200, "teste")
	})
	outcomeRoutes.GET("/", func(c echo.Context) error {
		outcomeController := outcome.OutComeController{Repo: outcomeRepo}

		// Retornando todos os outcomes
		outcomeController.GetAllOutcomeOfMonth()
		return c.String(200, "teste")
	})
	outcomeRoutes.GET("/metrics/outcomes_by_category", func(c echo.Context) error {
		outcomeController := outcome.OutComeController{Repo: outcomeRepo}

		// Retornando métricas
		outcomeController.GetMetricsOfOutcomeByCategory()
		return c.String(200, "")
	})

	// Income routes
	incomeRoutes.POST("/", func(c echo.Context) error {
		incomeController := income.IncomeController{Repo: incomeRepo}

		// Criando Outcome
		incomeController.CreateIncome(model.Income{})
		return c.String(200, "teste")
	})
	incomeRoutes.GET("/total", func(c echo.Context) error {
		incomeController := income.IncomeController{Repo: incomeRepo}

		// Criando Outcome
		incomeController.GetAllIncomeOfMonth(time.Now())
		return c.String(200, "teste")
	})

	port, err := config.GetPort()
	if err != nil {
		panic("Error to read the environment variables")
	}
	e.Logger.Fatal(e.Start(port))
}
