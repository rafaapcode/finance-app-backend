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

	err := json.NewDecoder(r.Body).Decode(&income)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	inc, err := entity.NewIncome(income.UserId, income.Value)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = inc.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	status, err := incHand.IncomeDb.CreateIncome(inc)
	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Income created successfully !",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&response)
}

func (incHand *IncomeHandler) GetTotalIncomeOfUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userid")
	var total float64
	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	incomeValue, status, err := incHand.IncomeDb.GetIncomeValueByUserId(userId)

	if err != nil {
		w.WriteHeader(status)
		return
	}
	total = incomeValue

	currentMonth := time.Now()

	extraIncomeValue, status, err := incHand.ExtraIncomeDb.GetTotalValueOfExtracIncomeOfTheMonth(int(currentMonth.Month()), userId)

	if status == 500 {
		w.WriteHeader(status)
		return
	}

	if err == nil {
		total += extraIncomeValue
	}

	response := struct {
		TotalIncomeValue float64 `json:"total_income_value"`
	}{
		TotalIncomeValue: total,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)
}

func (incHand *IncomeHandler) DeleteIncomeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	status, err := incHand.IncomeDb.DeleteIncome(id)

	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Income deleted successfully !",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&response)
}

func (incHand *IncomeHandler) UpdateIncome(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	newValue := chi.URLParam(r, "value")

	if userid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if newValue == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	valueParsed, err := strconv.ParseFloat(newValue, 64)

	if err == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	status, err := incHand.IncomeDb.UpdateIncome(userid, valueParsed)

	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Income deleted successfully !",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&response)
}
