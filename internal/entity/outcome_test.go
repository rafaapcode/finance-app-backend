package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOutcomeWithExpireDate(t *testing.T) {
	outcome, err := NewOutcome("Fixo", "Home", "Pix", "01949d63-8f64-7a1e-8b9b-f62344689113", 5.00, true, 1)
	assert.Nil(t, err)

	assert.Equal(t, "Fixo", outcome.Type)
	assert.Equal(t, "Home", outcome.Category)
	assert.Equal(t, "Pix", outcome.PaymentMethod)
	assert.Equal(t, "01949d63-8f64-7a1e-8b9b-f62344689113", outcome.Userid)
	assert.Equal(t, 5.00, outcome.Value)
	assert.True(t, outcome.Notification)
}

func TestCreateOutcomeWithoutExpireDate(t *testing.T) {

	outcome, err := NewOutcome("Fixo", "Home", "Pix", "01949d63-8f64-7a1e-8b9b-f62344689113", 5.00, false, 0)
	assert.Nil(t, err)

	assert.Equal(t, "Fixo", outcome.Type)
	assert.Equal(t, "Home", outcome.Category)
	assert.Equal(t, "Pix", outcome.PaymentMethod)
	assert.Equal(t, "01949d63-8f64-7a1e-8b9b-f62344689113", outcome.Userid)
	assert.Equal(t, 5.00, outcome.Value)
	assert.False(t, outcome.Notification)
	assert.Equal(t, 0, outcome.ExpireDate)
}

func TestOutcomeUserIdIsRequired(t *testing.T) {
	outcome, err := NewOutcome("Fixo", "Home", "Pix", "", 5.00, false, 1)
	assert.Nil(t, err)

	err = outcome.Validate()

	assert.Equal(t, ErrUserIdIsRequired, err)
}

func TestOutcomeUserIdIsInvalid(t *testing.T) {
	outcome, err := NewOutcome("Fixo", "Home", "Pix", "01949d63-8f64-7a1e8b9b-f62344689113", 5.00, false, 1)
	assert.Nil(t, err)

	err = outcome.Validate()

	assert.Equal(t, ErrUserIdIsInvalid, err)
}

func TestOutcomeTypeRequired(t *testing.T) {
	outcome, err := NewOutcome("", "Home", "Pix", "01949d63-8f64-7a1e-8b9b-f62344689113", 5.00, false, 1)
	assert.Nil(t, err)

	err = outcome.Validate()

	assert.Equal(t, ErrTypeIsRequired, err)
}

func TestOutcomeCategoryRequired(t *testing.T) {
	outcome, err := NewOutcome("Fixo", "", "Pix", "01949d63-8f64-7a1e-8b9b-f62344689113", 5.00, false, 1)
	assert.Nil(t, err)

	err = outcome.Validate()

	assert.Equal(t, ErrCategoryIsRequired, err)
}

func TestOutcomePaymentMethodRequired(t *testing.T) {
	outcome, err := NewOutcome("Fixo", "Home", "", "01949d63-8f64-7a1e-8b9b-f62344689113", 5.00, false, 1)
	assert.Nil(t, err)

	err = outcome.Validate()

	assert.Equal(t, ErrPaymentMethodIsRequired, err)
}

func TestOutcomePaymentMethodInvalid(t *testing.T) {
	outcome, err := NewOutcome("Fixo", "Home", "Pi", "01949d63-8f64-7a1e-8b9b-f62344689113", 5.00, false, 1)
	assert.Nil(t, err)

	err = outcome.Validate()

	assert.Equal(t, ErrPaymentMethodIsInvalid, err)
}

func TestOutcomeValueInvalid(t *testing.T) {
	outcome, err := NewOutcome("Fixo", "Home", "Pix", "01949d63-8f64-7a1e-8b9b-f62344689113", -5.00, false, 1)
	assert.Nil(t, err)

	err = outcome.Validate()

	assert.Equal(t, ErrValueIsInvalid, err)
}
