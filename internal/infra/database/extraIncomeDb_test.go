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
	inc, err := entity.NewIncome("0194dc4c-39cf-7b69-8f16-5ee8daa95aa2", 4500.00)

	assert.NoError(t, err)
	assert.Equal(t, 4500.00, inc.Value)

	incDb := NewIncomeDB(db)

	status, err := incDb.CreateIncome(inc)

	assert.NoError(t, err)
	assert.Equal(t, 201, status)

	testsCtx = context.WithValue(defaultContext, idIncome, inc.Id.String())
}
