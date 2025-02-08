package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
)

type InvestmentDb struct {
	DB *sql.DB
}

func NewInvestmentDB(db *sql.DB) *InvestmentDb {
	return &InvestmentDb{
		DB: db,
	}
}

func (invDb *InvestmentDb) CreateInvestment(invest *entity.Investment) (int, error) {
	stmt, err := invDb.DB.Prepare(`
		INSERT INTO investment VALUES(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11,
			$12,
			$13,
			$14
		);
	`)

	if err != nil {
		fmt.Println(err.Error())
		return 500, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(invest.Id.String(), invest.Userid, invest.Category, invest.StockCode, invest.TotalQuantity, invest.BuyPrice, invest.SellPrice, invest.Percentage, invest.Value, invest.Profit, invest.BuyDate, invest.SellDate, invest.CreatedAt, invest.LastSupplyDate)

	if err != nil {
		fmt.Println(err.Error())
		return 500, err
	}

	return 201, nil
}

func (invDb *InvestmentDb) GetTotalOfInvestment(userid string) (float64, int, error) {
	var total float64
	stmt, err := invDb.DB.Prepare("SELECT sum(value) FROM investment WHERE userid = $1")

	if err != nil {
		fmt.Println(err.Error())
		return 0, 500, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(userid).Scan(&total)

	if err != nil {
		fmt.Println(err.Error())
		return 0, 404, err
	}

	return total, 200, nil
}

func (invDb *InvestmentDb) GetAllOfInvestment(userid string) ([]entity.Investment, int, error) {
	var investments []entity.Investment
	stmt, err := invDb.DB.Prepare("SELECT id, userid, category, stockcode, totalquantity, buyprice, sellprice, percentage, value, profit, buydate, selldate, createdat, lastsupplydate FROM investment WHERE userid = $1 ORDER BY createdat ASC LIMIT 5")

	if err != nil {
		fmt.Println(err.Error())
		return investments, 500, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(userid)

	if err != nil {
		fmt.Println(err.Error())
		return investments, 404, err
	}

	for rows.Next() {
		var investment entity.Investment

		err = rows.Scan(&investment.Id, &investment.Userid, &investment.Category, &investment.StockCode, &investment.TotalQuantity, &investment.BuyPrice, &investment.SellPrice, &investment.Percentage, &investment.Value, &investment.Profit, &investment.BuyDate, &investment.SellDate, &investment.CreatedAt, &investment.LastSupplyDate)

		if err != nil {
			return investments, 404, err
		}

		investments = append(investments, investment)
	}

	if len(investments) == 0 {
		return investments, 404, fmt.Errorf("nenhum investimento encontrado")
	}

	return investments, 200, nil
}

func (invDb *InvestmentDb) GetNextPageAllOfInvestment(lastInvId, userid string) ([]entity.Investment, int, error) {
	var investments []entity.Investment
	stmt, err := invDb.DB.Prepare("SELECT id, userid, category, stockcode, totalquantity, buyprice, sellprice, percentage, value, profit, buydate, selldate, createdat, lastsupplydate FROM investment WHERE id > $1 AND userid = $2 ORDER BY createdat ASC LIMIT 5")

	if err != nil {
		fmt.Println(err.Error())
		return investments, 500, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(lastInvId, userid)

	if err != nil {
		fmt.Println(err.Error())
		return investments, 404, err
	}

	for rows.Next() {
		var investment entity.Investment

		err = rows.Scan(&investment.Id, &investment.Userid, &investment.Category, &investment.StockCode, &investment.TotalQuantity, &investment.BuyPrice, &investment.SellPrice, &investment.Percentage, &investment.Value, &investment.Profit, &investment.BuyDate, &investment.SellDate, &investment.CreatedAt, &investment.LastSupplyDate)

		if err != nil {
			return investments, 404, err
		}

		investments = append(investments, investment)
	}

	if len(investments) == 0 {
		return investments, 404, fmt.Errorf("nenhum investimento encontrado")
	}

	return investments, 200, nil
}

func (invDb *InvestmentDb) GetPreviousPageAllOfInvestment(firstInvId, userid string) ([]entity.Investment, int, error) {
	var investments []entity.Investment
	stmt, err := invDb.DB.Prepare("SELECT id, userid, category, stockcode, totalquantity, buyprice, sellprice, percentage, value, profit, buydate, selldate, createdat, lastsupplydate FROM investment WHERE id < $1 AND userid = $2 ORDER BY createdat ASC LIMIT 5")

	if err != nil {
		fmt.Println(err.Error())
		return investments, 500, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(firstInvId, userid)

	if err != nil {
		fmt.Println(err.Error())
		return investments, 404, err
	}

	for rows.Next() {
		var investment entity.Investment

		err = rows.Scan(&investment.Id, &investment.Userid, &investment.Category, &investment.StockCode, &investment.TotalQuantity, &investment.BuyPrice, &investment.SellPrice, &investment.Percentage, &investment.Value, &investment.Profit, &investment.BuyDate, &investment.SellDate, &investment.CreatedAt, &investment.LastSupplyDate)

		if err != nil {
			return investments, 404, err
		}

		investments = append(investments, investment)
	}

	if len(investments) == 0 {
		return investments, 404, fmt.Errorf("nenhum investimento encontrado")
	}

	return investments, 200, nil
}

func (invDb *InvestmentDb) GetInvestmentByName(name, userid string) (entity.Investment, int, error) {
	var investment entity.Investment
	stmt, err := invDb.DB.Prepare("SELECT id, userid, category, stockcode, totalquantity, buyprice, sellprice, percentage, value, profit, buydate, selldate, createdat, lastsupplydate FROM investment WHERE stockcode = $1 AND userid = $2")

	if err != nil {
		fmt.Println(err.Error())
		return investment, 500, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(name, userid).Scan(&investment.Id, &investment.Userid, &investment.Category, &investment.StockCode, &investment.TotalQuantity, &investment.BuyPrice, &investment.SellPrice, &investment.Percentage, &investment.Value, &investment.Profit, &investment.BuyDate, &investment.SellDate, &investment.CreatedAt, &investment.LastSupplyDate)

	if err != nil {
		fmt.Println(err.Error())
		return investment, 404, err
	}

	return investment, 200, nil
}

func (invDb *InvestmentDb) GetInvestmentById(id string) (entity.Investment, int, error) {
	var investment entity.Investment
	stmt, err := invDb.DB.Prepare("SELECT id, userid, category, stockcode, totalquantity, buyprice, sellprice, percentage, value, profit, buydate, selldate, createdat, lastsupplydate FROM investment WHERE id = $1")

	if err != nil {
		fmt.Println(err.Error())
		return investment, 500, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&investment.Id, &investment.Userid, &investment.Category, &investment.StockCode, &investment.TotalQuantity, &investment.BuyPrice, &investment.SellPrice, &investment.Percentage, &investment.Value, &investment.Profit, &investment.BuyDate, &investment.SellDate, &investment.CreatedAt, &investment.LastSupplyDate)

	if err != nil {
		fmt.Println(err.Error())
		return investment, 404, err
	}

	return investment, 200, nil
}

func (invDb *InvestmentDb) GetInvestmentByCategory(category, userid string) ([]entity.Investment, int, error) {
	var investments []entity.Investment
	stmt, err := invDb.DB.Prepare("SELECT id, userid, category, stockcode, totalquantity, buyprice, sellprice, percentage, value, profit, buydate, selldate, createdat, lastsupplydate FROM investment WHERE category = $1 AND userid = $2 ORDER BY createdat ASC LIMIT 5")

	if err != nil {
		fmt.Println(err.Error())
		return investments, 500, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(category, userid)

	if err != nil {
		fmt.Println(err.Error())
		return investments, 404, err
	}

	for rows.Next() {
		var investment entity.Investment

		err = rows.Scan(&investment.Id, &investment.Userid, &investment.Category, &investment.StockCode, &investment.TotalQuantity, &investment.BuyPrice, &investment.SellPrice, &investment.Percentage, &investment.Value, &investment.Profit, &investment.BuyDate, &investment.SellDate, &investment.CreatedAt, &investment.LastSupplyDate)

		if err != nil {
			return investments, 404, err
		}

		investments = append(investments, investment)
	}

	if len(investments) == 0 {
		return investments, 404, fmt.Errorf("nenhum investimento encontrado")
	}

	return investments, 200, nil
}

func (invDb *InvestmentDb) GetNextPageInvestmentByCategory(category, userid, lastInvId string) ([]entity.Investment, int, error) {
	var investments []entity.Investment
	stmt, err := invDb.DB.Prepare("SELECT id, userid, category, stockcode, totalquantity, buyprice, sellprice, percentage, value, profit, buydate, selldate, createdat, lastsupplydate FROM investment WHERE category = $1 AND userid = $2 AND id > $3 ORDER BY createdat ASC LIMIT 5")

	if err != nil {
		fmt.Println(err.Error())
		return investments, 500, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(category, userid, lastInvId)

	if err != nil {
		fmt.Println(err.Error())
		return investments, 404, err
	}

	for rows.Next() {
		var investment entity.Investment

		err = rows.Scan(&investment.Id, &investment.Userid, &investment.Category, &investment.StockCode, &investment.TotalQuantity, &investment.BuyPrice, &investment.SellPrice, &investment.Percentage, &investment.Value, &investment.Profit, &investment.BuyDate, &investment.SellDate, &investment.CreatedAt, &investment.LastSupplyDate)

		if err != nil {
			return investments, 404, err
		}

		investments = append(investments, investment)
	}

	if len(investments) == 0 {
		return investments, 404, fmt.Errorf("nenhum investimento encontrado")
	}

	return investments, 200, nil
}

func (invDb *InvestmentDb) GetPreviousPageInvestmentByCategory(category, userid, firstInvId string) ([]entity.Investment, int, error) {
	var investments []entity.Investment
	stmt, err := invDb.DB.Prepare("SELECT id, userid, category, stockcode, totalquantity, buyprice, sellprice, percentage, value, profit, buydate, selldate, createdat, lastsupplydate FROM investment WHERE category = $1 AND userid = $2 AND id < $3 ORDER BY createdat ASC LIMIT 5")

	if err != nil {
		fmt.Println(err.Error())
		return investments, 500, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(category, userid, firstInvId)

	if err != nil {
		fmt.Println(err.Error())
		return investments, 404, err
	}

	for rows.Next() {
		var investment entity.Investment

		err = rows.Scan(&investment.Id, &investment.Userid, &investment.Category, &investment.StockCode, &investment.TotalQuantity, &investment.BuyPrice, &investment.SellPrice, &investment.Percentage, &investment.Value, &investment.Profit, &investment.BuyDate, &investment.SellDate, &investment.CreatedAt, &investment.LastSupplyDate)

		if err != nil {
			return investments, 404, err
		}

		investments = append(investments, investment)
	}

	if len(investments) == 0 {
		return investments, 404, fmt.Errorf("nenhum investimento encontrado")
	}

	return investments, 200, nil
}

func (invDb *InvestmentDb) UpdateInvestment(newData *entity.Investment) (int, error) {
	stmt, err := invDb.DB.Prepare("UPDATE investment SET totalquantity = $1, sellprice = $2, percentage = $3, value = $4, profit = $5, buydate = $6, selldate = $7, lastsupplydate = $8 WHERE id = $9 AND userid = $10")

	if err != nil {
		fmt.Println(err.Error())
		return 500, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newData.TotalQuantity, newData.SellPrice, newData.Percentage, newData.Value, newData.Profit, newData.BuyDate, newData.SellDate, newData.LastSupplyDate, newData.Id.String(), newData.Userid)

	if err != nil {
		fmt.Println(err.Error())
		return 404, err
	}

	return 200, nil
}

func (invDb *InvestmentDb) GetAssetGrowth(userid string) (entity.Metrics, int, error) {
	var metrics map[string]float64 = make(map[string]float64)

	stmt, err := invDb.DB.Prepare(`
		SELECT 
			date_part('month', buydate) AS month,
			SUM(value) AS total
		FROM investment
		WHERE userid = '0194565b-673e-7f4b-aff0-83345f230df3' AND sellprice = 0.00
		GROUP BY month
		ORDER BY month;
	`)

	if err != nil {
		fmt.Println(err.Error())
		return entity.Metrics{}, 500, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(userid)

	if err != nil {
		fmt.Println(err.Error())
		return entity.Metrics{}, 404, err
	}

	for rows.Next() {
		var month int
		var totalValue float64

		err = rows.Scan(&month, &totalValue)

		if err != nil {
			fmt.Println(err.Error())
			return entity.Metrics{}, 404, err
		}

		formatMonth := time.Month(month)

		_, ok := metrics[formatMonth.String()]

		if !ok {
			metrics[formatMonth.String()] = totalValue
		} else {
			metrics[formatMonth.String()] += totalValue
		}

	}

	return *entity.NewMetrics(metrics), 200, nil
}

func (invDb *InvestmentDb) GetPortfolioDiversification(userid string) (entity.Metrics, int, error) {
	var metrics map[string]float64 = make(map[string]float64)
	stmt, err := invDb.DB.Prepare(`
		SELECT 
			category,
			SUM(value) AS total
		FROM investment
		WHERE userid = $1 AND sellprice = 0.00
		GROUP BY category;
	`)

	if err != nil {
		fmt.Println(err.Error())
		return entity.Metrics{}, 500, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(userid)

	if err != nil {
		fmt.Println(err.Error())
		return entity.Metrics{}, 404, err
	}

	for rows.Next() {
		var category string
		var totalValue float64

		err = rows.Scan(&category, &totalValue)

		if err != nil {
			fmt.Println(err.Error())
			return entity.Metrics{}, 404, err
		}

		_, ok := metrics[category]

		if !ok {
			metrics[category] = totalValue
		} else {
			metrics[category] += totalValue
		}

	}

	return *entity.NewMetrics(metrics), 200, nil
}

func (invDb *InvestmentDb) GetMonthInvestment(userid string, month int) (entity.Metrics, int, error) {
	var metrics map[string]float64 = make(map[string]float64)

	stmt, err := invDb.DB.Prepare("SELECT buydate, value FROM investment WHERE date_part('month', buydate) = $1 AND userid = $2")

	if err != nil {
		fmt.Println(err.Error())
		return entity.Metrics{}, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(month, userid)

	if err != nil {
		fmt.Println(err.Error())
		return entity.Metrics{}, 404, err
	}

	for rows.Next() {
		var buyDate string
		var value float64

		err = rows.Scan(&buyDate, &value)

		if err != nil {
			fmt.Println(err.Error())
			return entity.Metrics{}, 404, err
		}

		_, ok := metrics[buyDate]

		if !ok {
			metrics[buyDate] = value
		} else {
			metrics[buyDate] += value
		}

	}

	return *entity.NewMetrics(metrics), 200, nil
}
