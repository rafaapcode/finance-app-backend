package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateSellOperation(t *testing.T) {
	sellOp, err := NewSellOperation("01949d63-8f64-7a1e-8b9b-f62344689113", "Ações", "RBN", 10, 2.733, 54.67, time.Now())

	assert.Nil(t, err)

	assert.Equal(t, "01949d63-8f64-7a1e-8b9b-f62344689113", sellOp.InvestimentId)
	assert.Equal(t, "Ações", sellOp.Category)
	assert.Equal(t, "RBN", sellOp.StockCode)
	assert.Equal(t, 10, sellOp.Quantity)
	assert.Equal(t, 2.733, sellOp.Value)
	assert.Equal(t, 54.67, sellOp.SellPrice)
	assert.NotNil(t, sellOp.SellDate)
	assert.NotNil(t, sellOp.CreatedAt)
}

func TestSellOpInvestmentIdIsRequired(t *testing.T) {
	sellOp, err := NewSellOperation("", "Ações", "RBN", 50, 54.67, 2.733, time.Now())
	assert.Nil(t, err)

	err = sellOp.Validate()
	assert.Equal(t, ErrInvestmentIdIsRequired, err)
}

func TestSellOpInvestmentIdIsInvalid(t *testing.T) {
	sellOp, err := NewSellOperation("01949d63-8f64-7a1e-8b9bf62344689113", "Ações", "RBN", 50, 54.67, 2.733, time.Now())
	assert.Nil(t, err)

	err = sellOp.Validate()
	assert.Equal(t, ErrInvestmentIdIsInvalid, err)
}

func TestSellOpCategoryIsRequired(t *testing.T) {
	sellOp, err := NewSellOperation("01949d63-8f64-7a1e-8b9b-f62344689113", "", "RBN", 50, 54.67, 2.733, time.Now())
	assert.Nil(t, err)

	err = sellOp.Validate()
	assert.Equal(t, ErrCategoryIsRequired, err)
}

func TestSellOpStockCodeIsRequired(t *testing.T) {
	sellOp, err := NewSellOperation("01949d63-8f64-7a1e-8b9b-f62344689113", "Ações", "", 50, 54.67, 2.733, time.Now())
	assert.Nil(t, err)

	err = sellOp.Validate()
	assert.Equal(t, ErrStockCodeIsRequired, err)
}

func TestSellOpQuantityIsInvalid(t *testing.T) {
	sellOp, err := NewSellOperation("01949d63-8f64-7a1e-8b9b-f62344689113", "Ações", "RBN", -10, 54.67, 2.733, time.Now())
	assert.Nil(t, err)

	err = sellOp.Validate()
	assert.Equal(t, ErrQuantityIsInvalid, err)
}

func TestSellOpSellPriceIsInvalid(t *testing.T) {
	sellOp, err := NewSellOperation("01949d63-8f64-7a1e-8b9b-f62344689113", "Ações", "RBN", 10, 2.733, -54.67, time.Now())
	assert.Nil(t, err)

	err = sellOp.Validate()
	assert.Equal(t, ErrSellPriceIsInvalid, err)
}

func TestSellOpValueIsInvalid(t *testing.T) {
	sellOp, err := NewSellOperation("01949d63-8f64-7a1e-8b9b-f62344689113", "Ações", "RBN", 10, -2.733, 54.67, time.Now())
	assert.Nil(t, err)

	err = sellOp.Validate()
	assert.Equal(t, ErrValueIsInvalid, err)
}
