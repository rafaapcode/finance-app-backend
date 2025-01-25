package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateOutcomeWithExpireDate(t *testing.T) {
	outcome, err := NewOutcome("Fixo", "Home", "Pix", "01949d63-8f64-7a1e-8b9b-f62344689113", 5.00, true, time.Now())
	assert.Nil(t, err)

	var defaultDate time.Time

	assert.Equal(t, "Fixo", outcome.Type)
	assert.Equal(t, "Home", outcome.Category)
	assert.Equal(t, "Pix", outcome.PaymentMethod)
	assert.Equal(t, "01949d63-8f64-7a1e-8b9b-f62344689113", outcome.Userid)
	assert.Equal(t, 5.00, outcome.Value)
	assert.True(t, outcome.Notification)
	assert.NotEqual(t, defaultDate, outcome.ExpireDate)
}

func TestCreateOutcomeWithoutExpireDate(t *testing.T) {

	outcome, err := NewOutcome("Fixo", "Home", "Pix", "01949d63-8f64-7a1e-8b9b-f62344689113", 5.00, false, time.Now())
	assert.Nil(t, err)

	var defaultDate time.Time

	assert.Equal(t, "Fixo", outcome.Type)
	assert.Equal(t, "Home", outcome.Category)
	assert.Equal(t, "Pix", outcome.PaymentMethod)
	assert.Equal(t, "01949d63-8f64-7a1e-8b9b-f62344689113", outcome.Userid)
	assert.Equal(t, 5.00, outcome.Value)
	assert.False(t, outcome.Notification)
	assert.Equal(t, defaultDate, outcome.ExpireDate)
}
