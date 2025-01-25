package entity

import (
	"time"

	"github.com/rafaapcode/finance-app-backend/pkg"
)

type BuyOperation struct {
	Id            pkg.ID     `json:"id"`
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

func NewBuyOperation(investmentId, category, stockCode string, quantity int, buyPrice, value float64, buyDate time.Time) (*BuyOperation, error) {
	id, err := pkg.NewUUID()

	if err != nil {
		return nil, err
	}

	return &BuyOperation{
		Id:            id,
		InvestimentId: investmentId,
		Category:      category,
		StockCode:     stockCode,
		Quantity:      quantity,
		BuyPrice:      buyPrice,
		Value:         value,
		BuyDate:       buyDate,
		CreatedAt:     time.Now(),
	}, nil
}
