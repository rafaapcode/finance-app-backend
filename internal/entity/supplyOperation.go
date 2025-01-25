package entity

import (
	"time"

	"github.com/rafaapcode/finance-app-backend/pkg"
)

type SupplyOperation struct {
	Id            pkg.ID     `json:"id"`
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

func NewSupplyOperation(investmentId, category, stockCode string, quantity int, supplyPrice, value float64, supplyDate time.Time) (*SupplyOperation, error) {
	id, err := pkg.NewUUID()

	if err != nil {
		return nil, err
	}

	return &SupplyOperation{
		Id:            id,
		InvestimentId: investmentId,
		Category:      category,
		StockCode:     stockCode,
		Quantity:      quantity,
		SupplyPrice:   supplyPrice,
		Value:         value,
		SupplyDate:    supplyDate,
		CreatedAt:     time.Now(),
	}, nil
}

func (inv *SupplyOperation) Validate() error {
	if inv.Id.String() == "" {
		return ErrIdIdRequired
	}
	if _, err := pkg.ParseID(inv.Id.String()); err != nil {
		return ErrIdInvalidId
	}
	if inv.InvestimentId == "" {
		return ErrInvestmentIdIsRequired
	}
	if _, err := pkg.ParseID(inv.InvestimentId); err != nil {
		return ErrInvestmentIdIsInvalid
	}
	if inv.Category == "" {
		return ErrCategoryIsRequired
	}
	if inv.StockCode == "" {
		return ErrStockCodeIsRequired
	}
	if inv.Quantity < 0 {
		return ErrQuantityIsInvalid
	}
	if inv.SupplyPrice < 0 {
		return ErrSupplyPriceIsInvalid
	}
	if inv.Value < 0 {
		return ErrValueIsInvalid
	}

	return nil
}
