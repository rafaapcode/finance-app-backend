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

type OutcomeHandler struct {
	OutcomeDb database.OutcomeInterface
}

func NewOutcomeHandler(outcome database.OutcomeInterface) *OutcomeHandler {
	return &OutcomeHandler{
		OutcomeDb: outcome,
	}
}

func (outHand *OutcomeHandler) CreateOutcome(w http.ResponseWriter, r *http.Request) {
	var outcome dto.CreateOutcomeDto
	var res pkg.MessageResponse

	err := json.NewDecoder(r.Body).Decode(&outcome)
	if err != nil {
		res = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	out, err := entity.NewOutcome(outcome.OutcomeType, outcome.Category, outcome.PaymentMethod, outcome.UserId, outcome.Value, outcome.Notification, outcome.ExpireDate)

	if err != nil {
		res = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	err = out.Validate()

	if err != nil {
		res = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	status, err := outHand.OutcomeDb.CreateOutcome(out)
	if err != nil {
		res = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(res)
		return
	}

	res = pkg.MessageResponse{
		Message: "Outcome created successfully !",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (outHand *OutcomeHandler) GetOutcomeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var errMsg pkg.MessageResponse
	var messageData pkg.DataResponse

	if id == "" {
		errMsg = pkg.MessageResponse{
			Message: "Id is required !",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	outcome, status, err := outHand.OutcomeDb.GetOutcomeById(id)
	if err != nil {
		errMsg = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	messageData = pkg.DataResponse{
		Data: *outcome,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(messageData)
}

func (outHand *OutcomeHandler) GetAllOutcomeOfMonth(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	month := chi.URLParam(r, "month")
	var errMsg pkg.MessageResponse
	var messageData pkg.DataResponse

	if userid == "" || month == "" {
		errMsg = pkg.MessageResponse{
			Message: "Id and Month is required !",
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
			Message: "Must be a valid month",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetAllOutcomeOfMonth(monthvalue, userid)
	if err != nil {
		errMsg = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	messageData = pkg.DataResponse{
		Data: outcomes,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(messageData)
}

func (outHand *OutcomeHandler) GetAllFixedOutcome(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	var errMsg pkg.MessageResponse
	var messageData pkg.DataResponse

	if userid == "" {
		errMsg = pkg.MessageResponse{
			Message: "UserId is required !",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetAllFixedOutcome(userid)
	if err != nil {
		errMsg = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	messageData = pkg.DataResponse{
		Data: outcomes,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(messageData)
}

func (outHand *OutcomeHandler) GetAllOutcomeByCategory(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	category := chi.URLParam(r, "category")
	var errMsg pkg.MessageResponse
	var messageData pkg.DataResponse

	if userid == "" || category == "" {
		errMsg = pkg.MessageResponse{
			Message: "UserId and category is required !",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetAllOutcomeByCategory(category, userid)
	if err != nil {
		errMsg = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	messageData = pkg.DataResponse{
		Data: outcomes,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(messageData)
}

func (outHand *OutcomeHandler) GetAllOutcomeByPaymentMethod(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	paymentMethod := chi.URLParam(r, "paymentmethod")
	var errMsg pkg.MessageResponse
	var messageData pkg.DataResponse

	if userid == "" || paymentMethod == "" {
		errMsg = pkg.MessageResponse{
			Message: "UserId and PaymentMethod is required !",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetAllOutcomeByPaymentMethod(paymentMethod, userid)
	if err != nil {
		errMsg = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	messageData = pkg.DataResponse{
		Data: outcomes,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(messageData)
}

func (outHand *OutcomeHandler) GetAllOutcomeByType(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	outcomeType := chi.URLParam(r, "type")
	var errMsg pkg.MessageResponse
	var messageData pkg.DataResponse

	if userid == "" || outcomeType == "" {
		errMsg = pkg.MessageResponse{
			Message: "UserId and OutcomeType is required !",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetAllOutcomeByType(outcomeType, userid)
	if err != nil {
		errMsg = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	messageData = pkg.DataResponse{
		Data: outcomes,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(messageData)
}

func (outHand *OutcomeHandler) GetOutcomeAboutToExpire(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	daysToexpire := chi.URLParam(r, "daysToExpire")
	var errMsg pkg.MessageResponse
	var messageData pkg.DataResponse

	if userid == "" || daysToexpire == "" {
		errMsg = pkg.MessageResponse{
			Message: "UserId and DaysToExpire is required !",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	days, err := strconv.Atoi(daysToexpire)

	if err != nil {
		errMsg = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetOutcomeAboutToExpire(days, userid)
	if err != nil {
		errMsg = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	messageData = pkg.DataResponse{
		Data: outcomes,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(messageData)
}

func (outHand *OutcomeHandler) GetOutcomeLessThan(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	value := chi.URLParam(r, "value")
	var errMsg pkg.MessageResponse
	var messageData pkg.DataResponse

	if userid == "" || value == "" {
		errMsg = pkg.MessageResponse{
			Message: "UserId and value is required !",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	v, err := strconv.ParseFloat(value, 64)

	if err != nil {
		errMsg = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetOutcomeLessThan(v, userid)
	if err != nil {
		errMsg = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	messageData = pkg.DataResponse{
		Data: outcomes,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(messageData)
}

func (outHand *OutcomeHandler) GetOutcomeHigherThan(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	value := chi.URLParam(r, "value")
	var errMsg pkg.MessageResponse
	var messageData pkg.DataResponse

	if userid == "" || value == "" {
		errMsg = pkg.MessageResponse{
			Message: "UserId and value is required",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	v, err := strconv.ParseFloat(value, 64)

	if err != nil {
		errMsg = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetOutcomeHigherThan(v, userid)
	if err != nil {
		errMsg = pkg.MessageResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	messageData = pkg.DataResponse{
		Data: outcomes,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(messageData)
}

func (outHand *OutcomeHandler) DeleteOutcome(w http.ResponseWriter, r *http.Request) {
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

	msg, status, err := outHand.OutcomeDb.DeleteOutcome(id)
	if err != nil {
		msgRes = pkg.MessageResponse{
			Message: msg,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	msgRes = pkg.MessageResponse{
		Message: msg,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(msgRes)
}
