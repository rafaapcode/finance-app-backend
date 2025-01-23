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
}

func (inv InvestmenController) CreateInvestment() (model.Investment, int, error) {
	return model.Investment{}, 200, nil
}

func (inv InvestmenController) GetInvestmentByName() (model.Investment, int, error) {
	return model.Investment{}, 200, nil
}

func (inv InvestmenController) GetInvestmentByCategory() ([]model.Investment, int, error) {
	return []model.Investment{}, 200, nil
}

func (inv InvestmenController) GetInvestmentByType() ([]model.Investment, int, error) {
	return []model.Investment{}, 200, nil
}

// VOU TER QUE CRIAR 3 tabelas para investimentos ( COMPRA, VENDA e ATUAL )
func (inv InvestmenController) UpdateInvestment(id string) (model.Investment, int, error) {
	return model.Investment{}, 200, nil
}
