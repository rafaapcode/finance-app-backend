package entity

import (
	"time"

	"github.com/rafaapcode/finance-app-backend/pkg"
)

type Goals struct {
	Id         pkg.ID    `json:"id"`
	Category   string    `json:"category"`
	Percentage float64   `json:"percentage"`
	Userid     string    `json:"userid"`
	User       User      `json:"user"`
	CreatedAt  time.Time `json:"createdAt"`
}

func NewGoals(userId, category string, percentage float64) (*Goals, error) {
	id, err := pkg.NewUUID()

	if err != nil {
		return nil, err
	}

	return &Goals{
		Id:         id,
		Userid:     userId,
		Category:   category,
		Percentage: percentage,
		CreatedAt:  time.Now(),
	}, nil
}

func (goal *Goals) Validate() error {
	if goal.Id.String() == "" {
		return ErrIdIdRequired
	}
	if _, err := pkg.ParseID(goal.Id.String()); err != nil {
		return ErrIdInvalidId
	}
	if goal.Userid == "" {
		return ErrUserIdIsRequired
	}
	if _, err := pkg.ParseID(goal.Userid); err != nil {
		return ErrUserIdIsInvalid
	}
	if goal.Category == "" {
		return ErrCategoryIsRequired
	}
	if goal.Percentage < 0.0 {
		return ErrPercentageIsInvalid
	}
	return nil
}
