package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rafaapcode/finance-app-backend/internal/dto"
	"github.com/rafaapcode/finance-app-backend/internal/entity"
	"github.com/rafaapcode/finance-app-backend/internal/infra/database"
	"github.com/rafaapcode/finance-app-backend/pkg"
)

type IncomeHandler struct {
	IncomeDb      database.IncomeInterface
	ExtraIncomeDb database.ExtraIncomeInterface
}

func NewIncomeHandler(income database.IncomeInterface, extraIncomeDb database.ExtraIncomeInterface) *IncomeHandler {
	return &IncomeHandler{
		IncomeDb:      income,
		ExtraIncomeDb: extraIncomeDb,
	}
}

func (incHand *IncomeHandler) CreateIncome(w http.ResponseWriter, r *http.Request) {
	var income dto.CreateIncomeDto
	var msgRes pkg.MessageResponse

	err := json.NewDecoder(r.Body).Decode(&income)
	if err != nil {
		msgRes = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	inc, err := entity.NewIncome(income.UserId, income.Value)

	if err != nil {
		msgRes = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}
	err = inc.Validate()
	if err != nil {
		msgRes = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	status, err := incHand.IncomeDb.CreateIncome(inc)
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
		Message: "Income created successfully !",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msgRes)
}

func (incHand *IncomeHandler) GetTotalIncomeOfUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userid")
	var msgRes pkg.MessageResponse
	var msgData pkg.DataResponse

	var total float64
	if userId == "" {
		msgRes = pkg.MessageResponse{
			Message: "UserId is required",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	incomeValue, status, err := incHand.IncomeDb.GetIncomeValueByUserId(userId)

	if err != nil {
		msgRes = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}
	total = incomeValue

	currentMonth := time.Now()

	extraIncomeValue, status, err := incHand.ExtraIncomeDb.GetTotalValueOfExtracIncomeOfTheMonth(int(currentMonth.Month()), userId)

	if status == 500 {
		msgRes = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	if err == nil {
		total += extraIncomeValue
	}

	msgData = pkg.DataResponse{
		Data: total,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msgData)
}

func (incHand *IncomeHandler) DeleteIncomeById(w http.ResponseWriter, r *http.Request) {
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

	status, err := incHand.IncomeDb.DeleteIncome(id)

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
		Message: "Income deleted successfully !",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(msgRes)
}

func (incHand *IncomeHandler) UpdateIncome(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	newValue := chi.URLParam(r, "value")
	var msgRes pkg.MessageResponse

	if userid == "" {
		msgRes = pkg.MessageResponse{
			Message: "UserId is required",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	if newValue == "" {
		msgRes = pkg.MessageResponse{
			Message: "Value is required",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	valueParsed, err := strconv.ParseFloat(newValue, 64)

	if err == nil {
		msgRes = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	status, err := incHand.IncomeDb.UpdateIncome(userid, valueParsed)

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
		Message: "Income deleted successfully !",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(msgRes)
}
