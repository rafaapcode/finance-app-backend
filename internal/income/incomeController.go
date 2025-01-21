package income

import (
	"time"

	"github.com/rafaapcode/finance-app-backend/api/model"
)

type IncomeController struct {
	Id     string  `json:"id"`
	Value  float64 `json:"value"`
	Userid string  `json:"userid"`
	Repo   IncomeRepository
}

func (incController IncomeController) CreateIncome(income model.Income) (model.Income, int, error) {

	incController.Repo.CreateIncome(income)

	return model.Income{}, 200, nil
}

func (incController IncomeController) GetAllIncomeOfMonth(month time.Time) ([]model.Income, int, error) {

	incController.Repo.GetAllIncomeOfMonth(month)

	return []model.Income{}, 200, nil
}
