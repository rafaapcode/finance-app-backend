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
	Repo          OutcomeRepository
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

func (outcome OutComeController) GetTotalOfOutcomeOfMonth() (string, int, error) {
	return "", 200, nil
}

func (outcome OutComeController) GetMetricsOfOutcomeByCategory() (model.OutcomeMetrics, int, error) {
	return model.OutcomeMetrics{}, 200, nil
}
