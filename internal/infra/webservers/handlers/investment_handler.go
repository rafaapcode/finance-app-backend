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

type InvestmentHandler struct {
	InvestmentDb database.InvestmentInterface
	BuyOpDb      database.BuyOperationInterface
	SellOpDb     database.SellOperationInterface
	SupplyOpDb   database.SupplyOperationInterface
}

func NewInvestmentHandler(income database.IncomeInterface, extraIncomeDb database.ExtraIncomeInterface) *IncomeHandler {
	return &IncomeHandler{
		IncomeDb:      income,
		ExtraIncomeDb: extraIncomeDb,
	}
}

func (invDb *InvestmentHandler) CreateInvestmentHandler(w http.ResponseWriter, r *http.Request) {
	var investmentBody dto.CreateInvestmentDto
	var msgRes = pkg.NewMessageResponse("")

	if err := json.NewDecoder(r.Body).Decode(&investmentBody); err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	totalInvested, status, err := invDb.InvestmentDb.GetTotalOfInvestment(investmentBody.UserId)

	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	var value float64 = investmentBody.BuyPrice * float64(investmentBody.Quantity)

	totalInvested += value

	var percentageOfTotal = (value / totalInvested)

	inv, err := entity.NewInvestment(investmentBody.Category, investmentBody.UserId, investmentBody.StockCode, investmentBody.Quantity, investmentBody.BuyPrice, 0.00, value, 0.00, float32(percentageOfTotal), time.Now())

	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	err = inv.Validate()
	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	buyOp, err := entity.NewBuyOperation(inv.Id.String(), inv.Category, inv.StockCode, investmentBody.Quantity, investmentBody.BuyPrice, value, inv.BuyDate)

	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	err = buyOp.Validate()
	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	status, err = invDb.BuyOpDb.CreateBuyOperation(buyOp)
	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	status, err = invDb.InvestmentDb.CreateInvestment(inv)
	if err != nil {
		msgRes.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	msgRes.Message = "Investment created with success"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(msgRes)
}

func (invDb *InvestmentHandler) GetTotalOfInvestment(w http.ResponseWriter, r *http.Request) {
	var msgRes = pkg.NewMessageResponse("")
	userid := chi.URLParam(r, "userid")

	if userid == "" {
		msgRes.Message = "User Id is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	totalOfInvest, status, err := invDb.InvestmentDb.GetTotalOfInvestment(userid)

	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	data := pkg.NewDataResponse(totalOfInvest)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (invDb *InvestmentHandler) GetAllOfInvestment(w http.ResponseWriter, r *http.Request) {
	var msgRes = pkg.NewMessageResponse("")
	userid := chi.URLParam(r, "userid")

	if userid == "" {
		msgRes.Message = "User Id is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	allInvestments, status, err := invDb.InvestmentDb.GetAllOfInvestment(userid)

	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	data := pkg.NewDataResponse(allInvestments)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (invDb *InvestmentHandler) GetNextPageOfAllInvestment(w http.ResponseWriter, r *http.Request) {
	var msgRes = pkg.NewMessageResponse("")
	userid := chi.URLParam(r, "userid")
	lastInvdId := chi.URLParam(r, "lastinvid")

	if userid == "" || lastInvdId == "" {
		msgRes.Message = "User Id and Invesment Id is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	allInvestments, status, err := invDb.InvestmentDb.GetNextPageAllOfInvestment(lastInvdId, userid)

	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	data := pkg.NewDataResponse(allInvestments)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (invDb *InvestmentHandler) GetPreviousPageOfAllInvestment(w http.ResponseWriter, r *http.Request) {
	var msgRes = pkg.NewMessageResponse("")
	userid := chi.URLParam(r, "userid")
	firstInvdId := chi.URLParam(r, "firstinvid")

	if userid == "" || firstInvdId == "" {
		msgRes.Message = "User Id and Invesment Id is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	allInvestments, status, err := invDb.InvestmentDb.GetPreviousPageAllOfInvestment(firstInvdId, userid)

	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	data := pkg.NewDataResponse(allInvestments)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (invDb *InvestmentHandler) GetInvesmentByName(w http.ResponseWriter, r *http.Request) {
	var msgRes = pkg.NewMessageResponse("")
	userid := chi.URLParam(r, "userid")
	stockName := chi.URLParam(r, "stockname")

	if userid == "" || stockName == "" {
		msgRes.Message = "User Id and Name of Stock is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	investment, status, err := invDb.InvestmentDb.GetInvestmentByName(stockName, userid)

	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	data := pkg.NewDataResponse(investment)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (invDb *InvestmentHandler) GetInvesmentById(w http.ResponseWriter, r *http.Request) {
	var msgRes = pkg.NewMessageResponse("")
	invId := chi.URLParam(r, "invId")

	if invId == "" {
		msgRes.Message = "Investment Id is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	investment, status, err := invDb.InvestmentDb.GetInvestmentById(invId)

	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	data := pkg.NewDataResponse(investment)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (invDb *InvestmentHandler) GetInvesmentByCategory(w http.ResponseWriter, r *http.Request) {
	var msgRes = pkg.NewMessageResponse("")
	userid := chi.URLParam(r, "userid")
	category := chi.URLParam(r, "category")

	if userid == "" || category == "" {
		msgRes.Message = "UserId and category is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	investments, status, err := invDb.InvestmentDb.GetInvestmentByCategory(category, userid)

	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	data := pkg.NewDataResponse(investments)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (invDb *InvestmentHandler) GetNextPageInvesmentByCategory(w http.ResponseWriter, r *http.Request) {
	var msgRes = pkg.NewMessageResponse("")
	userid := chi.URLParam(r, "userid")
	category := chi.URLParam(r, "category")
	lastInvId := chi.URLParam(r, "lastinvid")

	if userid == "" || category == "" || lastInvId == "" {
		msgRes.Message = "UserId, Category and Investment Id is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	investments, status, err := invDb.InvestmentDb.GetNextPageInvestmentByCategory(category, userid, lastInvId)

	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	data := pkg.NewDataResponse(investments)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (invDb *InvestmentHandler) GetPreviousPageInvesmentByCategory(w http.ResponseWriter, r *http.Request) {
	var msgRes = pkg.NewMessageResponse("")
	userid := chi.URLParam(r, "userid")
	category := chi.URLParam(r, "category")
	firstInvId := chi.URLParam(r, "firstInvId")

	if userid == "" || category == "" || firstInvId == "" {
		msgRes.Message = "UserId, Category and Investment Id is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	investments, status, err := invDb.InvestmentDb.GetPreviousPageInvestmentByCategory(category, userid, firstInvId)

	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	data := pkg.NewDataResponse(investments)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (invDb *InvestmentHandler) GetAssetGrowth(w http.ResponseWriter, r *http.Request) {
	var msgRes = pkg.NewMessageResponse("")
	userid := chi.URLParam(r, "userid")

	if userid == "" {
		msgRes.Message = "UserId is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	metrics, status, err := invDb.InvestmentDb.GetAssetGrowth(userid)

	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	data := pkg.NewDataResponse(metrics)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (invDb *InvestmentHandler) GetPortfolioDiversification(w http.ResponseWriter, r *http.Request) {
	var msgRes = pkg.NewMessageResponse("")
	userid := chi.URLParam(r, "userid")

	if userid == "" {
		msgRes.Message = "UserId is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	metrics, status, err := invDb.InvestmentDb.GetPortfolioDiversification(userid)

	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	data := pkg.NewDataResponse(metrics)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (invDb *InvestmentHandler) GetMonthInvestment(w http.ResponseWriter, r *http.Request) {
	var msgRes = pkg.NewMessageResponse("")
	userid := chi.URLParam(r, "userid")
	month := chi.URLParam(r, "month")

	if userid == "" || month == "" {
		msgRes.Message = "UserId and Month is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	monthValue, err := strconv.Atoi(month)

	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	if monthValue <= 0 || monthValue > 12 {
		msgRes.Message = "Month is invalid"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	metrics, status, err := invDb.InvestmentDb.GetMonthInvestment(userid, monthValue)

	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	data := pkg.NewDataResponse(metrics)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (invDb *InvestmentHandler) UpdateSellInvesment(w http.ResponseWriter, r *http.Request) {
	var msgRes = pkg.NewMessageResponse("")
	id := chi.URLParam(r, "id")

	if id == "" {
		msgRes.Message = "InvestmentId is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	var updateSellInv dto.UpdateSellInvestmentDto

	if err := json.NewDecoder(r.Body).Decode(&updateSellInv); err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	inv, status, err := invDb.InvestmentDb.GetInvestmentById(id)

	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	totalValue, status, err := invDb.InvestmentDb.GetTotalOfInvestment(inv.Userid)
	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	if updateSellInv.Quantity > inv.TotalQuantity {
		msgRes.Message = "You can not sell more STOCKS than you have."
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}
	var newQuantity = inv.TotalQuantity - updateSellInv.Quantity
	var totalSelled = updateSellInv.SellPrice * float64(updateSellInv.Quantity)
	var profit = inv.Value - totalSelled
	var newTotalInvestedValue = totalValue - totalSelled
	var percentage = (totalSelled / newTotalInvestedValue)

	invEntity := entity.NewUpdateSellInvestment(&inv, updateSellInv.SellPrice, profit, newQuantity, float32(percentage))

	err = invEntity.Validate()

	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	status, err = invDb.InvestmentDb.UpdateInvestment(invEntity)

	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	sellOpEntity, err := entity.NewSellOperation(invEntity.Id.String(), invEntity.Category, invEntity.StockCode, updateSellInv.Quantity, totalSelled, updateSellInv.SellPrice, time.Now())

	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	err = sellOpEntity.Validate()
	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	status, err = invDb.SellOpDb.CreateSellOperation(sellOpEntity)
	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	msgRes.Message = "Investment selled with sucesss"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(msgRes)
}

func (invDb *InvestmentHandler) UpdateSupplyInvesment(w http.ResponseWriter, r *http.Request) {
	var msgRes = pkg.NewMessageResponse("")
	id := chi.URLParam(r, "id")

	if id == "" {
		msgRes.Message = "InvestmentId is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	var updateSupplyInv dto.UpdateSupplyInvestmentDto

	if err := json.NewDecoder(r.Body).Decode(&updateSupplyInv); err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	inv, status, err := invDb.InvestmentDb.GetInvestmentById(id)

	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	totalValue, status, err := invDb.InvestmentDb.GetTotalOfInvestment(inv.Userid)
	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	var newQuantity = inv.TotalQuantity + updateSupplyInv.Quantity
	var totalSuplied = updateSupplyInv.SupplyPrice * float64(updateSupplyInv.Quantity)
	var newTotalInvestedValue = totalValue + totalSuplied
	var percentage = (totalSuplied / newTotalInvestedValue)

	invEntity := entity.NewUpdateSupplyInvestment(&inv, updateSupplyInv.SupplyPrice, newTotalInvestedValue, newQuantity, float32(percentage))

	err = invEntity.Validate()

	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	status, err = invDb.InvestmentDb.UpdateInvestment(invEntity)

	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	supplyEntity, err := entity.NewSupplyOperation(invEntity.Id.String(), invEntity.Category, invEntity.StockCode, updateSupplyInv.Quantity, updateSupplyInv.SupplyPrice, totalSuplied, time.Now())

	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	err = supplyEntity.Validate()
	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	status, err = invDb.SupplyOpDb.CreateSupplyOperation(supplyEntity)
	if err != nil {
		msgRes.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(msgRes)
		return
	}

	msgRes.Message = "Investment suplied with sucesss"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(msgRes)
}
