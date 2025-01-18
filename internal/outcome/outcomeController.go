package outcome

import (
	"time"

	"github.com/rafaapcode/finance-app-backend/api/model"
)

type OutComeController struct {
	Type          string
	Category      string
	Value         float64
	PaymentMethod string
	Notification  bool
	Userid        string
	ExpireDate    time.Time
}

func (outcome OutComeController) CreateOutcome() (model.Outcome, int, error) {
	return model.Outcome{}, 200, nil
}

func (outcome OutComeController) GetOneOutcome() (model.Outcome, int, error) {
	return model.Outcome{}, 200, nil
}

func (outcome OutComeController) GetAllOutcomeOfMonth() ([]model.Outcome, int, error) {
	return []model.Outcome{}, 200, nil
}

func (outcome OutComeController) GetAllFixedOutcome() ([]model.Outcome, int, error) {
	return []model.Outcome{}, 200, nil
}
