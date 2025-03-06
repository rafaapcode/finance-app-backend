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

	err := json.NewDecoder(r.Body).Decode(&outcome)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	out, err := entity.NewOutcome(outcome.OutcomeType, outcome.Category, outcome.PaymentMethod, outcome.UserId, outcome.Value, outcome.Notification, outcome.ExpireDate)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = out.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	status, err := outHand.OutcomeDb.CreateOutcome(out)
	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Outcome created successfully !",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&response)
}

func (outHand *OutcomeHandler) GetOutcomeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	outcome, status, err := outHand.OutcomeDb.GetOutcomeById(id)
	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Data entity.Outcome `json:"data"`
	}{
		Data: *outcome,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&response)
}

func (outHand *OutcomeHandler) GetAllOutcomeOfMonth(w http.ResponseWriter, r *http.Request) {
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

	outcomes, status, err := outHand.OutcomeDb.GetAllOutcomeOfMonth(monthvalue, userid)
	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Data []entity.Outcome `json:"data"`
	}{
		Data: outcomes,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&response)
}

func (outHand *OutcomeHandler) GetAllFixedOutcome(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")

	if userid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetAllFixedOutcome(userid)
	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Data []entity.Outcome `json:"data"`
	}{
		Data: outcomes,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&response)
}

func (outHand *OutcomeHandler) GetAllOutcomeByCategory(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	category := chi.URLParam(r, "category")

	if userid == "" || category == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetAllOutcomeByCategory(category, userid)
	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Data []entity.Outcome `json:"data"`
	}{
		Data: outcomes,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&response)
}

func (outHand *OutcomeHandler) GetAllOutcomeByPaymentMethod(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	paymentMethod := chi.URLParam(r, "paymentmethod")

	if userid == "" || paymentMethod == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetAllOutcomeByPaymentMethod(paymentMethod, userid)
	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Data []entity.Outcome `json:"data"`
	}{
		Data: outcomes,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&response)
}

func (outHand *OutcomeHandler) GetAllOutcomeByType(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	outcomeType := chi.URLParam(r, "type")

	if userid == "" || outcomeType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetAllOutcomeByType(outcomeType, userid)
	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Data []entity.Outcome `json:"data"`
	}{
		Data: outcomes,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&response)
}

func (outHand *OutcomeHandler) GetOutcomeAboutToExpire(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	daysToexpire := chi.URLParam(r, "daysToExpire")

	if userid == "" || daysToexpire == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	days, err := strconv.Atoi(daysToexpire)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetOutcomeAboutToExpire(days, userid)
	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Data []entity.Outcome `json:"data"`
	}{
		Data: outcomes,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&response)
}

func (outHand *OutcomeHandler) GetOutcomeLessThan(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	value := chi.URLParam(r, "value")

	if userid == "" || value == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	v, err := strconv.ParseFloat(value, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetOutcomeLessThan(v, userid)
	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Data []entity.Outcome `json:"data"`
	}{
		Data: outcomes,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&response)
}

func (outHand *OutcomeHandler) GetOutcomeHigherThan(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	value := chi.URLParam(r, "value")

	if userid == "" || value == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	v, err := strconv.ParseFloat(value, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	outcomes, status, err := outHand.OutcomeDb.GetOutcomeHigherThan(v, userid)
	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Data []entity.Outcome `json:"data"`
	}{
		Data: outcomes,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&response)
}

func (outHand *OutcomeHandler) DeleteOutcome(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	msg, status, err := outHand.OutcomeDb.DeleteOutcome(id)
	if err != nil {
		w.WriteHeader(status)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&response)
}
