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
	UserDb  database.UserInterface
}

func NewGoalsHandler(goals database.GoalsInterface, userdb database.UserInterface) *GoalsHandler {
	return &GoalsHandler{
		GoalsDb: goals,
		UserDb:  userdb,
	}
}

func (goalsHand *GoalsHandler) CreateGoals(w http.ResponseWriter, r *http.Request) {
	var goals dto.CreateGoalDto
	var msgRes = pkg.NewMessageResponse("")

	err := json.NewDecoder(r.Body).Decode(&goals)
	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	_, _, err = goalsHand.UserDb.GetUser(goals.UserId)

	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	if goals.Percentage > 1 {
		msgRes.Message = "percentage must be less than 100%"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	_, status, err := goalsHand.GoalsDb.SumPercentageOfAllGoals(goals.UserId, goals.Percentage)

	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	goal, err := entity.NewGoals(goals.UserId, goals.Category, goals.Percentage)

	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}
	err = goal.Validate()
	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	status, err = goalsHand.GoalsDb.CreateGoal(goal)
	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	msgRes.Message = "Goal created successfully !"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msgRes)
}

func (goalsHand *GoalsHandler) ListAllGoals(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	var errMsg = pkg.NewMessageResponse("")

	if userid == "" {
		errMsg.Message = "UserId is required"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	allGoals, status, err := goalsHand.GoalsDb.ListAllGoals(userid)
	if err != nil {
		errMsg.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	dataRes := pkg.NewDataResponse(allGoals)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dataRes)
}

func (goalsHand *GoalsHandler) UpdateGoal(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userid := chi.URLParam(r, "userid")
	newPercentage := chi.URLParam(r, "percentage")
	var msgRes = pkg.NewMessageResponse("")

	if id == "" || newPercentage == "" {
		msgRes.Message = "Id and percentage is required"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	percentageValue, err := strconv.ParseFloat(newPercentage, 64)

	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	_, status, err := goalsHand.GoalsDb.SumPercentageForUpdateGoals(userid, id, percentageValue)

	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	status, err = goalsHand.GoalsDb.UpdateGoal(id, percentageValue)
	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	msgRes.Message = "Goal updated successfully"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msgRes)
}

func (goalsHand *GoalsHandler) DeleteGoal(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var msgRes = pkg.NewMessageResponse("")

	if id == "" {
		msgRes.Message = "Id is required"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	_, status, err := goalsHand.GoalsDb.GetGoal(id)

	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	status, err = goalsHand.GoalsDb.DeleteGoal(id)
	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	msgRes.Message = "Goal deleted successfully"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msgRes)
}
