package database

import (
	"database/sql"
	"fmt"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
)

type BuyOperationDB struct {
	DB *sql.DB
}

func NewBuyOperationDB(db *sql.DB) *BuyOperationDB {
	return &BuyOperationDB{
		DB: db,
	}
}

func (buyop *BuyOperationDB) CreateBuyOperation(buyOpData *entity.BuyOperation) (int, error) {
	stmt, err := buyop.DB.Prepare(`
		INSERT INTO buyoperation VALUES (
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

	_, err = stmt.Exec(buyOpData.Id.String(), buyOpData.InvestimentId, buyOpData.Category, buyOpData.StockCode, buyOpData.Quantity, buyOpData.BuyPrice, buyOpData.Value, buyOpData.CreatedAt, buyOpData.BuyDate)

	if err != nil {
		fmt.Println(err.Error())
		return 500, err
	}

	return 201, nil
}

func (buyop *BuyOperationDB) GetBuyOperationByName(investmentId string, stockName string) (*entity.BuyOperation, int, error) {
	var buyOp entity.BuyOperation
	stmt, err := buyop.DB.Prepare("SELECT category, stockcode, quantity, buyprice, value, buydate FROM buyoperation WHERE investimentid = $1 AND stockcode = $2")

	if err != nil {
		fmt.Println(err.Error())
		return nil, 500, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(investmentId, stockName).Scan(&buyOp.Category, &buyOp.StockCode, &buyOp.Quantity, &buyOp.BuyPrice, &buyOp.Value, &buyOp.BuyDate)

	if err != nil {
		fmt.Println(err.Error())
		return nil, 404, err
	}

	return &buyOp, 200, nil
}

func (buyop *BuyOperationDB) GetBuyOperationsLessThan(investmentId string, value float64) ([]entity.BuyOperation, int, error) {
	var buyOps []entity.BuyOperation
	stmt, err := buyop.DB.Prepare("SELECT category, stockcode, quantity, buyprice, value, buydate FROM buyoperation WHERE investimentid = $1 AND value < $2")

	if err != nil {
		fmt.Println(err.Error())
		return buyOps, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(investmentId, value)

	if err != nil {
		fmt.Println(err.Error())
		return buyOps, 404, err
	}

	for rows.Next() {
		var buy entity.BuyOperation

		err = rows.Scan(&buy.Category, &buy.StockCode, &buy.Quantity, &buy.BuyPrice, &buy.Value, &buy.BuyDate)
		if err != nil {
			fmt.Println(err.Error())
			return buyOps, 404, err
		}

		buyOps = append(buyOps, buy)
	}

	if len(buyOps) == 0 {
		return buyOps, 404, fmt.Errorf("nenhuma operação de compra encontrada")
	}

	return buyOps, 200, nil
}

func (buyop *BuyOperationDB) GetBuyOperationsHigherThan(investmentId string, value float64) ([]entity.BuyOperation, int, error) {
	var buyOps []entity.BuyOperation
	stmt, err := buyop.DB.Prepare("SELECT category, stockcode, quantity, buyprice, value, buydate FROM buyoperation WHERE investimentid = $1 AND value > $2")

	if err != nil {
		fmt.Println(err.Error())
		return buyOps, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(investmentId, value)

	if err != nil {
		fmt.Println(err.Error())
		return buyOps, 404, err
	}

	for rows.Next() {
		var buy entity.BuyOperation

		err = rows.Scan(&buy.Category, &buy.StockCode, &buy.Quantity, &buy.BuyPrice, &buy.Value, &buy.BuyDate)
		if err != nil {
			fmt.Println(err.Error())
			return buyOps, 404, err
		}

		buyOps = append(buyOps, buy)
	}

	if len(buyOps) == 0 {
		return buyOps, 404, fmt.Errorf("nenhuma operação de compra encontrada")
	}

	return buyOps, 200, nil
}

func (buyop *BuyOperationDB) GetBuyOperationsByCategory(investmentId, category string) ([]entity.BuyOperation, int, error) {
	var buyOps []entity.BuyOperation

	stmt, err := buyop.DB.Prepare("SELECT category, stockcode, quantity, buyprice, value, buydate FROM buyoperation WHERE investimentid = $1 AND category = $2")

	if err != nil {
		fmt.Println(err.Error())
		return buyOps, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(investmentId, category)

	if err != nil {
		fmt.Println(err.Error())
		return buyOps, 404, err
	}

	for rows.Next() {
		var buy entity.BuyOperation

		err = rows.Scan(&buy.Category, &buy.StockCode, &buy.Quantity, &buy.BuyPrice, &buy.Value, &buy.BuyDate)

		if err != nil {
			fmt.Println(err.Error())
			return buyOps, 404, err
		}

		buyOps = append(buyOps, buy)
	}

	if len(buyOps) == 0 {
		return buyOps, 404, fmt.Errorf("nenhum operação de compra encontrada")
	}
	return buyOps, 200, nil
}

func (buyop *BuyOperationDB) ListAllBuyOperation(investmentId string) ([]entity.BuyOperation, int, error) {
	var buyOps []entity.BuyOperation

	stmt, err := buyop.DB.Prepare("SELECT category, stockcode, quantity, buyprice, value, buydate FROM buyoperation WHERE investimentid = $1")

	if err != nil {
		fmt.Println(err.Error())
		return buyOps, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(investmentId)

	if err != nil {
		fmt.Println(err.Error())
		return buyOps, 404, err
	}

	for rows.Next() {
		var buy entity.BuyOperation

		err = rows.Scan(&buy.Category, &buy.StockCode, &buy.Quantity, &buy.BuyPrice, &buy.Value, &buy.BuyDate)

		if err != nil {
			fmt.Println(err.Error())
			return buyOps, 404, err
		}

		buyOps = append(buyOps, buy)
	}

	if len(buyOps) == 0 {
		return buyOps, 404, fmt.Errorf("nenhum operação de compra encontrada")
	}
	return buyOps, 200, nil
}

func (buyop *BuyOperationDB) ListAllBuyOperationOfMonth(investmentId string, month int) ([]entity.BuyOperation, int, error) {
	var buyOps []entity.BuyOperation

	stmt, err := buyop.DB.Prepare("SELECT category, stockcode, quantity, buyprice, value, buydate FROM buyoperation WHERE investimentid = $1 AND date_part('month', buydate) = $2")

	if err != nil {
		fmt.Println(err.Error())
		return buyOps, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(investmentId, month)

	if err != nil {
		fmt.Println(err.Error())
		return buyOps, 404, err
	}

	for rows.Next() {
		var buy entity.BuyOperation

		err = rows.Scan(&buy.Category, &buy.StockCode, &buy.Quantity, &buy.BuyPrice, &buy.Value, &buy.BuyDate)

		if err != nil {
			fmt.Println(err.Error())
			return buyOps, 404, err
		}

		buyOps = append(buyOps, buy)
	}

	if len(buyOps) == 0 {
		return buyOps, 404, fmt.Errorf("nenhum operação de compra encontrada")
	}
	return buyOps, 200, nil
}
