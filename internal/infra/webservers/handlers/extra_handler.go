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

type ExtraHandler struct {
	ExtraIncomeDb database.ExtraIncomeInterface
}

func NewExtraHandler(extraIncomeDb database.ExtraIncomeInterface) *ExtraHandler {
	return &ExtraHandler{
		ExtraIncomeDb: extraIncomeDb,
	}
}

func (extHand *ExtraHandler) CreateExtraIncome(w http.ResponseWriter, r *http.Request) {
	var extraIncome dto.CreateExtraIncomeDto

	err := json.NewDecoder(r.Body).Decode(&extraIncome)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ext, err := entity.NewExtraIncome(extraIncome.UserId, extraIncome.Category, extraIncome.Value)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = ext.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	status, err := extHand.ExtraIncomeDb.CreateExtraIncome(ext)

	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "ExtraIncome created successfully !",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&response)
}

func (extHand *ExtraHandler) GetAllExtraIncomeOfMonth(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	month := chi.URLParam(r, "month")

	if userid == "" || month == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	monthvalue, err := strconv.Atoi(month)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if monthvalue <= 0 || monthvalue > 12 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	extraIncomes, status, err := extHand.ExtraIncomeDb.GetAllExtraIncomeOfMonth(monthvalue, userid)

	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Data []entity.ExtraIncome `json:"data"`
	}{
		Data: extraIncomes,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)
}

func (extHand *ExtraHandler) GetExtraIncomeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	extraIncome, status, err := extHand.ExtraIncomeDb.GetExtraIncomeById(id)

	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Data entity.ExtraIncome `json:"data"`
	}{
		Data: *extraIncome,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)
}

func (extHand *ExtraHandler) DeleteExtraIncome(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	status, err := extHand.ExtraIncomeDb.DeleteExtraIncome(id)

	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Extra income deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)
}
