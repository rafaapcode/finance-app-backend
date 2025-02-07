package database

import (
	"context"
	"testing"
	"time"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
	"github.com/stretchr/testify/assert"
)

type investmentId string

var idInvestment investmentId = investmentId("investmentId")

func TestCreateInvestment(t *testing.T) {
	inv, err := entity.NewInvestment("Ações", "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2", "RBN", 5, 20.00, 0.0, 100.00, 0.0, 0.3, time.Date(2025, 01, 10, 0, 0, 0, 0, time.UTC))

	assert.NoError(t, err)

	invDb := NewInvestmentDB(db)

	status, err := invDb.CreateInvestment(inv)

	assert.NoError(t, err)
	assert.Equal(t, 201, status)
	testsCtx = context.WithValue(defaultContext, idInvestment, inv.Id.String())
}

func TestGetInvestmentById(t *testing.T) {
	invDb := NewInvestmentDB(db)
	invId := testsCtx.Value(idInvestment).(string)

	inv, status, err := invDb.GetInvestmentById(invId)

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEmpty(t, inv)
}

func TestGetTotalOfInvestment(t *testing.T) {
	invDb := NewInvestmentDB(db)

	total, status, err := invDb.GetTotalOfInvestment("0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, 0, total)
}

func TestGetInvestmentByName(t *testing.T) {
	invDb := NewInvestmentDB(db)

	inv, status, err := invDb.GetInvestmentByName("RBN", "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, "RBN", inv.StockCode)
	assert.Equal(t, "Ações", inv.Category)
}

func TestGetInvestmentByInvalidName(t *testing.T) {
	invDb := NewInvestmentDB(db)

	inv, status, err := invDb.GetInvestmentByName("RNB", "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Empty(t, inv)
}

func TestGetDiversificationPortfolio(t *testing.T) {
	invDb := NewInvestmentDB(db)

	metrics, status, err := invDb.GetPortfolioDiversification("0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")
	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEmpty(t, metrics.Metrics["Ações"])
}

func TestUpdateInvestment(t *testing.T) {
	invDb := NewInvestmentDB(db)
	invId := testsCtx.Value(idInvestment).(string)

	inv, status, err := invDb.GetInvestmentById(invId)
	assert.NoError(t, err)
	assert.Equal(t, 200, status)

	inv.Profit = 10.10
	inv.Percentage = 0.5

	status, err = invDb.UpdateInvestment(&inv)

	assert.NoError(t, err)
	assert.Equal(t, 200, status)

	inv, _, _ = invDb.GetInvestmentById(invId)

	assert.Equal(t, 10.10, inv.Profit)
}

func TestGetFirstPageOfAllInvestment(t *testing.T) {
	invDb := NewInvestmentDB(db)

	inv, status, err := invDb.GetAllOfInvestment("0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, 5, len(inv))
}

func TestGetFirstPageOfAllInvestmentInvalidUserid(t *testing.T) {
	invDb := NewInvestmentDB(db)

	inv, status, err := invDb.GetAllOfInvestment("0194dc4c-39cf-7b69-8f16-5ee8daa95aa")

	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(inv))
}

func TestGetFirstPageOfAllInvestmentByCategory(t *testing.T) {
	invDb := NewInvestmentDB(db)

	inv, status, err := invDb.GetInvestmentByCategory("Ações", "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, 5, len(inv))
}

func TestGetFirstPageOfAllInvestmentInvalidUseridByCategory(t *testing.T) {
	invDb := NewInvestmentDB(db)

	inv, status, err := invDb.GetInvestmentByCategory("Acoes", "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(inv))
}

func TestGetNextPageOfAllInvestment(t *testing.T) {
	invDb := NewInvestmentDB(db)

	inv, status, err := invDb.GetAllOfInvestment("0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, 5, len(inv))

	nextInv, status, err := invDb.GetNextPageAllOfInvestment(inv[4].Id.String(), "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, 5, len(nextInv))
	assert.NotEqual(t, nextInv[0].Id.String(), inv[0].Id.String())
}

func TestGetPreviousPageOfAllInvestment(t *testing.T) {
	invDb := NewInvestmentDB(db)

	inv, status, err := invDb.GetAllOfInvestment("0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, 5, len(inv))

	nextInv, status, err := invDb.GetNextPageAllOfInvestment(inv[4].Id.String(), "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, 5, len(nextInv))
	assert.NotEqual(t, nextInv[0].Id.String(), inv[0].Id.String())

	prevInv, status, err := invDb.GetPreviousPageAllOfInvestment(nextInv[0].Id.String(), "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, 5, len(nextInv))
	assert.Equal(t, prevInv[0].Id.String(), inv[0].Id.String())
	assert.NotEqual(t, prevInv[0].Id.String(), nextInv[0].Id.String())
}

func TestGetNextPageOfAllInvestmentByCategory(t *testing.T) {
	invDb := NewInvestmentDB(db)

	inv, status, err := invDb.GetInvestmentByCategory("Ações", "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, 5, len(inv))

	nextInv, status, err := invDb.GetNextPageInvestmentByCategory("Ações", "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2", inv[4].Id.String())

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, 5, len(inv))
	assert.NotEqual(t, nextInv[0].Id.String(), inv[0].Id.String())
}

func TestGetPreviousPageOfAllInvestmentByCategory(t *testing.T) {
	invDb := NewInvestmentDB(db)

	inv, status, err := invDb.GetInvestmentByCategory("Ações", "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, 5, len(inv))

	nextInv, status, err := invDb.GetNextPageInvestmentByCategory("Ações", "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2", inv[4].Id.String())

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, 5, len(inv))
	assert.NotEqual(t, nextInv[0].Id.String(), inv[0].Id.String())

	prevInv, status, err := invDb.GetPreviousPageInvestmentByCategory("Ações", "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2", nextInv[0].Id.String())

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, 5, len(nextInv))
	assert.Equal(t, prevInv[0].Id.String(), inv[0].Id.String())
	assert.NotEqual(t, prevInv[0].Id.String(), nextInv[0].Id.String())
}
