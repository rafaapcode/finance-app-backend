package goals

import "github.com/rafaapcode/finance-app-backend/api/model"

type GoalsController struct {
	Category   string
	Percentage float64
	Userid     string

	Repo GoalsRepository
}

func (g GoalsController) CreateGoal() (model.Goals, int, error) {
	return model.Goals{}, 200, nil
}

func (g GoalsController) ListAllGoals() ([]model.Goals, int, error) {
	return []model.Goals{}, 200, nil
}

func (g GoalsController) DeleteGoal() (string, int, error) {
	return "", 200, nil
}
