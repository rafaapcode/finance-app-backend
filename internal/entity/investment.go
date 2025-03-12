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
		CreatedAt:     time.Now(),
	}, nil
}

func NewUpdateSellInvestment(inv *Investment, sellPrice, profit float64, newQuantity int, newpercentage float32) *Investment {

	return &Investment{
		Id:            inv.Id,
		Category:      inv.Category,
		Userid:        inv.Userid,
		StockCode:     inv.StockCode,
		TotalQuantity: newQuantity,
		BuyPrice:      inv.BuyPrice,
		SellPrice:     sellPrice,
		Value:         inv.Value,
		Profit:        profit,
		Percentage:    newpercentage,
		BuyDate:       inv.BuyDate,
		SellDate:      time.Now(),
		CreatedAt:     inv.CreatedAt,
	}
}

func NewUpdateSupplyInvestment(inv *Investment, buyPrice, newValue float64, newQuantity int, newpercentage float32) *Investment {

	return &Investment{
		Id:             inv.Id,
		Category:       inv.Category,
		Userid:         inv.Userid,
		StockCode:      inv.StockCode,
		TotalQuantity:  newQuantity,
		BuyPrice:       buyPrice,
		SellPrice:      inv.SellPrice,
		Value:          newValue,
		Profit:         inv.Profit,
		Percentage:     newpercentage,
		BuyDate:        inv.BuyDate,
		SellDate:       inv.SellDate,
		LastSupplyDate: time.Now(),
		CreatedAt:      inv.CreatedAt,
	}
}

func (inv *Investment) Validate() error {
	if inv.Id.String() == "" {
		return ErrIdIdRequired
	}
	if _, err := pkg.ParseID(inv.Id.String()); err != nil {
		return ErrIdInvalidId
	}
	if inv.Userid == "" {
		return ErrUserIdIsRequired
	}
	if _, err := pkg.ParseID(inv.Userid); err != nil {
		return ErrUserIdIsInvalid
	}
	if inv.Category == "" {
		return ErrCategoryIsRequired
	}
	if inv.StockCode == "" {
		return ErrStockCodeIsRequired
	}
	if inv.TotalQuantity < 0 {
		return ErrQuantityIsInvalid
	}
	if inv.BuyPrice <= 0.0 {
		return ErrBuyPriceIsInvalid
	}
	if inv.SellPrice < 0.0 {
		return ErrSellPriceIsInvalid
	}
	if inv.Value < 0.0 {
		return ErrValueIsInvalid
	}
	if inv.Percentage < 0.0 || inv.Percentage > 1.0 {
		return ErrPercentageIsInvalid
	}
	return nil
}
