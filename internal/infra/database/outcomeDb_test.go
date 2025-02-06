package database

import (
	"testing"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateOutcome(t *testing.T) {
	outcome, err := entity.NewOutcome("fixo", "assinatura", "Cartão de Crédito", "0194db72-ea40-72b0-8e5d-4a9b50519732", 40.00, true, 5)
	assert.NoError(t, err)

	outcomeDb := NewOutcomeDb(db)

	status, err := outcomeDb.CreateOutcome(outcome)

	assert.NoError(t, err)
	assert.Equal(t, 201, status)
}
