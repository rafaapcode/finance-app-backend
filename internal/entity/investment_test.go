package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateInvestment(t *testing.T) {
	investment, err := NewInvestment("FIIS", "01949d63-8f64-7a1e-8b9b-f62344689113", "RBN", 30, 54.00, 0.0, 1000.0, 0.0, 0.5, time.Now())

	assert.Nil(t, err)

	var defaultDate time.Time

	assert.Equal(t, "FIIS", investment.Category)
	assert.Equal(t, "01949d63-8f64-7a1e-8b9b-f62344689113", investment.Userid)
	assert.Equal(t, "RBN", investment.StockCode)
	assert.Equal(t, 30, investment.TotalQuantity)
	assert.Equal(t, 54.00, investment.BuyPrice)
	assert.Equal(t, 0.0, investment.SellPrice)
	assert.Equal(t, 1000.0, investment.Value)
	assert.Equal(t, 0.0, investment.Profit)
	assert.Equal(t, float32(0.5), investment.Percentage)
	assert.NotNil(t, investment.CreatedAt)
	assert.Equal(t, defaultDate, investment.SellDate)
}

func TestInvestmentUserIdIsRequired(t *testing.T) {
	investment, err := NewInvestment("FIIS", "", "RBN", 30, 54.00, 0.0, 1000.0, 0.0, 0.5, time.Now())
	assert.Nil(t, err)

	err = investment.Validate()

	assert.Equal(t, ErrUserIdIsRequired, err)
}

func TestInvestmentUserIdIsInvalid(t *testing.T) {
	investment, err := NewInvestment("FIIS", "01949d63-8f647a1e-8b9b-f62344689113", "RBN", 30, 54.00, 0.0, 1000.0, 0.0, 0.5, time.Now())
	assert.Nil(t, err)

	err = investment.Validate()

	assert.Equal(t, ErrUserIdIsInvalid, err)
}

func TestINvestmentCategoryIsRequired(t *testing.T) {
	investment, err := NewInvestment("", "01949d63-8f64-7a1e-8b9b-f62344689113", "RBN", 30, 54.00, 0.0, 1000.0, 0.0, 0.5, time.Now())

	assert.Nil(t, err)

	err = investment.Validate()

	assert.Equal(t, ErrCategoryIsRequired, err)
}

func TestINvestmentStockCodeIsRequired(t *testing.T) {
	investment, err := NewInvestment("FIIS", "01949d63-8f64-7a1e-8b9b-f62344689113", "", 30, 54.00, 0.0, 1000.0, 0.0, 0.5, time.Now())

	assert.Nil(t, err)

	err = investment.Validate()

	assert.Equal(t, ErrStockCodeIsRequired, err)
}

func TestINvestmentTotalQuantityIsInvalid(t *testing.T) {
	investment, err := NewInvestment("FIIS", "01949d63-8f64-7a1e-8b9b-f62344689113", "RBN", -30, 54.00, 0.0, 1000.0, 0.0, 0.5, time.Now())

	assert.Nil(t, err)

	err = investment.Validate()

	assert.Equal(t, ErrQuantityIsInvalid, err)
}

func TestINvestmentBuyPriceIsInvalid(t *testing.T) {
	investment, err := NewInvestment("FIIS", "01949d63-8f64-7a1e-8b9b-f62344689113", "RBN", 30, -54.00, 0.0, 1000.0, 0.0, 0.5, time.Now())

	assert.Nil(t, err)

	err = investment.Validate()

	assert.Equal(t, ErrBuyPriceIsInvalid, err)
}

func TestINvestmentSellPriceIsInvalid(t *testing.T) {
	investment, err := NewInvestment("FIIS", "01949d63-8f64-7a1e-8b9b-f62344689113", "RBN", 30, 54.00, -1.0, 1000.0, 0.0, 0.5, time.Now())

	assert.Nil(t, err)

	err = investment.Validate()

	assert.Equal(t, ErrSellPriceIsInvalid, err)
}

func TestINvestmentValueIsInvalid(t *testing.T) {
	investment, err := NewInvestment("FIIS", "01949d63-8f64-7a1e-8b9b-f62344689113", "RBN", 30, 54.00, 0.0, -1000.0, 0.0, 0.5, time.Now())

	assert.Nil(t, err)

	err = investment.Validate()

	assert.Equal(t, ErrValueIsInvalid, err)
}

func TestINvestmentPercentageIsNegativeInvalid(t *testing.T) {
	investment, err := NewInvestment("FIIS", "01949d63-8f64-7a1e-8b9b-f62344689113", "RBN", 30, 54.00, 0.0, 1000.0, 0.0, -0.5, time.Now())

	assert.Nil(t, err)

	err = investment.Validate()

	assert.Equal(t, ErrPercentageIsInvalid, err)
}

func TestINvestmentPercentageIsHigherInvalid(t *testing.T) {
	investment, err := NewInvestment("FIIS", "01949d63-8f64-7a1e-8b9b-f62344689113", "RBN", 30, 54.00, 0.0, 1000.0, 0.0, 1.1, time.Now())

	assert.Nil(t, err)

	err = investment.Validate()

	assert.Equal(t, ErrPercentageIsInvalid, err)
}
