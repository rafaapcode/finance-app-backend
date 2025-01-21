package model

import "time"

type SupplyOperation struct {
	Id            string     `json:"id"`
	InvestimentId string     `json:"investmentId"`
	Investment    Investment `json:"investment"`
	Category      string     `json:"category"`
	StockCode     string     `json:"stockCode"`
	Quantity      int        `json:"quantity"`
	SupplyPrice   float64    `json:"supplyPrice"`
	Value         float64    `json:"value"`
	CreatedAt     time.Time  `json:"createdAt"`
	SupplyDate    time.Time  `json:"supplyDate"`
}
