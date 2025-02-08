package database

import (
	"time"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
)

type UserInterface interface {
	CreateUser(user *entity.User) (int, error)
	GetUser(id string) (*entity.User, int, error)
	GetUserByEmail(email string) (*entity.User, int, error)
	DeleteUser(id string) (string, int, error)
	UpdateUser(newUserData *entity.User) (int, error)
}

type OutcomeInterface interface {
	CreateOutcome(outcome *entity.Outcome) (int, error)
	GetOutcomeById(id string) (*entity.Outcome, int, error)
	GetAllOutcomeOfMonth(month int, userId string) ([]entity.Outcome, int, error)
	GetAllFixedOutcome(userId string) ([]entity.Outcome, int, error)
	GetAllOutcomeByCategory(category string, userId string) ([]entity.Outcome, int, error)
	GetAllOutcomeByPaymentMethod(method string, userId string) ([]entity.Outcome, int, error)
	GetAllOutcomeByType(typeOutcome string, userId string) ([]entity.Outcome, int, error)
	GetOutcomeAboutToExpire(daysToExpire int, userId string) ([]entity.Outcome, int, error)
	GetOutcomeLessThan(value float64, userId string) ([]entity.Outcome, int, error)
	GetOutcomeHigherThan(value float64, userId string) ([]entity.Outcome, int, error)
	DeleteOutcome(id string, userId string) (string, int, error)
}

type InvestmentInterface interface {
	CreateInvestment(invest *entity.Investment) (int, error)
	GetTotalOfInvestment(userid string) (int, int, error)
	GetAllOfInvestment(pageNumber int, sort, userid string) ([]entity.Investment, int, error)
	GetInvestmentByName(name, userid string) (entity.Investment, int, error)
	GetInvestmentById(id string) (entity.Investment, int, error)
	GetInvestmentByCategory(pageNumber int, category, userid string) ([]entity.Investment, int, error)
	UpdateInvestment(newData *entity.Investment) (int, error)
	GetAssetGrowth(userid string) (entity.Metrics, int, error)
	GetPortfolioDiversification(userid string) (entity.Metrics, int, error)
	GetMonthInvestment(userid string, month int) (entity.Metrics, int, error)
}

type IncomeInterface interface {
	CreateIncome(income *entity.Income) (int, error)
	GetIncomeById(id string) (*entity.Income, int, error)
	GetAllIncomeOfMonth(month time.Time, userId string) ([]entity.Income, int, error) // AllExtraIncomeOfMonth + Income
	UpdateIncome(newData *entity.Income) (int, error)
	DeleteIncome(id string) (int, error)
}

type GoalsInterface interface {
	CreateGoal(goals entity.Goals) (int, error)
	UpdateGoal(newGoal entity.Goals) (int, error)
	DeleteGoal(goalId string) (string, int, error)
	ListAllGoals(userId string) ([]entity.Goals, int, error)
}

type SellOperationInterface interface {
	CreateSellOperation(sellOp entity.SellOperation) (int, error)
	GetSellOperationByName(investmentId string, stockName string) (entity.SellOperation, int, error)
	GetSellOperationsLessThan(investmentId string, value float64) ([]entity.SellOperation, int, error)
	GetSellOperationsHigherThan(investmentId string, value float64) ([]entity.SellOperation, int, error)
	GetSellOperationsByCategory(investmentId string, category float64) ([]entity.SellOperation, int, error)
	ListAllSellOperation(investmentId string) ([]entity.SellOperation, int, error)
	ListAllSellOperationOfMonth(investmentId string, month time.Time) ([]entity.SellOperation, int, error)
}

type BuyOperationInterface interface {
	CreateBuyOperation(buyOp entity.BuyOperation) (int, error)
	GetBuyOperationByName(investmentId string, stockName string) (entity.BuyOperation, int, error)
	GetBuyOperationsLessThan(investmentId string, value float64) ([]entity.BuyOperation, int, error)
	GetBuyOperationsHigherThan(investmentId string, value float64) ([]entity.BuyOperation, int, error)
	GetBuyOperationsByCategory(investmentId string, category float64) ([]entity.BuyOperation, int, error)
	ListAllBuyOperation(investmentId string) ([]entity.BuyOperation, int, error)
	ListAllBuyOperationOfMonth(investmentId string, month time.Time) ([]entity.BuyOperation, int, error)
}

type SupplyOperationInterface interface {
	CreateSupplyOperation(supplyOp entity.SupplyOperation) (int, error)
	GetSupplyOperationByName(investmentId string, stockName string) ([]entity.SupplyOperation, int, error)
	GetSupplyOperationsLessThan(investmentId string, value float64) ([]entity.SupplyOperation, int, error)
	GetSupplyOperationsHigherThan(investmentId string, value float64) ([]entity.SupplyOperation, int, error)
	GetSupplyOperationsByCategory(investmentId string, category float64) ([]entity.SupplyOperation, int, error)
	ListAllSupplyOperation(investmentId string) ([]entity.SupplyOperation, int, error)
	ListAllSupplyOperationOfMonth(investmentId string, month time.Time) ([]entity.SupplyOperation, int, error)
}
