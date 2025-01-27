package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGoals(t *testing.T) {
	goals, err := NewGoals("0194a9f9-2d99-7adb-8946-edffe1ff5d21", "HOME", 0.5)

	assert.Nil(t, err)
	assert.NotNil(t, goals.Id)
	assert.Equal(t, "0194a9f9-2d99-7adb-8946-edffe1ff5d21", goals.Userid)
	assert.Equal(t, "HOME", goals.Category)
	assert.Equal(t, 0.5, goals.Percentage)
	assert.NotNil(t, goals.CreatedAt)
}

func TestGoalsUserIdIsRequired(t *testing.T) {
	goals, err := NewGoals("", "HOME", 0.5)
	assert.Nil(t, err)

	err = goals.Validate()

	assert.Equal(t, ErrUserIdIsRequired, err)
}

func TestGoalsUserIdIsInvalid(t *testing.T) {
	goals, err := NewGoals("0194a9f92d99-7adb-8946-edffe1ff5d21", "HOME", 0.5)
	assert.Nil(t, err)

	err = goals.Validate()

	assert.Equal(t, ErrUserIdIsInvalid, err)
}

func TestGoalsCategoryIsRequired(t *testing.T) {
	goals, err := NewGoals("0194a9f9-2d99-7adb-8946-edffe1ff5d21", "", 0.5)
	assert.Nil(t, err)

	err = goals.Validate()

	assert.Equal(t, ErrCategoryIsRequired, err)
}

func TestGoalsPercentageIsInvalid(t *testing.T) {
	goals, err := NewGoals("0194a9f9-2d99-7adb-8946-edffe1ff5d21", "Home", -0.5)
	assert.Nil(t, err)

	err = goals.Validate()

	assert.Equal(t, ErrPercentageIsInvalid, err)
}
