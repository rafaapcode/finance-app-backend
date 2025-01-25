package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateSupplyOperation(t *testing.T) {
	supplyOp, err := NewSupplyOperation("01949d63-8f64-7a1e-8b9b-f62344689113", "Ações", "RBN", 50, 54.67, 2.733, time.Now())
	assert.Nil(t, err)

	err = supplyOp.Validate()
	assert.Nil(t, err)
	assert.Equal(t, "01949d63-8f64-7a1e-8b9b-f62344689113", supplyOp.InvestimentId)
	assert.Equal(t, "Ações", supplyOp.Category)
	assert.Equal(t, "RBN", supplyOp.StockCode)
	assert.Equal(t, 50, supplyOp.Quantity)
	assert.Equal(t, 2.733, supplyOp.Value)
	assert.Equal(t, 54.67, supplyOp.SupplyPrice)
	assert.NotNil(t, supplyOp.SupplyDate)
	assert.NotNil(t, supplyOp.CreatedAt)
}

func TestSupplyInvestmentIdIsRequired(t *testing.T) {
	supplyOp, err := NewSupplyOperation("", "Ações", "RBN", 50, 54.67, 2.733, time.Now())
	assert.Nil(t, err)

	err = supplyOp.Validate()
	assert.Equal(t, ErrInvestmentIdIsRequired, err)
}

func TestSupplyInvestmentIdIsInvalid(t *testing.T) {
	supplyOp, err := NewSupplyOperation("01949d63-8f64-7a1e-8b9bf62344689113", "Ações", "RBN", 50, 54.67, 2.733, time.Now())
	assert.Nil(t, err)

	err = supplyOp.Validate()
	assert.Equal(t, ErrInvestmentIdIsInvalid, err)
}

func TestSupplyCategoryIsRequired(t *testing.T) {
	supplyOp, err := NewSupplyOperation("01949d63-8f64-7a1e-8b9b-f62344689113", "", "RBN", 50, 54.67, 2.733, time.Now())
	assert.Nil(t, err)

	err = supplyOp.Validate()
	assert.Equal(t, ErrCategoryIsRequired, err)
}

func TestSupplyStockCodeIsRequired(t *testing.T) {
	supplyOp, err := NewSupplyOperation("01949d63-8f64-7a1e-8b9b-f62344689113", "Ações", "", 50, 54.67, 2.733, time.Now())
	assert.Nil(t, err)

	err = supplyOp.Validate()
	assert.Equal(t, ErrStockCodeIsRequired, err)
}

func TestSupplyQuantityIsInvalid(t *testing.T) {
	supplyOp, err := NewSupplyOperation("01949d63-8f64-7a1e-8b9b-f62344689113", "Ações", "RBN", -10, 54.67, 2.733, time.Now())
	assert.Nil(t, err)

	err = supplyOp.Validate()
	assert.Equal(t, ErrQuantityIsInvalid, err)
}

func TestSupplySupplyPriceIsInvalid(t *testing.T) {
	supplyOp, err := NewSupplyOperation("01949d63-8f64-7a1e-8b9b-f62344689113", "Ações", "RBN", 10, -54.67, 2.733, time.Now())
	assert.Nil(t, err)

	err = supplyOp.Validate()
	assert.Equal(t, ErrSupplyPriceIsInvalid, err)
}

func TestSupplyValueIsInvalid(t *testing.T) {
	supplyOp, err := NewSupplyOperation("01949d63-8f64-7a1e-8b9b-f62344689113", "Ações", "RBN", 10, 54.67, -2.733, time.Now())
	assert.Nil(t, err)

	err = supplyOp.Validate()
	assert.Equal(t, ErrValueIsInvalid, err)
}
