package database

import "database/sql"

type IncomeDb struct {
	DB *sql.DB
}

func NewIncomeDB(db *sql.DB) *InvestmentDb {
	return &InvestmentDb{
		DB: db,
	}
}
