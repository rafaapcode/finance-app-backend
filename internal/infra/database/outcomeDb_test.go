package database

import (
	"context"
	"testing"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
	"github.com/stretchr/testify/assert"
)

type outcomeId string

var idOutcome outcomeId = outcomeId("outcomeId")

func TestCreateOutcome(t *testing.T) {
	outcome, err := entity.NewOutcome("fixo", "assinatura", "Cartão de Crédito", "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2", 40.00, true, 5)
	assert.NoError(t, err)

	outcomeDb := NewOutcomeDb(db)

	status, err := outcomeDb.CreateOutcome(outcome)

	assert.NoError(t, err)
	assert.Equal(t, 201, status)
	testsCtx = context.WithValue(defaultContext, idOutcome, outcome.Id.String())
}

func TestGetOutcomeById(t *testing.T) {
	outcomeDb := NewOutcomeDb(db)
	outcomeId := testsCtx.Value(idOutcome).(string)

	outcome, status, err := outcomeDb.GetOutcomeById(outcomeId)

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, 40.00, outcome.Value)
	assert.Equal(t, "fixo", outcome.Type)
	assert.Equal(t, "assinatura", outcome.Category)
}

func TestGetAllFixedOutcome(t *testing.T) {
	outcomedb := NewOutcomeDb(db)

	outcomes, status, err := outcomedb.GetAllFixedOutcome("0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, 0, len(outcomes))
}

func TestGetAllFixedOutcomeWithInvalidUserId(t *testing.T) {
	outcomedb := NewOutcomeDb(db)

	outcomes, status, err := outcomedb.GetAllFixedOutcome("0194dc4c-39cf-7b69-8f16-5ee8daa95aa")

	assert.Error(t, err)
	assert.Equal(t, "nenhuma saída encontrada", err.Error())
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(outcomes))
}

func TestGetOutcomeByCategory(t *testing.T) {
	outcomedb := NewOutcomeDb(db)

	outcomes, status, err := outcomedb.GetAllOutcomeByCategory("assinatura", "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, 0, len(outcomes))
}

func TestGetOutcomeByInvalidCategory(t *testing.T) {
	outcomedb := NewOutcomeDb(db)

	outcomes, status, err := outcomedb.GetAllOutcomeByCategory("aluguel", "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.Error(t, err)
	assert.Equal(t, "nenhuma saída encontrada", err.Error())
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(outcomes))
}

func TestGetOutcomeByPaymentMethod(t *testing.T) {
	outcomedb := NewOutcomeDb(db)

	outcomes, status, err := outcomedb.GetAllOutcomeByPaymentMethod("Cartão de Crédito", "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, 0, len(outcomes))
}

func TestGetOutcomeByInvalidPaymentMethod(t *testing.T) {
	outcomedb := NewOutcomeDb(db)

	outcomes, status, err := outcomedb.GetAllOutcomeByPaymentMethod("Cart de Crédito", "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.Error(t, err)
	assert.Equal(t, "nenhuma saída encontrada", err.Error())
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(outcomes))
}

func TestGetOutcomeByType(t *testing.T) {
	outcomedb := NewOutcomeDb(db)

	outcomes, status, err := outcomedb.GetAllOutcomeByType("fixo", "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, 0, len(outcomes))
}

func TestGetOutcomeByInvalidType(t *testing.T) {
	outcomedb := NewOutcomeDb(db)

	outcomes, status, err := outcomedb.GetAllOutcomeByType("fio", "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.Error(t, err)
	assert.Equal(t, "nenhuma saída encontrada", err.Error())
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(outcomes))
}

func TestGetOutcomeLessThan(t *testing.T) {
	outcomedb := NewOutcomeDb(db)

	outcomes, status, err := outcomedb.GetOutcomeLessThan(50.00, "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, 0, len(outcomes))
}

func TestGetOutcomeLessThanInvalid(t *testing.T) {
	outcomedb := NewOutcomeDb(db)

	outcomes, status, err := outcomedb.GetOutcomeLessThan(30.00, "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.Error(t, err)
	assert.Equal(t, "nenhuma saída encontrada", err.Error())
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(outcomes))
}

func TestGetOutcomeHigherThan(t *testing.T) {
	outcomedb := NewOutcomeDb(db)

	outcomes, status, err := outcomedb.GetOutcomeHigherThan(30.00, "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, 0, len(outcomes))
}

func TestGetOutcomeHigherThanInvalid(t *testing.T) {
	outcomedb := NewOutcomeDb(db)

	outcomes, status, err := outcomedb.GetOutcomeHigherThan(40.00, "0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.Error(t, err)
	assert.Equal(t, "nenhuma saída encontrada", err.Error())
	assert.Equal(t, 404, status)
	assert.Equal(t, 0, len(outcomes))
}
