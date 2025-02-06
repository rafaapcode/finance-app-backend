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

func (invDb *InvestmentDb) GetAllOfInvestment(pageNumber int, sort, userid string) ([]entity.Investment, int, error) {
	var investments []entity.Investment

	return investments, 500, nil
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

func (invDb *InvestmentDb) GetInvestmentByCategory(pageNumber int, category, userid string) ([]entity.Investment, int, error) {
	var investments []entity.Investment

	return investments, 500, nil
}

func (invDb *InvestmentDb) UpdateInvestment(newData *entity.Investment) (int, error) {
	return 500, nil
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

// A testar
func (invDb *InvestmentDb) GetMonthInvestment(userid string, month time.Month) (entity.Metrics, int, error) {
	return entity.Metrics{}, 200, nil
}
