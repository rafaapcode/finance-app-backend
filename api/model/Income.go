package model

import (
	"time"
)

type Income struct {
	Id       string    `json:"id"`
	Type     string    `json:"type"`
	Category string    `json:"category"`
	Value    float64   `json:"value"`
	Date     time.Time `json:"date"`
	Userid   string    `json:"userid"`
	User     User      `json:"user"`
}
