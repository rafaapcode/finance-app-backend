package model

import (
	"time"
)

type Outcome struct {
	Id            string    `json:"id"`
	Type          string    `json:"type"`
	Category      string    `json:"category"`
	Value         float64   `json:"value"`
	PaymentMethod string    `json:"paymentMethod"`
	Notification  bool      `json:"notification"`
	Userid        string    `json:"userid"`
	User          User      `json:"user"`
	ExpireDate    time.Time `json:"expireDate"`
	CreatedAt     time.Time `json:"createdAt"`
	Date          time.Time `json:"date"`
}
