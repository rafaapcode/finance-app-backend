package model

import "time"

type ExtraIncome struct {
	Id       string    `json:"id"`
	Userid   string    `json:"userId"`
	User     User      `json:"user"`
	Category string    `json:"category"`
	Value    float64   `json:"value"`
	Date     time.Time `json:"date"`
}
