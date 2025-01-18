package model

import "time"

type Investment struct {
	Id         string    `json:"id"`
	Type       string    `json:"type"`
	Category   string    `json:"category"`
	Quantity   int       `json:"quantity"`
	Userid     string    `json:"userid"`
	StockPrice float64   `json:"stockPrice"`
	StockName  string    `json:"stockName"`
	User       User      `json:"user"`
	Date       time.Time `json:"date"`
	CreatedAt  time.Time `json:"createdat"`
	SupplyDate time.Time `json:"supplyDate"`
	SellDate   time.Time `json:"sellDate"`
}
