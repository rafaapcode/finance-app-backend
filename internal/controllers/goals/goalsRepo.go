package goals

import (
	"github.com/rafaapcode/finance-app-backend/api/model"
	"gorm.io/gorm"
)

type GoalsRepository interface {
	CreateGoal(goals model.Goals) (model.Goals, int, error)
	UpdateGoal(goalId string, newGoal model.Goals) (model.Goals, int, error)
	DeleteGoal(goalId string) (string, int, error)
	ListAllGoals(userId string) ([]model.Goals, int, error)
}

type GoalsRepo struct {
	DB *gorm.DB
}

func (gr GoalsRepo) CreateGoal(goals model.Goals) (model.Goals, int, error) {
	return model.Goals{}, 200, nil
}

func (gr GoalsRepo) UpdateGoal(goalId string, newGoal model.Goals) (model.Goals, int, error) {
	return model.Goals{}, 200, nil
}

func (gr GoalsRepo) DeleteGoal(goalId string) (string, int, error) {
	return "", 200, nil
}

func (gr GoalsRepo) ListAllGoals(userId string) ([]model.Goals, int, error) {
	return []model.Goals{}, 200, nil
}
