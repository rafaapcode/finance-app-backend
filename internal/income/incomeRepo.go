package income

import (
	"time"

	"github.com/rafaapcode/finance-app-backend/api/model"
	"gorm.io/gorm"
)

type IncomeRepository interface {
	CreateIncome(income model.Income) (model.Income, int, error)
	GetIncomeById(id string) (model.Income, int, error)
	GetAllIncomeOfMonth(month time.Time) ([]model.Income, int, error)
	UpdateIncome(id string, newData model.Income) (model.Income, int, error)
	DeleteIncome(id string) (string, int, error)
}

type IncomeRepo struct {
	DB *gorm.DB
}

func (inc IncomeRepo) CreateIncome(income model.Income) (model.Income, int, error) {
	return model.Income{}, 200, nil
}

func (inc IncomeRepo) GetIncomeById(id string) (model.Income, int, error) {
	return model.Income{}, 200, nil
}

func (inc IncomeRepo) GetAllIncomeOfMonth(month time.Time) ([]model.Income, int, error) {
	return []model.Income{}, 200, nil
}

func (inc IncomeRepo) UpdateIncome(id string, newData model.Income) (model.Income, int, error) {
	return model.Income{}, 200, nil
}

func (inc IncomeRepo) DeleteIncome(id string) (string, int, error) {
	return "", 200, nil
}
