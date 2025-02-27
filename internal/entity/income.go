package entity

import (
	"time"

	"github.com/rafaapcode/finance-app-backend/pkg"
)

type Income struct {
	Id        pkg.ID    `json:"id"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
	Userid    string    `json:"userid"`
	User      User      `json:"user"`
}

func NewIncome(userId string, value float64) (*Income, error) {
	id, err := pkg.NewUUID()

	if err != nil {
		return nil, err
	}

	return &Income{
		Id:        id,
		Userid:    userId,
		Value:     value,
		CreatedAt: time.Now(),
	}, nil
}

func (inc *Income) Validate() error {
	if inc.Id.String() == "" {
		return ErrIdIdRequired
	}
	if _, err := pkg.ParseID(inc.Id.String()); err != nil {
		return ErrIdInvalidId
	}
	if inc.Userid == "" {
		return ErrUserIdIsRequired
	}
	if _, err := pkg.ParseID(inc.Userid); err != nil {
		return ErrUserIdIsInvalid
	}
	if inc.Value < 0.0 {
		return ErrValueIsInvalid
	}
	return nil
}
