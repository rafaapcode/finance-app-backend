package database

import (
	"context"
	"testing"
	"time"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
	"github.com/stretchr/testify/assert"
)

type buyOpId string

var idBuyOp buyOpId = buyOpId("buyoperationId")

func TestCreateBuyOperation(t *testing.T) {
	buyOp, err := entity.NewBuyOperation("0194e2b8-6a48-7847-9e00-04df301b736a", "Ações", "RBN", 10, 10.00, 21.00, time.Date(2025, 01, 10, 0, 0, 0, 0, time.UTC))
	assert.NoError(t, err)
	assert.Equal(t, "RBN", buyOp.StockCode)

	buyOpDb := NewBuyOperationDB(db)

	status, err := buyOpDb.CreateBuyOperation(buyOp)

	assert.NoError(t, err)
	assert.Equal(t, 201, status)
	testsCtx = context.WithValue(defaultContext, idBuyOp, buyOp.Id.String())
}

func TestGetBuyOpByName(t *testing.T) {
	buyOpDb := NewBuyOperationDB(db)

	buyOp, status, err := buyOpDb.GetBuyOperationByName("0194e2b8-6a48-7847-9e00-04df301b736a", "RBN")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, "RBN", buyOp.StockCode)
}

func TestGetBuyOpByInvalidName(t *testing.T) {
	buyOpDb := NewBuyOperationDB(db)

	_, status, err := buyOpDb.GetBuyOperationByName("0194e2b8-6a48-7847-9e00-04df301b736a", "RNB")

	assert.Error(t, err)
	assert.Equal(t, 404, status)
}

func TestGetBuyOpLessThan(t *testing.T) {
	buyOpDb := NewBuyOperationDB(db)

	buyOp, status, err := buyOpDb.GetBuyOperationsLessThan("0194e2b8-6a48-7847-9e00-04df301b736a", 25.00)

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, 0, len(buyOp))
}

func TestGetBuyOpLessThanInvalid(t *testing.T) {
	buyOpDb := NewBuyOperationDB(db)

	buyop, status, err := buyOpDb.GetBuyOperationsLessThan("0194e2b8-6a48-7847-9e00-04df301b736a", 1.00)

	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(buyop))
}

func TestGetBuyOpHigherThan(t *testing.T) {
	buyOpDb := NewBuyOperationDB(db)

	buyOp, status, err := buyOpDb.GetBuyOperationsHigherThan("0194e2b8-6a48-7847-9e00-04df301b736a", 20.00)

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, 0, len(buyOp))
}

func TestGetBuyOpHigherThanInvalid(t *testing.T) {
	buyOpDb := NewBuyOperationDB(db)

	buyop, status, err := buyOpDb.GetBuyOperationsHigherThan("0194e2b8-6a48-7847-9e00-04df301b736a", 25.00)

	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(buyop))
}

func TestGetBuyOpByCategory(t *testing.T) {
	buyOpDb := NewBuyOperationDB(db)

	buyop, status, err := buyOpDb.GetBuyOperationsByCategory("0194e2b8-6a48-7847-9e00-04df301b736a", "Ações")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, 0, len(buyop))
}

func TestGetBuyOpByCategoryInvalid(t *testing.T) {
	buyOpDb := NewBuyOperationDB(db)

	buyop, status, err := buyOpDb.GetBuyOperationsByCategory("0194e2b8-6a48-7847-9e00-04df301b736a", "Teste")

	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(buyop))
}

func TestGetAllBuyOp(t *testing.T) {
	buyOpDb := NewBuyOperationDB(db)

	buyop, status, err := buyOpDb.ListAllBuyOperation("0194e2b8-6a48-7847-9e00-04df301b736a")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, 0, len(buyop))
}

func TestGetAllBuyOpInvalid(t *testing.T) {
	buyOpDb := NewBuyOperationDB(db)

	buyop, status, err := buyOpDb.ListAllBuyOperation("0194e2b8-6a48-7847-9e00-04df301b736")

	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(buyop))
}
