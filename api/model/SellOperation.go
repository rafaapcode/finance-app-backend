package model

import "time"

type SellOperation struct {
	Id            string     `json:"id"`
	InvestimentId string     `json:"investmentId"`
	Investment    Investment `json:"investment"`
	Category      string     `json:"category"`
	StockCode     string     `json:"stockCode"`
	Quantity      int        `json:"quantity"`
	SellPrice     float64    `json:"sellPrice"`
	Value         float64    `json:"value"`
	CreatedAt     time.Time  `json:"createdAt"`
	SellDate      time.Time  `json:"sellDate"`
}
