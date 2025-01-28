package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateBuyOperation(t *testing.T) {
	buyOp, err := NewBuyOperation("0194a9f9-2d99-7adb-8946-edffe1ff5d21", "HOME", "RBN", 10, 10.0, 1000.0, time.Now())

	assert.Nil(t, err)
	assert.NotNil(t, buyOp.Id)
	assert.Equal(t, "0194a9f9-2d99-7adb-8946-edffe1ff5d21", buyOp.InvestimentId)
	assert.Equal(t, "HOME", buyOp.Category)
	assert.Equal(t, "RBN", buyOp.StockCode)
	assert.Equal(t, 10, buyOp.Quantity)
	assert.Equal(t, 10.0, buyOp.BuyPrice)
	assert.Equal(t, 1000.0, buyOp.Value)
	assert.NotNil(t, buyOp.CreatedAt)
	assert.NotNil(t, buyOp.BuyDate)
}

func TestBuyOperationInvestmentIdIsRequired(t *testing.T) {
	buyOp, err := NewBuyOperation("", "HOME", "RBN", 10, 10.0, 1000.0, time.Now())
	assert.NoError(t, err)

	err = buyOp.Validate()

	assert.Equal(t, ErrInvestmentIdIsRequired, err)
}

func TestBuyOperationInvestmentIdIsInvalid(t *testing.T) {
	buyOp, err := NewBuyOperation("0194a9f92d99-7adb-8946-edffe1ff5d21", "HOME", "RBN", 10, 10.0, 1000.0, time.Now())
	assert.NoError(t, err)

	err = buyOp.Validate()

	assert.Equal(t, ErrInvestmentIdIsInvalid, err)
}

func TestBuyOperationCategoryIsRequired(t *testing.T) {
	buyOp, err := NewBuyOperation("0194a9f9-2d99-7adb-8946-edffe1ff5d21", "", "RBN", 10, 10.0, 1000.0, time.Now())
	assert.NoError(t, err)

	err = buyOp.Validate()

	assert.Equal(t, ErrCategoryIsRequired, err)
}

func TestBuyOperationStockCodeIsRequired(t *testing.T) {
	buyOp, err := NewBuyOperation("0194a9f9-2d99-7adb-8946-edffe1ff5d21", "Home", "", 10, 10.0, 1000.0, time.Now())
	assert.NoError(t, err)

	err = buyOp.Validate()

	assert.Equal(t, ErrStockCodeIsRequired, err)
}

func TestBuyOperationQuantityIsInvalid(t *testing.T) {
	buyOp, err := NewBuyOperation("0194a9f9-2d99-7adb-8946-edffe1ff5d21", "Home", "RBN", -10, 10.0, 1000.0, time.Now())
	assert.NoError(t, err)

	err = buyOp.Validate()

	assert.Equal(t, ErrQuantityIsInvalid, err)
}

func TestBuyOperationBuyPriceIsInvalid(t *testing.T) {
	buyOp, err := NewBuyOperation("0194a9f9-2d99-7adb-8946-edffe1ff5d21", "Home", "RBN", 10, -10.0, 1000.0, time.Now())
	assert.NoError(t, err)

	err = buyOp.Validate()

	assert.Equal(t, ErrBuyPriceIsInvalid, err)
}

func TestBuyOperationValueIsInvalid(t *testing.T) {
	buyOp, err := NewBuyOperation("0194a9f9-2d99-7adb-8946-edffe1ff5d21", "Home", "RBN", 10, 10.0, -1000.0, time.Now())
	assert.NoError(t, err)

	err = buyOp.Validate()

	assert.Equal(t, ErrValueIsInvalid, err)
}
