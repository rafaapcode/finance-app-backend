package investment

import (
	"time"

	"github.com/rafaapcode/finance-app-backend/api/model"
)

type InvestmenController struct {
	Category   string
	Quantity   int
	Userid     string
	StockPrice float64
	StockName  string
	Date       time.Time
	SupplyDate time.Time
	SellDate   time.Time

	Repo InvestmentRepository
}

func (inv InvestmenController) CreateInvestment() (model.Investment, int, error) {
	return model.Investment{}, 200, nil
}

func (inv InvestmenController) GetTotalOfInvestment() ([]model.Investment, int, error) {
	return []model.Investment{}, 200, nil
}

func (inv InvestmenController) GetAllOfInvestment(pageNumber string, sort string) ([]model.Investment, int, error) {
	return []model.Investment{}, 200, nil
}

func (inv InvestmenController) GetInvestmentByName(pageNumber string) (model.Investment, int, error) {
	return model.Investment{}, 200, nil
}

func (inv InvestmenController) GetInvestmentByCategory(pageNumber string) ([]model.Investment, int, error) {
	return []model.Investment{}, 200, nil
}

func (inv InvestmenController) GetInvestmentByType() ([]model.Investment, int, error) {
	return []model.Investment{}, 200, nil
}

func (inv InvestmenController) UpdateInvestment(id string) (model.Investment, int, error) {
	return model.Investment{}, 200, nil
}

func (inv InvestmenController) GetAssetGrowth() (model.Metrics, int, error) {
	return model.Metrics{}, 200, nil
}

func (inv InvestmenController) GetPortfolioDiversification() (model.Metrics, int, error) {
	return model.Metrics{}, 200, nil
}

func (inv InvestmenController) GetMonthInvestment() (model.Metrics, int, error) {
	return model.Metrics{}, 200, nil
}
