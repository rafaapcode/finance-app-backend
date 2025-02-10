package database

import (
	"context"
	"testing"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
	"github.com/stretchr/testify/assert"
)

type goalsId string

var idGoals goalsId = goalsId("goalsId")

func TestCreateGoals(t *testing.T) {
	goals, err := entity.NewGoals("0194dc4c-39cf-7b69-8f16-5ee8daa95aa2", "Home", 0.3)

	assert.NoError(t, err)
	assert.Equal(t, 0.3, goals.Percentage)
	assert.Equal(t, "Home", goals.Category)

	goalsDb := NewGoalsDB(db)

	status, err := goalsDb.CreateGoal(goals)

	assert.NoError(t, err)
	assert.Equal(t, 201, status)

	testsCtx = context.WithValue(defaultContext, idGoals, goals.Id.String())
}

func TestUpdateGoal(t *testing.T) {
	goalsDb := NewGoalsDB(db)
	goalId := testsCtx.Value(idGoals).(string)

	status, err := goalsDb.UpdateGoal(goalId, 0.5)

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
}

func TestDeleteGoal(t *testing.T) {
	goalsDb := NewGoalsDB(db)
	goalId := testsCtx.Value(idGoals).(string)

	status, err := goalsDb.DeleteGoal(goalId)

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
}

func TestAllGoals(t *testing.T) {
	goalsDb := NewGoalsDB(db)

	goals, status, err := goalsDb.ListAllGoals("0194dc4c-39cf-7b69-8f16-5ee8daa95aa2")

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEqual(t, 0, len(goals))
}
