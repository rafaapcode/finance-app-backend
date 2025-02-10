package database

import (
	"context"
	"testing"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
	"github.com/stretchr/testify/assert"
)

type extraIncomeId string

var idExtraIncome extraIncomeId = extraIncomeId("extraIncomeId")

func TestCreateExtraIncome(t *testing.T) {
	extInc, err := entity.NewExtraIncome("0194dc4c-39cf-7b69-8f16-5ee8daa95aa2", "Investimentos", 500.00)

	assert.NoError(t, err)
	assert.Equal(t, 500.00, extInc.Value)

	extIncDb := NewExtraIncomeDB(db)

	status, err := extIncDb.CreateExtraIncome(extInc)

	assert.NoError(t, err)
	assert.Equal(t, 201, status)

	testsCtx = context.WithValue(defaultContext, idExtraIncome, extInc.Id.String())
}

func TestGetExtraIncomeById(t *testing.T) {
	extIncDb := NewExtraIncomeDB(db)
	extraIncId := testsCtx.Value(idExtraIncome).(string)

	extInc, status, err := extIncDb.GetExtraIncomeById(extraIncId)

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, 500.00, extInc.Value)
}

func TestDeleteExtraIncomeById(t *testing.T) {
	extIncDb := NewExtraIncomeDB(db)
	extraIncId := testsCtx.Value(idExtraIncome).(string)

	status, err := extIncDb.DeleteExtraIncome(extraIncId)

	assert.NoError(t, err)
	assert.Equal(t, 200, status)

	_, status, err = extIncDb.GetExtraIncomeById(extraIncId)

	assert.Error(t, err)
	assert.Equal(t, 404, status)
}
