package database

import (
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
	DeleteOutcome(id string) (string, int, error)
}

type InvestmentInterface interface {
	CreateInvestment(invest *entity.Investment) (int, error)
	GetTotalOfInvestment(userid string) (float64, int, error)
	GetAllOfInvestment(userid string) ([]entity.Investment, int, error)
	GetNextPageAllOfInvestment(lastInvId, userid string) ([]entity.Investment, int, error)
	GetPreviousPageAllOfInvestment(firstInvId, userid string) ([]entity.Investment, int, error)
	GetInvestmentByName(name, userid string) (entity.Investment, int, error)
	GetInvestmentById(id string) (entity.Investment, int, error)
	GetInvestmentByCategory(category, userid string) ([]entity.Investment, int, error)
	GetNextPageInvestmentByCategory(category, userid, lastInvId string) ([]entity.Investment, int, error)
	GetPreviousPageInvestmentByCategory(category, userid, firstInvId string) ([]entity.Investment, int, error)
	UpdateInvestment(newData *entity.Investment) (int, error)
	GetAssetGrowth(userid string) (entity.Metrics, int, error)
	GetPortfolioDiversification(userid string) (entity.Metrics, int, error)
	GetMonthInvestment(userid string, month int) (entity.Metrics, int, error)
}

type IncomeInterface interface {
	CreateIncome(income *entity.Income) (int, error)
	GetIncomeValueByUserId(userId string) (float64, int, error)
	DeleteIncome(id string) (int, error)
	UpdateIncome(userId string, newValue float64) (int, error)
	GetIncomeByUserId(userId string) (bool, int, error)
}

type ExtraIncomeInterface interface {
	CreateExtraIncome(extraincome *entity.ExtraIncome) (int, error)
	GetExtraIncomeById(id string) (*entity.ExtraIncome, int, error)
	GetAllExtraIncomeOfMonth(month int, userId string) ([]entity.ExtraIncome, int, error)
	GetTotalValueOfExtracIncomeOfTheMonth(month int, userId string) (float64, int, error)
	DeleteExtraIncome(id string) (int, error)
}

type GoalsInterface interface {
	CreateGoal(goals *entity.Goals) (int, error)
	UpdateGoal(id string, newPercentage float64) (int, error)
	DeleteGoal(goalId string) (int, error)
	ListAllGoals(userId string) ([]entity.Goals, int, error)
	SumPercentageOfAllGoals(userId string, percentage float64) (float64, int, error)
	SumPercentageForUpdateGoals(userId, goalid string, percentage float64) (float64, int, error)
	GetGoal(id string) (*entity.Goals, int, error)
}

type SellOperationInterface interface {
	CreateSellOperation(sellOp *entity.SellOperation) (int, error)
	GetSellOperationByName(investmentId string, stockName string) (*entity.SellOperation, int, error)
	GetSellOperationsLessThan(investmentId string, value float64) ([]entity.SellOperation, int, error)
	GetSellOperationsHigherThan(investmentId string, value float64) ([]entity.SellOperation, int, error)
	GetSellOperationsByCategory(investmentId, category string) ([]entity.SellOperation, int, error)
	ListAllSellOperation(investmentId string) ([]entity.SellOperation, int, error)
	ListAllSellOperationOfMonth(investmentId string, month int) ([]entity.SellOperation, int, error)
}

type BuyOperationInterface interface {
	CreateBuyOperation(buyOp *entity.BuyOperation) (int, error)
	GetBuyOperationByName(investmentId string, stockName string) (*entity.BuyOperation, int, error)
	GetBuyOperationsLessThan(investmentId string, value float64) ([]entity.BuyOperation, int, error)
	GetBuyOperationsHigherThan(investmentId string, value float64) ([]entity.BuyOperation, int, error)
	GetBuyOperationsByCategory(investmentId, category string) ([]entity.BuyOperation, int, error)
	ListAllBuyOperation(investmentId string) ([]entity.BuyOperation, int, error)
	ListAllBuyOperationOfMonth(investmentId string, month int) ([]entity.BuyOperation, int, error)
}

type SupplyOperationInterface interface {
	CreateSupplyOperation(supplyOp *entity.SupplyOperation) (int, error)
	GetSupplyOperationByName(investmentId string, stockName string) ([]entity.SupplyOperation, int, error)
	GetSupplyOperationsLessThan(investmentId string, value float64) ([]entity.SupplyOperation, int, error)
	GetSupplyOperationsHigherThan(investmentId string, value float64) ([]entity.SupplyOperation, int, error)
	GetSupplyOperationsByCategory(investmentId, category string) ([]entity.SupplyOperation, int, error)
	ListAllSupplyOperation(investmentId string) ([]entity.SupplyOperation, int, error)
	ListAllSupplyOperationOfMonth(investmentId string, month int) ([]entity.SupplyOperation, int, error)
}
