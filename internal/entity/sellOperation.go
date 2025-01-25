package entity

import (
	"time"

	"github.com/rafaapcode/finance-app-backend/pkg"
)

type SellOperation struct {
	Id            pkg.ID     `json:"id"`
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

func NewSellOperation(invesmentId, category, stockCode string, quantity int, value, sellPrice float64, sellDate time.Time) (*SellOperation, error) {
	id, err := pkg.NewUUID()

	if err != nil {
		return nil, err
	}

	return &SellOperation{
		Id:            id,
		InvestimentId: invesmentId,
		Category:      category,
		StockCode:     stockCode,
		Quantity:      quantity,
		Value:         value,
		SellPrice:     sellPrice,
		SellDate:      sellDate,
		CreatedAt:     time.Now(),
	}, nil
}
