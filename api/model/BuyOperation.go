package model

import "time"

type BuyOperation struct {
	Id            string     `json:"id"`
	InvestimentId string     `json:"investmentId"`
	Investment    Investment `json:"investment"`
	Category      string     `json:"category"`
	StockCode     string     `json:"stockCode"`
	Quantity      int        `json:"quantity"`
	BuyPrice      float64    `json:"buyPrice"`
	Value         float64    `json:"value"`
	CreatedAt     time.Time  `json:"createdAt"`
	BuyDate       time.Time  `json:"buyDate"`
}
