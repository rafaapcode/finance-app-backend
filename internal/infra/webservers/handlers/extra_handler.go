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
	var msgRes pkg.MessageResponse

	err := json.NewDecoder(r.Body).Decode(&extraIncome)
	if err != nil {
		msgRes = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	ext, err := entity.NewExtraIncome(extraIncome.UserId, extraIncome.Category, extraIncome.Value)

	if err != nil {
		msgRes = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}
	err = ext.Validate()
	if err != nil {
		msgRes = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	status, err := extHand.ExtraIncomeDb.CreateExtraIncome(ext)

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
		Message: "ExtraIncome created successfully !",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msgRes)
}

func (extHand *ExtraHandler) GetAllExtraIncomeOfMonth(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	month := chi.URLParam(r, "month")
	var errMsg pkg.MessageResponse
	var dataRes pkg.DataResponse

	if userid == "" || month == "" {
		errMsg = pkg.MessageResponse{
			Message: "UserId and month is required",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	monthvalue, err := strconv.Atoi(month)

	if err != nil {
		errMsg = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	if monthvalue <= 0 || monthvalue > 12 {
		errMsg = pkg.MessageResponse{
			Message: "Month must be valid",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	extraIncomes, status, err := extHand.ExtraIncomeDb.GetAllExtraIncomeOfMonth(monthvalue, userid)

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
		Data: extraIncomes,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dataRes)
}

func (extHand *ExtraHandler) GetExtraIncomeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var errMsg pkg.MessageResponse
	var dataRes pkg.DataResponse

	if id == "" {
		errMsg = pkg.MessageResponse{
			Message: "Id is required",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	extraIncome, status, err := extHand.ExtraIncomeDb.GetExtraIncomeById(id)

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
		Data: *extraIncome,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dataRes)
}

func (extHand *ExtraHandler) DeleteExtraIncome(w http.ResponseWriter, r *http.Request) {
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

	status, err := extHand.ExtraIncomeDb.DeleteExtraIncome(id)

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
		Message: "Extra income deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msgRes)
}
