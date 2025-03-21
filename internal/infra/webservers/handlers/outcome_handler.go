package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

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
	var res = pkg.NewMessageResponse("")

	err := json.NewDecoder(r.Body).Decode(&outcome)
	if err != nil {
		res.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	out, err := entity.NewOutcome(strings.ToLower(outcome.OutcomeType), outcome.Category, outcome.PaymentMethod, outcome.UserId, outcome.Value, outcome.Notification, outcome.ExpireDate)
	if err != nil {
		res.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}
	err = out.Validate()

	if err != nil {
		res.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	status, err := outHand.OutcomeDb.CreateOutcome(out)
	if err != nil {
		res.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Message = "Outcome created successfully !"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (outHand *OutcomeHandler) GetOutcomeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var errMsg = pkg.NewMessageResponse("")

	if id == "" {
		errMsg.Message = "Id is required !"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	outcome, status, err := outHand.OutcomeDb.GetOutcomeById(id)
	if err != nil {
		errMsg.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	messageData := pkg.NewDataResponse(*outcome)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(messageData)
}

func (outHand *OutcomeHandler) GetAllOutcomeOfMonth(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	month := chi.URLParam(r, "month")
	var errMsg = pkg.NewMessageResponse("")

	if userid == "" || month == "" {
		errMsg.Message = "Id and Month is required !"
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
		errMsg.Message = "Must be a valid month"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetAllOutcomeOfMonth(monthvalue, userid)
	if err != nil {
		errMsg.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	if outcomes == nil {
		errMsg.Message = "Nenhuma saída encontrada para esse mês"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	messageData := pkg.NewDataResponse(outcomes)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(messageData)
}

func (outHand *OutcomeHandler) GetAllFixedOutcome(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	var errMsg = pkg.NewMessageResponse("")

	if userid == "" {
		errMsg.Message = "UserId is required !"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetAllFixedOutcome(userid)
	if err != nil {
		errMsg.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	messageData := pkg.NewDataResponse(outcomes)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(messageData)
}

func (outHand *OutcomeHandler) GetAllOutcomeByCategory(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	category := chi.URLParam(r, "category")
	errMsg := pkg.NewMessageResponse("")

	if userid == "" || category == "" {
		errMsg.Message = "UserId and category is required !"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetAllOutcomeByCategory(category, userid)
	if err != nil {
		errMsg.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	messageData := pkg.NewDataResponse(outcomes)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(messageData)
}

func (outHand *OutcomeHandler) GetAllOutcomeByPaymentMethod(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	paymentMethod := chi.URLParam(r, "paymentmethod")
	var errMsg = pkg.NewMessageResponse("")

	if userid == "" || paymentMethod == "" {
		errMsg.Message = "UserId and PaymentMethod is required !"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetAllOutcomeByPaymentMethod(paymentMethod, userid)
	if err != nil {
		errMsg.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	messageData := pkg.NewDataResponse(outcomes)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(messageData)
}

func (outHand *OutcomeHandler) GetAllOutcomeByType(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	outcomeType := chi.URLParam(r, "type")
	var errMsg = pkg.NewMessageResponse("")

	if userid == "" || outcomeType == "" {
		errMsg.Message = "UserId and OutcomeType is required !"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetAllOutcomeByType(outcomeType, userid)
	if err != nil {
		errMsg.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	messageData := pkg.NewDataResponse(outcomes)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(messageData)
}

func (outHand *OutcomeHandler) GetOutcomeAboutToExpire(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	daysToexpire := chi.URLParam(r, "daysToExpire")
	var errMsg = pkg.NewMessageResponse("")

	if userid == "" || daysToexpire == "" {
		errMsg.Message = "UserId and DaysToExpire is required !"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	days, err := strconv.Atoi(daysToexpire)

	if err != nil {
		errMsg.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetOutcomeAboutToExpire(days, userid)
	if err != nil {
		errMsg.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	messageData := pkg.NewDataResponse(outcomes)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(messageData)
}

func (outHand *OutcomeHandler) GetOutcomeLessThan(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	value := chi.URLParam(r, "value")
	var errMsg = pkg.NewMessageResponse("")

	if userid == "" || value == "" {
		errMsg.Message = "UserId and value is required !"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	v, err := strconv.ParseFloat(value, 64)

	if err != nil {
		errMsg.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetOutcomeLessThan(v, userid)
	if err != nil {
		errMsg.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	messageData := pkg.NewDataResponse(outcomes)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(messageData)
}

func (outHand *OutcomeHandler) GetOutcomeHigherThan(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	value := chi.URLParam(r, "value")
	var errMsg = pkg.NewMessageResponse("")

	if userid == "" || value == "" {
		errMsg.Message = "UserId and value is required"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	v, err := strconv.ParseFloat(value, 64)

	if err != nil {
		errMsg.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetOutcomeHigherThan(v, userid)
	if err != nil {
		errMsg.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	messageData := pkg.NewDataResponse(outcomes)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(messageData)
}

func (outHand *OutcomeHandler) DeleteOutcome(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var msgRes = pkg.NewMessageResponse("")

	if id == "" {
		msgRes.Message = "Id is required"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	msg, status, err := outHand.OutcomeDb.DeleteOutcome(id)
	if err != nil {
		msgRes.Message = msg

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	msgRes.Message = msg

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(msgRes)
}
