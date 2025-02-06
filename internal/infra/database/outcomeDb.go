package database

import (
	"database/sql"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
)

type OutcomeDb struct {
	DB *sql.DB
}

func NewOutcomeDb(db *sql.DB) *OutcomeDb {
	return &OutcomeDb{
		DB: db,
	}
}

func (outDb *OutcomeDb) CreateOutcome(outcome *entity.Outcome) (int, error) {
	return 500, nil
}

func (outDb *OutcomeDb) GetOutcomeById(id string) (*entity.Outcome, int, error) {
	var outcome entity.Outcome
	return &outcome, 500, nil
}

func (outDb *OutcomeDb) GetAllOutcomeOfMonth(month int, userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome

	return outcomes, 500, nil
}

func (outDb *OutcomeDb) GetAllFixedOutcome(userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome

	return outcomes, 500, nil
}

func (outDb *OutcomeDb) GetAllOutcomeByCategory(category string, userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome

	return outcomes, 500, nil
}

func (outDb *OutcomeDb) GetAllOutcomeByPaymentMethod(method string, userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome

	return outcomes, 500, nil
}

func (outDb *OutcomeDb) GetAllOutcomeByType(typeOutcome string, userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome

	return outcomes, 500, nil
}

func (outDb *OutcomeDb) GetOutcomeAboutToExpire(daysToExpire int, userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome

	return outcomes, 500, nil
}

func (outDb *OutcomeDb) GetOutcomeLessThan(value float64, userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome

	return outcomes, 500, nil
}

func (outDb *OutcomeDb) GetOutcomeMoreThan(value float64, userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome

	return outcomes, 500, nil
}

func (outDb *OutcomeDb) DeleteOutcome(id string) (string, int, error) {

	return "Saída deletada com sucesso", 500, nil
}
