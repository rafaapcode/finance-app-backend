package entity

import (
	"time"

	"github.com/rafaapcode/finance-app-backend/pkg"
)

type Investment struct {
	Id              pkg.ID            `json:"id"`
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

func NewInvestment(category, userId, stockCode string, totalQuantity int, buyPrice, sellPrice, value, profit float64, percentage float32, buyDate time.Time) (*Investment, error) {
	id, err := pkg.NewUUID()

	if err != nil {
		return nil, err
	}

	return &Investment{
		Id:            id,
		Category:      category,
		Userid:        userId,
		StockCode:     stockCode,
		TotalQuantity: totalQuantity,
		BuyPrice:      buyPrice,
		SellPrice:     sellPrice,
		Value:         value,
		Profit:        profit,
		Percentage:    percentage,
		BuyDate:       buyDate,
	}, nil

}
