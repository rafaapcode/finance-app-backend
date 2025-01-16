package model

import "time"

type Goals struct {
	Id         string    `json:"id"`
	Category   string    `json:"category"`
	Percentage float64   `json:"percentage"`
	Userid     string    `json:"userid"`
	User       User      `json:"user"`
	CreatedAt  time.Time `json:"createdAt"`
}
