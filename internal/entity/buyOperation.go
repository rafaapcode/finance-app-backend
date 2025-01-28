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

func (buyOp *BuyOperation) Validate() error {
	if buyOp.Id.String() == "" {
		return ErrIdIdRequired
	}
	if _, err := pkg.ParseID(buyOp.Id.String()); err != nil {
		return ErrIdInvalidId
	}
	if buyOp.InvestimentId == "" {
		return ErrInvestmentIdIsRequired
	}
	if _, err := pkg.ParseID(buyOp.InvestimentId); err != nil {
		return ErrInvestmentIdIsInvalid
	}
	if buyOp.Category == "" {
		return ErrCategoryIsRequired
	}
	if buyOp.StockCode == "" {
		return ErrStockCodeIsRequired
	}
	if buyOp.Quantity < 0 {
		return ErrQuantityIsInvalid
	}
	if buyOp.BuyPrice < 0.0 {
		return ErrBuyPriceIsInvalid
	}
	if buyOp.Value < 0.0 || buyOp.Value != (buyOp.BuyPrice*float64(buyOp.Quantity)) {
		return ErrValueIsInvalid
	}
	return nil
}
