package income

import (
	"time"

	"github.com/rafaapcode/finance-app-backend/api/model"
)

type IncomeController struct {
	Value  float64
	Userid string
	Repo   IncomeRepository
}

func (incController IncomeController) CreateIncome() (model.Income, int, error) {

	// incController.Repo.CreateIncome(income)

	return model.Income{}, 200, nil
}

func (incController IncomeController) GetAllIncomeOfMonth(month time.Time) ([]model.Income, int, error) {

	incController.Repo.GetAllIncomeOfMonth(month, "")

	return []model.Income{}, 200, nil
}
