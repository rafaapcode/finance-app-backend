package model

import (
	"time"
)

type Income struct {
	Id        string    `json:"id"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
	Userid    string    `json:"userid"`
	User      User      `json:"user"`
}
