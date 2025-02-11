package database

import (
	"context"
	"testing"
	"time"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
	"github.com/stretchr/testify/assert"
)

type supplyId string

var idSupply supplyId = supplyId("supplyId")

func TestCreateSupplyOperation(t *testing.T) {
	supply, err := entity.NewSupplyOperation("0194e2b8-6a48-7847-9e00-04df301b736a", "Ações", "RBN", 10, 12.00, 120.00, time.Date(2025, 01, 10, 0, 0, 0, 0, time.UTC))
	assert.NoError(t, err)
	assert.Equal(t, 10, supply.Quantity)

	supplyDb := NewSupplyOperationDB(db)

	status, err := supplyDb.CreateSupplyOperation(supply)

	assert.NoError(t, err)
	assert.Equal(t, 201, status)

	testsCtx = context.WithValue(defaultContext, idSupply, supply.Id.String())
}

func TestGetSupplyOperationByName(t *testing.T) {
	supplyDb := NewSupplyOperationDB(db)

	supply, status, err := supplyDb.GetSupplyOperationByName("0194e2b8-6a48-7847-9e00-04df301b736a", "RBN")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, 0, len(supply))
}

func TestGetSupplyOperationByInvalidName(t *testing.T) {
	supplyDb := NewSupplyOperationDB(db)

	supply, status, err := supplyDb.GetSupplyOperationByName("0194e2b8-6a48-7847-9e00-04df301b736", "RBN")

	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(supply))
}

func TestGetSupplyOperationsLessThan(t *testing.T) {
	supplyDb := NewSupplyOperationDB(db)

	supply, status, err := supplyDb.GetSupplyOperationsLessThan("0194e2b8-6a48-7847-9e00-04df301b736a", 150.00)

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, 0, len(supply))
}

func TestGetSupplyOperationsLessThanInvalid(t *testing.T) {
	supplyDb := NewSupplyOperationDB(db)

	supply, status, err := supplyDb.GetSupplyOperationsLessThan("0194e2b8-6a48-7847-9e00-04df301b736a", 100.00)

	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(supply))
}

func TestGetSupplyOperationsLessThanInvalidId(t *testing.T) {
	supplyDb := NewSupplyOperationDB(db)

	supply, status, err := supplyDb.GetSupplyOperationsLessThan("0194e2b8-6a48-7847-9e00-04df301b736", 150.00)

	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(supply))
}

func TestGetSupplyOperationsHigherThan(t *testing.T) {
	supplyDb := NewSupplyOperationDB(db)

	supply, status, err := supplyDb.GetSupplyOperationsHigherThan("0194e2b8-6a48-7847-9e00-04df301b736a", 100.00)

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, 0, len(supply))
}

func TestGetSupplyOperationsHigherThanInvalid(t *testing.T) {
	supplyDb := NewSupplyOperationDB(db)

	supply, status, err := supplyDb.GetSupplyOperationsHigherThan("0194e2b8-6a48-7847-9e00-04df301b736a", 130.00)

	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(supply))
}

func TestGetSupplyOperationsHigherThanInvalidId(t *testing.T) {
	supplyDb := NewSupplyOperationDB(db)

	supply, status, err := supplyDb.GetSupplyOperationsHigherThan("0194e2b8-6a48-7847-9e00-04df301b736", 130.00)

	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(supply))
}

func TestGetSupplyOperationsByCategory(t *testing.T) {
	supplyDb := NewSupplyOperationDB(db)

	supply, status, err := supplyDb.GetSupplyOperationsByCategory("0194e2b8-6a48-7847-9e00-04df301b736a", "Ações")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, 0, len(supply))
}

func TestGetSupplyOperationsByCategoryInvalidId(t *testing.T) {
	supplyDb := NewSupplyOperationDB(db)

	supply, status, err := supplyDb.GetSupplyOperationsByCategory("0194e2b8-6a48-7847-9e00-04df301b736", "Ações")

	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(supply))
}

func TestGetSupplyOperationsByCategoryInvalid(t *testing.T) {
	supplyDb := NewSupplyOperationDB(db)

	supply, status, err := supplyDb.GetSupplyOperationsByCategory("0194e2b8-6a48-7847-9e00-04df301b736a", "Teste")

	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(supply))
}

func TestListAllSupplyOperation(t *testing.T) {
	supplyDb := NewSupplyOperationDB(db)

	supply, status, err := supplyDb.ListAllSupplyOperation("0194e2b8-6a48-7847-9e00-04df301b736a")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, 0, len(supply))
}

func TestListAllSupplyOperationInvalid(t *testing.T) {
	supplyDb := NewSupplyOperationDB(db)

	supply, status, err := supplyDb.ListAllSupplyOperation("0194e2b8-6a48-7847-9e00-04df301b736")

	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(supply))
}
