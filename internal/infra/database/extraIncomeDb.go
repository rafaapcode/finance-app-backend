package database

import (
	"database/sql"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
)

type ExtraIncomeDb struct {
	DB *sql.DB
}

func NewExtraIncomeDB(db *sql.DB) *ExtraIncomeDb {
	return &ExtraIncomeDb{
		DB: db,
	}
}

func (inc *IncomeDb) CreateExtraIncome(income *entity.Income) (int, error) {
	return 500, nil
}
