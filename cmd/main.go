package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
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
	googleApp := config.OauthConfig()

	r := chi.NewRouter()
	usersDb := database.NewUserDb(db)
	userHandler := handlers.NewUserHandler(usersDb, googleApp)

	incomeDb := database.NewIncomeDB(db)
	extraIncomeDb := database.NewExtraIncomeDB(db)
	incomeHandler := handlers.NewIncomeHandler(incomeDb, extraIncomeDb)
	extraHandler := handlers.NewExtraHandler(extraIncomeDb, usersDb)

	outcomeDb := database.NewOutcomeDb(db)
	outcomeHandler := handlers.NewOutcomeHandler(outcomeDb)

	goalsDb := database.NewGoalsDB(db)
	goalHandler := handlers.NewGoalsHandler(goalsDb, usersDb)

	investmentDb := database.NewInvestmentDB(db)
	buyOp := database.NewBuyOperationDB(db)
	sellOp := database.NewSellOperationDB(db)
	supply := database.NewSupplyOperationDB(db)
	investmentHandler := handlers.NewInvestmentHandler(investmentDb, buyOp, sellOp, supply)

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Middleware with Contexts
	r.Use(middleware.WithValue("token", jwtConfig.TokenAuth))
	r.Use(middleware.WithValue("jwtexp", jwtConfig.Jwtexpires))

	r.Route("/users", func(r chi.Router) {
		r.Get("/auth", userHandler.Auth)
		r.Get("/auth/callback", userHandler.CallbackAuth)
	})

	r.Route("/incomes", func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwtConfig.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/{userid}", incomeHandler.CreateIncome)
		r.Get("/{userid}", incomeHandler.GetTotalIncomeOfUser)
		r.Delete("/{id}", incomeHandler.DeleteIncomeById)
		r.Patch("/{userid}/{value}", incomeHandler.UpdateIncome)
	})

	r.Route("/extraincomes", func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwtConfig.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", extraHandler.CreateExtraIncome)
		r.Get("/allofmonth/{userid}/{month}", extraHandler.GetAllExtraIncomeOfMonth)
		r.Get("/{id}", extraHandler.GetExtraIncomeById)
	})

	r.Route("/outcome", func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwtConfig.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", outcomeHandler.CreateOutcome)
		r.Get("/{id}", outcomeHandler.GetOutcomeById)
		r.Get("/month/{userid}/{month}", outcomeHandler.GetAllOutcomeOfMonth)
		r.Get("/allfixed/{userid}", outcomeHandler.GetAllFixedOutcome)
		r.Get("/category/{userid}/{category}", outcomeHandler.GetAllOutcomeByCategory)
		r.Get("/payment/{userid}/{paymentmethod}", outcomeHandler.GetAllOutcomeByPaymentMethod)
		r.Get("/type/{userid}/{type}", outcomeHandler.GetAllOutcomeByType)
		r.Get("/aboutexpire/{userid}/{daysToExpire}", outcomeHandler.GetOutcomeAboutToExpire)
		r.Get("/less/{userid}/{value}", outcomeHandler.GetOutcomeLessThan)
		r.Get("/higher/{userid}/{value}", outcomeHandler.GetOutcomeHigherThan)
		r.Delete("/{id}", outcomeHandler.DeleteOutcome)
	})

	r.Route("/goals", func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwtConfig.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", goalHandler.CreateGoals)
		r.Get("/{userid}", goalHandler.ListAllGoals)
		r.Patch("/{id}/{percentage}", goalHandler.UpdateGoal)
		r.Delete("/{id}", goalHandler.DeleteGoal)
	})

	r.Route("/investments", func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwtConfig.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", investmentHandler.CreateInvestmentHandler)
		r.Get("/total/{userid}", investmentHandler.GetTotalOfInvestment)
		r.Get("/allinvestment/{userid}", investmentHandler.GetAllOfInvestment)
		r.Get("/allinvestment/next/{userid}/{lastinvid}", investmentHandler.GetNextPageOfAllInvestment)
		r.Get("/allinvestment/previous/{userid}/{firstinvid}", investmentHandler.GetPreviousPageOfAllInvestment)
		r.Get("/investment/name/{userid}/{stockname}", investmentHandler.GetInvesmentByName)
		r.Get("/investment/id/{invId}", investmentHandler.GetInvesmentById)
		r.Get("/investment/category/{userid}/{category}", investmentHandler.GetInvesmentByCategory)
		r.Get("/investment/category/next/{userid}/{category}/{lastinvid}", investmentHandler.GetNextPageInvesmentByCategory)
		r.Get("/investment/category/previous/{userid}/{category}/{firstInvId}", investmentHandler.GetPreviousPageInvesmentByCategory)
		r.Get("/investment/metrics/{userid}", investmentHandler.GetAssetGrowth)
		r.Get("/investment/portfolio/{userid}", investmentHandler.GetPortfolioDiversification)
		r.Get("/investment/month/{userid}/{month}", investmentHandler.GetMonthInvestment)
		r.Put("/sell", investmentHandler.UpdateSellInvesment)
		r.Put("/supply", investmentHandler.UpdateSupplyInvesment)
	})

	http.ListenAndServe(config.GetPort(), r)
}
