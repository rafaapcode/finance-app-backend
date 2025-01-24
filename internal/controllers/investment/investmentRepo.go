package investment

import (
	"github.com/rafaapcode/finance-app-backend/api/model"
	"gorm.io/gorm"
)

type InvestmentRepository interface {
	CreateInvestment(invest model.Investment) (model.Investment, int, error)
	GetTotalOfInvestment() ([]model.Investment, int, error)
	GetAllOfInvestment(pageNumber string, sort string) ([]model.Investment, int, error)
	GetInvestmentByName(pageNumber string) (model.Investment, int, error)
	GetInvestmentByCategory(pageNumber string) ([]model.Investment, int, error)
	GetInvestmentByType() ([]model.Investment, int, error)
	UpdateInvestment(id string) (model.Investment, int, error)
	GetAssetGrowth() (model.Metrics, int, error)
	GetPortfolioDiversification() (model.Metrics, int, error)
	GetMonthInvestment() (model.Metrics, int, error)
}

type InvestmentRepo struct {
	DB *gorm.DB
}

func (invRepo InvestmentRepo) CreateInvestment(invest model.Investment) (model.Investment, int, error) {
	return model.Investment{}, 200, nil
}

func (invRepo InvestmentRepo) GetTotalOfInvestment() ([]model.Investment, int, error) {
	return []model.Investment{}, 200, nil
}

func (invRepo InvestmentRepo) GetAllOfInvestment(pageNumber string, sort string) ([]model.Investment, int, error) {
	return []model.Investment{}, 200, nil
}

func (invRepo InvestmentRepo) GetInvestmentByName(pageNumber string) (model.Investment, int, error) {
	return model.Investment{}, 200, nil
}

func (invRepo InvestmentRepo) GetInvestmentByCategory(pageNumber string) ([]model.Investment, int, error) {
	return []model.Investment{}, 200, nil
}

func (invRepo InvestmentRepo) GetInvestmentByType() ([]model.Investment, int, error) {
	return []model.Investment{}, 200, nil
}

func (invRepo InvestmentRepo) UpdateInvestment(id string) (model.Investment, int, error) {
	return model.Investment{}, 200, nil
}

func (invRepo InvestmentRepo) GetAssetGrowth() (model.Metrics, int, error) {
	return model.Metrics{}, 200, nil
}

func (invRepo InvestmentRepo) GetPortfolioDiversification() (model.Metrics, int, error) {
	return model.Metrics{}, 200, nil
}

func (invRepo InvestmentRepo) GetMonthInvestment() (model.Metrics, int, error) {
	return model.Metrics{}, 200, nil
}
