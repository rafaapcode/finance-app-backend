package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateExtraIncome(t *testing.T) {
	extraIncome, err := NewExtraIncome("0194a9f9-2d99-7adb-8946-edffe1ff5d21", "HOME", 0.5)

	assert.Nil(t, err)
	assert.NotNil(t, extraIncome.Id)
	assert.Equal(t, "0194a9f9-2d99-7adb-8946-edffe1ff5d21", extraIncome.Userid)
	assert.Equal(t, "HOME", extraIncome.Category)
	assert.Equal(t, 0.5, extraIncome.Value)
	assert.NotNil(t, extraIncome.Date)
}

func TestExtraIncomeUserIdIsRequired(t *testing.T) {
	extraIncome, err := NewExtraIncome("", "HOME", 0.5)
	assert.NoError(t, err)

	err = extraIncome.Validate()

	assert.Equal(t, ErrUserIdIsRequired, err)
}

func TestExtraIncomeUserIdIsInvalid(t *testing.T) {
	extraIncome, err := NewExtraIncome("0194a9f92d99-7adb-8946-edffe1ff5d21", "HOME", 0.5)
	assert.NoError(t, err)

	err = extraIncome.Validate()

	assert.Equal(t, ErrUserIdIsInvalid, err)
}

func TestExtraIncomeCategoryIsRequired(t *testing.T) {
	extraIncome, err := NewExtraIncome("0194a9f9-2d99-7adb-8946-edffe1ff5d21", "", 0.5)
	assert.NoError(t, err)

	err = extraIncome.Validate()

	assert.Equal(t, ErrCategoryIsRequired, err)
}

func TestExtraIncomeValueIsInvalid(t *testing.T) {
	extraIncome, err := NewExtraIncome("0194a9f9-2d99-7adb-8946-edffe1ff5d21", "Home", -0.5)
	assert.NoError(t, err)

	err = extraIncome.Validate()

	assert.Equal(t, ErrValueIsInvalid, err)
}
