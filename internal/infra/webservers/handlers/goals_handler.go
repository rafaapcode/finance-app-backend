package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rafaapcode/finance-app-backend/internal/dto"
	"github.com/rafaapcode/finance-app-backend/internal/entity"
	"github.com/rafaapcode/finance-app-backend/internal/infra/database"
	"github.com/rafaapcode/finance-app-backend/pkg"
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
	var msgRes pkg.MessageResponse

	err := json.NewDecoder(r.Body).Decode(&goals)
	if err != nil {
		msgRes = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	goal, err := entity.NewGoals(goals.UserId, goals.Category, goals.Percentage)

	if err != nil {
		msgRes = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}
	err = goal.Validate()
	if err != nil {
		msgRes = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	status, err := goalsHand.GoalsDb.CreateGoal(goal)
	if err != nil {
		msgRes = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	msgRes = pkg.MessageResponse{
		Message: "Goal created successfully !",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msgRes)
}

func (goalsHand *GoalsHandler) ListAllGoals(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	var errMsg pkg.MessageResponse
	var dataRes pkg.DataResponse

	if userid == "" {
		errMsg = pkg.MessageResponse{
			Message: "UserId is required",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	allGoals, status, err := goalsHand.GoalsDb.ListAllGoals(userid)
	if err != nil {
		errMsg = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	dataRes = pkg.DataResponse{
		Data: allGoals,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dataRes)
}

func (goalsHand *GoalsHandler) UpdateGoal(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	newPercentage := chi.URLParam(r, "percentage")
	var msgRes pkg.MessageResponse

	if id == "" || newPercentage == "" {
		msgRes = pkg.MessageResponse{
			Message: "Id and percentage is required",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	percentageValue, err := strconv.ParseFloat(newPercentage, 64)

	if err != nil {
		msgRes = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	status, err := goalsHand.GoalsDb.UpdateGoal(id, percentageValue)
	if err != nil {
		msgRes = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	msgRes = pkg.MessageResponse{
		Message: "Goal updated successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msgRes)
}

func (goalsHand *GoalsHandler) DeleteGoal(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var msgRes pkg.MessageResponse

	if id == "" {
		msgRes = pkg.MessageResponse{
			Message: "Id is required",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	status, err := goalsHand.GoalsDb.DeleteGoal(id)
	if err != nil {
		msgRes = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	msgRes = pkg.MessageResponse{
		Message: "Goal deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msgRes)
}
