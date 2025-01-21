package model

import "time"

type Investment struct {
	Id              string            `json:"id"`
	Category        string            `json:"category"`
	TotalQuantity   int               `json:"totalQuantity"`
	Userid          string            `json:"userid"`
	BuyPrice        float64           `json:"buyPrice"`
	SellPrice       float64           `json:"sellPrice"`
	Percentage      float32           `json:"percentage"`
	StockCode       string            `json:"stockCode"`
	Value           float64           `json:"value"`
	Profit          float64           `json:"profit"`
	User            User              `json:"user"`
	BuyDate         time.Time         `json:"buyDate"`
	CreatedAt       time.Time         `json:"createdat"`
	LastSupplyDate  time.Time         `json:"lastSupplyDate"`
	SellDate        time.Time         `json:"sellDate"`
	BuyOperation    []BuyOperation    `json:"buyOperation"`
	SellOperation   []SellOperation   `json:"sellOperation"`
	SupplyOperation []SupplyOperation `json:"supplyOperation"`
}
