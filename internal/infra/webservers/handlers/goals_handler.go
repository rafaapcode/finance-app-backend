package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rafaapcode/finance-app-backend/internal/dto"
	"github.com/rafaapcode/finance-app-backend/internal/entity"
	"github.com/rafaapcode/finance-app-backend/internal/infra/database"
)

type GoalsHandler struct {
	GoalsDb database.GoalsInterface
}

func NewGoalsHandler(goals database.GoalsInterface) *GoalsHandler {
	return &GoalsHandler{
		GoalsDb: goals,
	}
}

func (goalsHand *GoalsHandler) CreateGoals(w http.ResponseWriter, r *http.Request) {
	var goals dto.CreateGoalDto

	err := json.NewDecoder(r.Body).Decode(&goals)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	goal, err := entity.NewGoals(goals.UserId, goals.Category, goals.Percentage)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = goal.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	status, err := goalsHand.GoalsDb.CreateGoal(goal)
	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Goal created successfully !",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&response)
}

func (goalsHand *GoalsHandler) ListAllGoals(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")

	if userid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	allGoals, status, err := goalsHand.GoalsDb.ListAllGoals(userid)
	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Data []entity.Goals `json:"data"`
	}{
		Data: allGoals,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)
}

func (goalsHand *GoalsHandler) UpdateGoal(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	newPercentage := chi.URLParam(r, "percentage")

	if id == "" || newPercentage == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	percentageValue, err := strconv.ParseFloat(newPercentage, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	status, err := goalsHand.GoalsDb.UpdateGoal(id, percentageValue)
	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Goal updated successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)
}

func (goalsHand *GoalsHandler) DeleteGoal(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	status, err := goalsHand.GoalsDb.DeleteGoal(id)
	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Goal deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)
}
