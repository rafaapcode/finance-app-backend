package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateIncome(t *testing.T) {
	income, err := NewIncome("0194a9f9-2d99-7adb-8946-edffe1ff5d21", 103.2)

	assert.Nil(t, err)
	assert.NotNil(t, income.Id)
	assert.Equal(t, "0194a9f9-2d99-7adb-8946-edffe1ff5d21", income.Userid)
	assert.Equal(t, 103.2, income.Value)
	assert.NotNil(t, income.CreatedAt)
}

func TestIncomeUserIdIsRequired(t *testing.T) {
	income, err := NewIncome("", 103.2)
	assert.Nil(t, err)

	err = income.Validate()

	assert.Equal(t, ErrUserIdIsRequired, err)
}

func TestIncomeUserIdIsInvalid(t *testing.T) {
	income, err := NewIncome("0194a9f9-2d99-7adb-8946edffe1ff5d21", 103.2)
	assert.Nil(t, err)

	err = income.Validate()

	assert.Equal(t, ErrUserIdIsInvalid, err)
}

func TestIncomeValueIsInvalid(t *testing.T) {
	income, err := NewIncome("0194a9f9-2d99-7adb-8946-edffe1ff5d21", -10.0)
	assert.Nil(t, err)

	err = income.Validate()

	assert.Equal(t, ErrValueIsInvalid, err)
}
