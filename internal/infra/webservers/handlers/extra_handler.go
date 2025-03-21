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
	UserDb        database.UserInterface
}

func NewExtraHandler(extraIncomeDb database.ExtraIncomeInterface, userDb database.UserInterface) *ExtraHandler {
	return &ExtraHandler{
		ExtraIncomeDb: extraIncomeDb,
		UserDb:        userDb,
	}
}

func (extHand *ExtraHandler) CreateExtraIncome(w http.ResponseWriter, r *http.Request) {
	var extraIncome dto.CreateExtraIncomeDto
	var msgRes = pkg.NewMessageResponse("")

	err := json.NewDecoder(r.Body).Decode(&extraIncome)
	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	_, status, err := extHand.UserDb.GetUser(extraIncome.UserId)

	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	ext, err := entity.NewExtraIncome(extraIncome.UserId, extraIncome.Category, extraIncome.Value)

	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}
	err = ext.Validate()
	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	status, err = extHand.ExtraIncomeDb.CreateExtraIncome(ext)

	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	msgRes.Message = "ExtraIncome created successfully !"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msgRes)
}

func (extHand *ExtraHandler) GetAllExtraIncomeOfMonth(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	month := chi.URLParam(r, "month")
	var errMsg = pkg.NewMessageResponse("")

	if userid == "" || month == "" {
		errMsg.Message = "UserId and month is required"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	monthvalue, err := strconv.Atoi(month)

	if err != nil {
		errMsg.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	if monthvalue <= 0 || monthvalue > 12 {
		errMsg.Message = "Month must be valid"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	extraIncomes, status, err := extHand.ExtraIncomeDb.GetAllExtraIncomeOfMonth(monthvalue, userid)

	if err != nil {
		errMsg.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	dataRes := pkg.NewDataResponse(extraIncomes)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dataRes)
}

func (extHand *ExtraHandler) GetExtraIncomeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var errMsg = pkg.NewMessageResponse("")

	if id == "" {
		errMsg.Message = "Id is required"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	extraIncome, status, err := extHand.ExtraIncomeDb.GetExtraIncomeById(id)

	if err != nil {
		errMsg.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	dataRes := pkg.NewDataResponse(*extraIncome)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dataRes)
}

func (extHand *ExtraHandler) DeleteExtraIncome(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var msgRes = pkg.NewMessageResponse("")

	if id == "" {
		msgRes.Message = "Id is required"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	status, err := extHand.ExtraIncomeDb.DeleteExtraIncome(id)

	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	msgRes.Message = "Extra income deleted successfully"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msgRes)
}
