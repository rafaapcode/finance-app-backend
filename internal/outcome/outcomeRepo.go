package outcome

import (
	"time"

	"github.com/rafaapcode/finance-app-backend/api/model"
	"gorm.io/gorm"
)

type OutcomeRepository interface {
	CreateOutcome(outcome model.Outcome) (model.Outcome, int, error)
	GetOutcomeById(id string) (model.Outcome, int, error)
	GetAllOutcomeOfMonth(month time.Time) ([]model.Outcome, int, error)
	GetAllFixedOutcome() ([]model.Outcome, int, error)
	GetAllOutcomeByCategory(category string) ([]model.Outcome, int, error)
	GetAllOutcomeByPaymentMethod(method string) ([]model.Outcome, int, error)
	GetAllOutcomeByType(typeOutcome string) ([]model.Outcome, int, error)
	GetOutcomeAboutToExpire(daysToExpire int) ([]model.Outcome, int, error)
	GetOutcomeLessThan(value float64) ([]model.Outcome, int, error)
	GetOutcomeMoreThan(value float64) ([]model.Outcome, int, error)
	UpdateOutcome(id string, newData model.Outcome) (model.Outcome, int, error)
	DeleteOutcome(id string) (string, int, error)
}

type OutcomeRepo struct {
	DB *gorm.DB
}

func (out OutcomeRepo) CreateOutcome(outcome model.Outcome) (model.Outcome, int, error) {
	return model.Outcome{}, 200, nil
}

func (out OutcomeRepo) GetAllOutcomeOfMonth(month time.Time) ([]model.Outcome, int, error) {
	return []model.Outcome{}, 200, nil
}

func (out OutcomeRepo) GetAllFixedOutcome() ([]model.Outcome, int, error) {
	return []model.Outcome{}, 200, nil
}

func (out OutcomeRepo) GetAllOutcomeByCategory(category string) ([]model.Outcome, int, error) {
	return []model.Outcome{}, 200, nil
}

func (out OutcomeRepo) GetAllOutcomeByPaymentMethod(method string) ([]model.Outcome, int, error) {
	return []model.Outcome{}, 200, nil
}

func (out OutcomeRepo) GetAllOutcomeByType(typeOutcome string) ([]model.Outcome, int, error) {
	return []model.Outcome{}, 200, nil
}

func (out OutcomeRepo) GetOutcomeAboutToExpire(daysToExpire int) ([]model.Outcome, int, error) {
	return []model.Outcome{}, 200, nil
}

func (out OutcomeRepo) GetOutcomeLessThan(value float64) ([]model.Outcome, int, error) {
	return []model.Outcome{}, 200, nil
}

func (out OutcomeRepo) GetOutcomeMoreThan(value float64) ([]model.Outcome, int, error) {
	return []model.Outcome{}, 200, nil
}

func (out OutcomeRepo) GetOutcomeById(id string) (model.Outcome, int, error) {
	return model.Outcome{}, 200, nil
}

func (out OutcomeRepo) UpdateOutcome(id string, newData model.Outcome) (model.Outcome, int, error) {
	return model.Outcome{}, 200, nil
}

func (out OutcomeRepo) DeleteOutcome(id string) (string, int, error) {
	return "", 200, nil
}
