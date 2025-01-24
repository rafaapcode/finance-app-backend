package main

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/rafaapcode/finance-app-backend/config"
	"github.com/rafaapcode/finance-app-backend/internal/controllers/goals"
	"github.com/rafaapcode/finance-app-backend/internal/controllers/income"
	"github.com/rafaapcode/finance-app-backend/internal/controllers/investment"
	"github.com/rafaapcode/finance-app-backend/internal/controllers/outcome"
	"github.com/rafaapcode/finance-app-backend/internal/controllers/user"
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

	// Repos
	userRepo := user.UserRepo{
		DB: db,
	}
	outcomeRepo := outcome.OutcomeRepo{DB: db}
	incomeRepo := income.IncomeRepo{DB: db}
	goalsRepo := goals.GoalsRepo{DB: db}
	investmentRepo := investment.InvestmentRepo{DB: db}

	// Groups Routes
	userRoutes := e.Group("/user")
	authRoutes := e.Group("/auth")
	outcomeRoutes := e.Group("/outcome")
	incomeRoutes := e.Group("/income")
	goalsRoutes := e.Group("/goals")
	investmentRoutes := e.Group("/investment")

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
		incomeController.CreateIncome()
		return c.String(200, "teste")
	})
	incomeRoutes.GET("/total", func(c echo.Context) error {
		incomeController := income.IncomeController{Repo: incomeRepo}

		// Criando Outcome
		incomeController.GetAllIncomeOfMonth(time.Now())
		return c.String(200, "teste")
	})

	// Goals Routes
	goalsRoutes.POST("/", func(c echo.Context) error {
		goalsController := goals.GoalsController{Repo: goalsRepo}

		// Create a Goals
		goalsController.CreateGoal()

		return c.String(200, "")
	})

	// InvestmentRoutes
	investmentRoutes.GET("/total", func(c echo.Context) error {
		investmentController := investment.InvestmenController{Repo: investmentRepo}

		investmentController.GetAllOfInvestment("1", "")

		return c.String(200, "")
	})
	investmentRoutes.GET("/:pageNumber", func(c echo.Context) error {
		pageNum := c.Param("pageNumber")
		investmentController := investment.InvestmenController{Repo: investmentRepo}

		investmentController.GetAllOfInvestment(pageNum, "")

		return c.String(200, "")
	})
	investmentRoutes.GET("/:pageNumber/:name", func(c echo.Context) error {
		pageNum := c.Param("pageNumber")
		name := c.Param("name")
		investmentController := investment.InvestmenController{StockName: name, Repo: investmentRepo}

		investmentController.GetInvestmentByName(pageNum)

		return c.String(200, "")
	})
	investmentRoutes.GET("/sort/:typeofSort", func(c echo.Context) error {
		sort := c.Param("typeofSort")
		investmentController := investment.InvestmenController{Repo: investmentRepo}

		investmentController.GetAllOfInvestment("", sort)

		return c.String(200, "")
	})
	investmentRoutes.GET("/category/:category", func(c echo.Context) error {
		cat := c.Param("category")
		investmentController := investment.InvestmenController{Category: cat, Repo: investmentRepo}

		investmentController.GetInvestmentByCategory("")

		return c.String(200, "")
	})
	investmentRoutes.GET("/metrics/asset_growth", func(c echo.Context) error {
		investmentController := investment.InvestmenController{Repo: investmentRepo}

		investmentController.GetAssetGrowth()

		return c.String(200, "")
	})
	investmentRoutes.GET("/metrics/portfolio_diversification", func(c echo.Context) error {
		investmentController := investment.InvestmenController{Repo: investmentRepo}

		investmentController.GetPortfolioDiversification()

		return c.String(200, "")
	})
	investmentRoutes.POST("/", func(c echo.Context) error {
		investmentController := investment.InvestmenController{Repo: investmentRepo}

		investmentController.CreateInvestment()

		return c.String(200, "")
	})
	investmentRoutes.POST("/metrics/month_investment", func(c echo.Context) error {
		investmentController := investment.InvestmenController{Repo: investmentRepo}

		investmentController.GetMonthInvestment()

		return c.String(200, "")
	})

	port := config.GetPort()
	e.Logger.Fatal(e.Start(port))
}
