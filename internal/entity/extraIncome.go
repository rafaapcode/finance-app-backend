package entity

import (
	"time"

	"github.com/rafaapcode/finance-app-backend/pkg"
)

type ExtraIncome struct {
	Id       pkg.ID    `json:"id"`
	Userid   string    `json:"userId"`
	User     User      `json:"user"`
	Category string    `json:"category"`
	Value    float64   `json:"value"`
	Date     time.Time `json:"date"`
}

func NewExtraIncome(userId, category string, value float64) (*ExtraIncome, error) {
	id, err := pkg.NewUUID()

	if err != nil {
		return nil, err
	}

	return &ExtraIncome{
		Id:       id,
		Userid:   userId,
		Category: category,
		Value:    value,
		Date:     time.Now(),
	}, nil
}
