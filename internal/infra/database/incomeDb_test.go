package database

import (
	"context"
	"testing"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
	"github.com/stretchr/testify/assert"
)

type incomeId string

var idIncome incomeId = incomeId("incomeId")

func TestCreateIncome(t *testing.T) {
	inc, err := entity.NewIncome("0194dc4c-39cf-7b69-8f16-5ee8daa95aa2", 4500.00)

	assert.NoError(t, err)
	assert.Equal(t, 4500.00, inc.Value)

	incDb := NewIncomeDB(db)

	status, err := incDb.CreateIncome(inc)

	assert.NoError(t, err)
	assert.Equal(t, 201, status)

	testsCtx = context.WithValue(defaultContext, idIncome, inc.Id.String())
}

func TestGetIncomeValueByuserId(t *testing.T) {
	incDb := NewIncomeDB(db)

	value, status, err := incDb.GetIncomeValueByUserId("0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, 4500.00, value)
}

func TestUpdateIncomeValueByuserId(t *testing.T) {
	incDb := NewIncomeDB(db)

	status, err := incDb.UpdateIncome("0194dc4c-39cf-7b69-8f16-5ee8daa95aa2", 2000.00)

	assert.NoError(t, err)
	assert.Equal(t, 200, status)

	value, status, err := incDb.GetIncomeValueByUserId("0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, 2000.00, value)

}

func TestGetIncomeValueWithInvaliduserId(t *testing.T) {
	incDb := NewIncomeDB(db)

	value, status, err := incDb.GetIncomeValueByUserId("0194dc4c-39cf-7b69-8f16-5ee8daa95aa")

	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Equal(t, 0.0, value)
}

func TestDeleteIncomeById(t *testing.T) {
	incDb := NewIncomeDB(db)
	incomeId := testsCtx.Value(idIncome).(string)

	status, err := incDb.DeleteIncome(incomeId)

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
}
