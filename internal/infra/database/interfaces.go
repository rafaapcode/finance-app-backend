package database

import (
	"time"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
)

type UserInterface interface {
	CreateUser(user *entity.User) (int, error)
	GetUser(id string) (*entity.User, int, error)
	DeleteUser(id string) (string, int, error)
	UpdateUser(newUserData *entity.User) (int, error)
}

type OutcomeInterface interface {
	CreateOutcome(outcome *entity.Outcome) (int, error)
	GetOutcomeById(id string) (*entity.Outcome, int, error)
	GetAllOutcomeOfMonth(month time.Time, userId string) ([]entity.Outcome, int, error)
	GetAllFixedOutcome(userId string) ([]entity.Outcome, int, error)
	GetAllOutcomeByCategory(category string, userId string) ([]entity.Outcome, int, error)
	GetAllOutcomeByPaymentMethod(method string, userId string) ([]entity.Outcome, int, error)
	GetAllOutcomeByType(typeOutcome string, userId string) ([]entity.Outcome, int, error)
	GetOutcomeAboutToExpire(daysToExpire int, userId string) ([]entity.Outcome, int, error)
	GetOutcomeLessThan(value float64, userId string) ([]entity.Outcome, int, error)
	GetOutcomeMoreThan(value float64, userId string) ([]entity.Outcome, int, error)
	UpdateOutcome(newData *entity.Outcome) (int, error)
	DeleteOutcome(id string, userId string) (string, int, error)
}

type InvestmentInterface interface {
	CreateInvestment(invest *entity.Investment) (int, error)
	GetTotalOfInvestment(userid string) (int, int, error)
	GetAllOfInvestment(pageNumber int, sort, userid string) ([]entity.Investment, int, error)
	GetInvestmentByName(name, userid string) (entity.Investment, int, error)
	GetInvestmentByCategory(pageNumber int, category, userid string) ([]entity.Investment, int, error)
	UpdateInvestment(newData *entity.Investment) (int, error)
	GetAssetGrowth(userid string) (entity.Metrics, int, error)
	GetPortfolioDiversification(userid string) (entity.Metrics, int, error)
	GetMonthInvestment(userid string, month time.Month) (entity.Metrics, int, error)
}

// type ProductInterface interface {
// 	Create(product *entity.Product) error
// 	FindAll(page, limit int, sort string) ([]*entity.Product, error)
// 	FindById(id string) (*entity.Product, error)
// 	Update(prduct *entity.Product) error
// 	Delete(id string) error
// }
