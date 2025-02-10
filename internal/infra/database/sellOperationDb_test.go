package database

import (
	"context"
	"testing"
	"time"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
	"github.com/stretchr/testify/assert"
)

type sellOpId string

var idSellOp sellOpId = sellOpId("selloperationId")

func TestCreateSellOperation(t *testing.T) {
	sellOp, err := entity.NewSellOperation("0194e2b8-6a48-7847-9e00-04df301b736a", "Ações", "RBN", 10, 21.00, 10.00, time.Date(2025, 01, 10, 0, 0, 0, 0, time.UTC))
	assert.NoError(t, err)
	assert.Equal(t, "RBN", sellOp.StockCode)

	sellOpDb := NewSellOperationDB(db)

	status, err := sellOpDb.CreateSellOperation(sellOp)

	assert.NoError(t, err)
	assert.Equal(t, 201, status)
	testsCtx = context.WithValue(defaultContext, idSellOp, sellOp.Id.String())
}

func TestGetSellOpByName(t *testing.T) {
	sellOpDb := NewSellOperationDB(db)

	sellOp, status, err := sellOpDb.GetSellOperationByName("0194e2b8-6a48-7847-9e00-04df301b736a", "RBN")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, "RBN", sellOp.StockCode)
}

func TestGetSellOpByInvalidName(t *testing.T) {
	sellOpDb := NewSellOperationDB(db)

	_, status, err := sellOpDb.GetSellOperationByName("0194e2b8-6a48-7847-9e00-04df301b736a", "RNB")

	assert.Error(t, err)
	assert.Equal(t, 404, status)
}

func TestGetSellOpLessThan(t *testing.T) {
	sellOpDb := NewSellOperationDB(db)

	sellop, status, err := sellOpDb.GetSellOperationsLessThan("0194e2b8-6a48-7847-9e00-04df301b736a", 25.00)

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, 0, len(sellop))
}

func TestGetSellOpLessThanInvalid(t *testing.T) {
	sellOpDb := NewSellOperationDB(db)

	sellop, status, err := sellOpDb.GetSellOperationsLessThan("0194e2b8-6a48-7847-9e00-04df301b736a", 15.00)

	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(sellop))
}

func TestGetSellOpHigherThan(t *testing.T) {
	sellOpDb := NewSellOperationDB(db)

	sellop, status, err := sellOpDb.GetSellOperationsHigherThan("0194e2b8-6a48-7847-9e00-04df301b736a", 20.00)

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, 0, len(sellop))
}

func TestGetSellOpHigherThanInvalid(t *testing.T) {
	sellOpDb := NewSellOperationDB(db)

	sellop, status, err := sellOpDb.GetSellOperationsHigherThan("0194e2b8-6a48-7847-9e00-04df301b736a", 25.00)

	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(sellop))
}

func TestGetSellOpByCategory(t *testing.T) {
	sellOpDb := NewSellOperationDB(db)

	sellop, status, err := sellOpDb.GetSellOperationsByCategory("0194e2b8-6a48-7847-9e00-04df301b736a", "Ações")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, 0, len(sellop))
}

func TestGetSellOpByCategoryInvalid(t *testing.T) {
	sellOpDb := NewSellOperationDB(db)

	sellop, status, err := sellOpDb.GetSellOperationsByCategory("0194e2b8-6a48-7847-9e00-04df301b736a", "Teste")

	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(sellop))
}

func TestGetAllSellOp(t *testing.T) {
	sellOpDb := NewSellOperationDB(db)

	sellop, status, err := sellOpDb.ListAllSellOperation("0194e2b8-6a48-7847-9e00-04df301b736a")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, 0, len(sellop))
}

func TestGetAllSellOpInvalid(t *testing.T) {
	sellOpDb := NewSellOperationDB(db)

	sellop, status, err := sellOpDb.ListAllSellOperation("0194e2b8-6a48-7847-9e00-04df301b736")

	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(sellop))
}
