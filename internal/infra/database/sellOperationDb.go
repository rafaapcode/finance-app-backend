package database

import (
	"database/sql"
	"fmt"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
)

type SellOperationDB struct {
	DB *sql.DB
}

func NewSellOperationDB(db *sql.DB) *SellOperationDB {
	return &SellOperationDB{
		DB: db,
	}
}

func (sellop *SellOperationDB) CreateSellOperation(sellOpData *entity.SellOperation) (int, error) {
	stmt, err := sellop.DB.Prepare(`
		INSERT INTO selloperation VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9
		);
	`)

	if err != nil {
		fmt.Println(err.Error())
		return 500, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(sellOpData.Id.String(), sellOpData.InvestimentId, sellOpData.Category, sellOpData.StockCode, sellOpData.Quantity, sellOpData.SellPrice, sellOpData.Value, sellOpData.CreatedAt, sellOpData.SellDate)

	if err != nil {
		fmt.Println(err.Error())
		return 500, err
	}

	return 201, nil
}

func (sellop *SellOperationDB) GetSellOperationByName(investmentId string, stockName string) (*entity.SellOperation, int, error) {
	var sellOp entity.SellOperation
	stmt, err := sellop.DB.Prepare("SELECT category, stockcode, quantity, sellprice, value, selldate FROM selloperation WHERE investimentid = $1 AND stockcode = $2")

	if err != nil {
		fmt.Println(err.Error())
		return nil, 500, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(investmentId, stockName).Scan(&sellOp.Category, &sellOp.StockCode, &sellOp.Quantity, &sellOp.SellPrice, &sellOp.Value, &sellOp.SellDate)

	if err != nil {
		fmt.Println(err.Error())
		return nil, 404, err
	}

	return &sellOp, 200, nil
}

func (sellop *SellOperationDB) GetSellOperationsLessThan(investmentId string, value float64) ([]entity.SellOperation, int, error) {
	var sellOps []entity.SellOperation
	stmt, err := sellop.DB.Prepare("SELECT category, stockcode, quantity, sellprice, value, selldate FROM selloperation WHERE investimentid = $1 AND value < $2")

	if err != nil {
		fmt.Println(err.Error())
		return sellOps, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(investmentId, value)

	if err != nil {
		fmt.Println(err.Error())
		return sellOps, 404, err
	}

	for rows.Next() {
		var sell entity.SellOperation

		err = rows.Scan(&sell.Category, &sell.StockCode, &sell.Quantity, &sell.SellPrice, &sell.Value, &sell.SellDate)
		if err != nil {
			fmt.Println(err.Error())
			return sellOps, 404, err
		}

		sellOps = append(sellOps, sell)
	}

	if len(sellOps) == 0 {
		return sellOps, 404, fmt.Errorf("nenhuma operação de venda encontrada")
	}

	return sellOps, 200, nil
}

func (sellop *SellOperationDB) GetSellOperationsHigherThan(investmentId string, value float64) ([]entity.SellOperation, int, error) {
	var sellOps []entity.SellOperation
	stmt, err := sellop.DB.Prepare("SELECT category, stockcode, quantity, sellprice, value, selldate FROM selloperation WHERE investimentid = $1 AND value > $2")

	if err != nil {
		fmt.Println(err.Error())
		return sellOps, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(investmentId, value)

	if err != nil {
		fmt.Println(err.Error())
		return sellOps, 404, err
	}

	for rows.Next() {
		var sell entity.SellOperation

		err = rows.Scan(&sell.Category, &sell.StockCode, &sell.Quantity, &sell.SellPrice, &sell.Value, &sell.SellDate)
		if err != nil {
			fmt.Println(err.Error())
			return sellOps, 404, err
		}

		sellOps = append(sellOps, sell)
	}

	if len(sellOps) == 0 {
		return sellOps, 404, fmt.Errorf("nenhuma operação de venda encontrada")
	}

	return sellOps, 200, nil
}

func (sellop *SellOperationDB) GetSellOperationsByCategory(investmentId, category string) ([]entity.SellOperation, int, error) {
	var sellOps []entity.SellOperation

	stmt, err := sellop.DB.Prepare("SELECT category, stockcode, quantity, sellprice, value, selldate FROM selloperation WHERE investimentid = $1 AND category = $2")

	if err != nil {
		fmt.Println(err.Error())
		return sellOps, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(investmentId, category)

	if err != nil {
		fmt.Println(err.Error())
		return sellOps, 404, err
	}

	for rows.Next() {
		var sell entity.SellOperation

		err = rows.Scan(&sell.Category, &sell.StockCode, &sell.Quantity, &sell.SellPrice, &sell.Value, &sell.SellDate)

		if err != nil {
			fmt.Println(err.Error())
			return sellOps, 404, err
		}

		sellOps = append(sellOps, sell)
	}

	if len(sellOps) == 0 {
		return sellOps, 404, fmt.Errorf("nenhum operação de venda encontrada")
	}
	return sellOps, 200, nil
}

func (sellop *SellOperationDB) ListAllSellOperation(investmentId string) ([]entity.SellOperation, int, error) {
	var sellOps []entity.SellOperation

	stmt, err := sellop.DB.Prepare("SELECT category, stockcode, quantity, sellprice, value, selldate FROM selloperation WHERE investimentid = $1")

	if err != nil {
		fmt.Println(err.Error())
		return sellOps, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(investmentId)

	if err != nil {
		fmt.Println(err.Error())
		return sellOps, 404, err
	}

	for rows.Next() {
		var sell entity.SellOperation

		err = rows.Scan(&sell.Category, &sell.StockCode, &sell.Quantity, &sell.SellPrice, &sell.Value, &sell.SellDate)

		if err != nil {
			fmt.Println(err.Error())
			return sellOps, 404, err
		}

		sellOps = append(sellOps, sell)
	}

	if len(sellOps) == 0 {
		return sellOps, 404, fmt.Errorf("nenhum operação de venda encontrada")
	}
	return sellOps, 200, nil
}

func (sellop *SellOperationDB) ListAllSellOperationOfMonth(investmentId string, month int) ([]entity.SellOperation, int, error) {
	var sellOps []entity.SellOperation

	stmt, err := sellop.DB.Prepare("SELECT category, stockcode, quantity, sellprice, value, selldate FROM selloperation WHERE investimentid = $1 AND date_part('month', selldate) = $2")

	if err != nil {
		fmt.Println(err.Error())
		return sellOps, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(investmentId, month)

	if err != nil {
		fmt.Println(err.Error())
		return sellOps, 404, err
	}

	for rows.Next() {
		var sell entity.SellOperation

		err = rows.Scan(&sell.Category, &sell.StockCode, &sell.Quantity, &sell.SellPrice, &sell.Value, &sell.SellDate)

		if err != nil {
			fmt.Println(err.Error())
			return sellOps, 404, err
		}

		sellOps = append(sellOps, sell)
	}

	if len(sellOps) == 0 {
		return sellOps, 404, fmt.Errorf("nenhum operação de venda encontrada")
	}
	return sellOps, 200, nil
}
