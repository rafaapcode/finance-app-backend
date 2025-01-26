package entity

import (
	"time"

	"github.com/rafaapcode/finance-app-backend/pkg"
)

type Outcome struct {
	Id            pkg.ID    `json:"id"`
	Type          string    `json:"type"`
	Category      string    `json:"category"`
	Value         float64   `json:"value"`
	PaymentMethod string    `json:"paymentMethod"`
	Notification  bool      `json:"notification"`
	Userid        string    `json:"userid"`
	User          User      `json:"user"`
	ExpireDate    time.Time `json:"expireDate"`
	CreatedAt     time.Time `json:"createdAt"`
}

func NewOutcome(outcomeType, category, paymentMethod, userId string, value float64, notification bool, expireDate time.Time) (*Outcome, error) {
	id, err := pkg.NewUUID()

	if err != nil {
		return nil, err
	}

	if !notification {
		return &Outcome{
			Id:            id,
			Type:          outcomeType,
			Category:      category,
			PaymentMethod: paymentMethod,
			Userid:        userId,
			Value:         value,
			Notification:  notification,
			CreatedAt:     time.Now(),
		}, nil
	}

	return &Outcome{
		Id:            id,
		Type:          outcomeType,
		Category:      category,
		PaymentMethod: paymentMethod,
		Userid:        userId,
		Value:         value,
		Notification:  notification,
		ExpireDate:    expireDate,
		CreatedAt:     time.Now(),
	}, nil
}

func (out *Outcome) Validate() error {
	if out.Id.String() == "" {
		return ErrIdIdRequired
	}
	if _, err := pkg.ParseID(out.Id.String()); err != nil {
		return ErrIdInvalidId
	}
	if out.Userid == "" {
		return ErrUserIdIsRequired
	}
	if _, err := pkg.ParseID(out.Userid); err != nil {
		return ErrUserIdIsInvalid
	}
	if out.Type == "" {
		return ErrTypeIsRequired
	}
	if out.Category == "" {
		return ErrCategoryIsRequired
	}
	if out.PaymentMethod == "" {
		return ErrPaymentMethodIsRequired
	}
	if len(out.PaymentMethod) < 3 {
		return ErrPaymentMethodIsInvalid
	}
	if out.Value < 0 {
		return ErrValueIsInvalid
	}
	return nil
}
