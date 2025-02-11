package database

import (
	"database/sql"
	"fmt"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
)

type SupplyOperationDB struct {
	DB *sql.DB
}

func NewSupplyOperationDB(db *sql.DB) *SupplyOperationDB {
	return &SupplyOperationDB{
		DB: db,
	}
}

func (supplyop *SupplyOperationDB) CreateSupplyOperation(supplyOp *entity.SupplyOperation) (int, error) {
	stmt, err := supplyop.DB.Prepare(`
		INSERT INTO supplyoperation VALUES (
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

	_, err = stmt.Exec(supplyOp.Id.String(), supplyOp.InvestimentId, supplyOp.Category, supplyOp.StockCode, supplyOp.Quantity, supplyOp.SupplyPrice, supplyOp.Value, supplyOp.CreatedAt, supplyOp.SupplyDate)

	if err != nil {
		fmt.Println(err.Error())
		return 500, err
	}

	return 201, nil
}

func (supplyop *SupplyOperationDB) GetSupplyOperationByName(investmentId string, stockName string) ([]entity.SupplyOperation, int, error) {
	var supplyOps []entity.SupplyOperation
	stmt, err := supplyop.DB.Prepare("SELECT category, stockcode, quantity, supplyprice, value, supplydate FROM supplyoperation WHERE investimentid = $1 AND stockcode = $2")

	if err != nil {
		fmt.Println(err.Error())
		return supplyOps, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(investmentId, stockName)

	if err != nil {
		fmt.Println(err.Error())
		return supplyOps, 500, err
	}

	for rows.Next() {
		var supplyOp entity.SupplyOperation

		err = rows.Scan(&supplyOp.Category, &supplyOp.StockCode, &supplyOp.Quantity, &supplyOp.SupplyPrice, &supplyOp.Value, &supplyOp.SupplyDate)

		if err != nil {
			fmt.Println(err.Error())
			return supplyOps, 404, nil
		}

		supplyOps = append(supplyOps, supplyOp)
	}

	if len(supplyOps) == 0 {
		return supplyOps, 404, fmt.Errorf("nenhuma operação encontrada")
	}

	return supplyOps, 200, nil
}

func (supplyop *SupplyOperationDB) GetSupplyOperationsLessThan(investmentId string, value float64) ([]entity.SupplyOperation, int, error) {
	var supplyOps []entity.SupplyOperation
	stmt, err := supplyop.DB.Prepare("SELECT category, stockcode, quantity, supplyprice, value, supplydate FROM supplyoperation WHERE investimentid = $1 AND value < $2")

	if err != nil {
		fmt.Println(err.Error())
		return supplyOps, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(investmentId, value)

	if err != nil {
		fmt.Println(err.Error())
		return supplyOps, 500, err
	}

	for rows.Next() {
		var supplyOp entity.SupplyOperation

		err = rows.Scan(&supplyOp.Category, &supplyOp.StockCode, &supplyOp.Quantity, &supplyOp.SupplyPrice, &supplyOp.Value, &supplyOp.SupplyDate)

		if err != nil {
			fmt.Println(err.Error())
			return supplyOps, 404, nil
		}

		supplyOps = append(supplyOps, supplyOp)
	}

	if len(supplyOps) == 0 {
		return supplyOps, 404, fmt.Errorf("nenhuma operação encontrada")
	}
	return supplyOps, 200, nil
}

func (supplyop *SupplyOperationDB) GetSupplyOperationsHigherThan(investmentId string, value float64) ([]entity.SupplyOperation, int, error) {
	var supplyOps []entity.SupplyOperation
	stmt, err := supplyop.DB.Prepare("SELECT category, stockcode, quantity, supplyprice, value, supplydate FROM supplyoperation WHERE investimentid = $1 AND value > $2")

	if err != nil {
		fmt.Println(err.Error())
		return supplyOps, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(investmentId, value)

	if err != nil {
		fmt.Println(err.Error())
		return supplyOps, 500, err
	}

	for rows.Next() {
		var supplyOp entity.SupplyOperation

		err = rows.Scan(&supplyOp.Category, &supplyOp.StockCode, &supplyOp.Quantity, &supplyOp.SupplyPrice, &supplyOp.Value, &supplyOp.SupplyDate)

		if err != nil {
			fmt.Println(err.Error())
			return supplyOps, 404, nil
		}

		supplyOps = append(supplyOps, supplyOp)
	}

	if len(supplyOps) == 0 {
		return supplyOps, 404, fmt.Errorf("nenhuma operação encontrada")
	}

	return supplyOps, 200, nil
}

func (supplyop *SupplyOperationDB) GetSupplyOperationsByCategory(investmentId, category string) ([]entity.SupplyOperation, int, error) {
	var supplyOps []entity.SupplyOperation
	stmt, err := supplyop.DB.Prepare("SELECT category, stockcode, quantity, supplyprice, value, supplydate FROM supplyoperation WHERE investimentid = $1 AND category = $2")

	if err != nil {
		fmt.Println(err.Error())
		return supplyOps, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(investmentId, category)

	if err != nil {
		fmt.Println(err.Error())
		return supplyOps, 500, err
	}

	for rows.Next() {
		var supplyOp entity.SupplyOperation

		err = rows.Scan(&supplyOp.Category, &supplyOp.StockCode, &supplyOp.Quantity, &supplyOp.SupplyPrice, &supplyOp.Value, &supplyOp.SupplyDate)

		if err != nil {
			fmt.Println(err.Error())
			return supplyOps, 404, nil
		}

		supplyOps = append(supplyOps, supplyOp)
	}

	if len(supplyOps) == 0 {
		return supplyOps, 404, fmt.Errorf("nenhuma operação encontrada")
	}

	return supplyOps, 200, nil
}

func (supplyop *SupplyOperationDB) ListAllSupplyOperation(investmentId string) ([]entity.SupplyOperation, int, error) {
	var supplyOps []entity.SupplyOperation

	stmt, err := supplyop.DB.Prepare("SELECT category, stockcode, quantity, supplyprice, value, supplydate FROM supplyoperation WHERE investimentid = $1")

	if err != nil {
		fmt.Println(err.Error())
		return supplyOps, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(investmentId)

	if err != nil {
		fmt.Println(err.Error())
		return supplyOps, 500, err
	}

	for rows.Next() {
		var supplyOp entity.SupplyOperation

		err = rows.Scan(&supplyOp.Category, &supplyOp.StockCode, &supplyOp.Quantity, &supplyOp.SupplyPrice, &supplyOp.Value, &supplyOp.SupplyDate)

		if err != nil {
			fmt.Println(err.Error())
			return supplyOps, 404, nil
		}

		supplyOps = append(supplyOps, supplyOp)
	}

	if len(supplyOps) == 0 {
		return supplyOps, 404, fmt.Errorf("nenhuma operação encontrada")
	}

	return supplyOps, 200, nil
}

func (supplyop *SupplyOperationDB) ListAllSupplyOperationOfMonth(investmentId string, month int) ([]entity.SupplyOperation, int, error) {
	var supplyOps []entity.SupplyOperation
	stmt, err := supplyop.DB.Prepare("SELECT category, stockcode, quantity, supplyprice, value, supplydate FROM supplyoperation WHERE investimentid = $1 AND date_part('month', supplydate) = $2")

	if err != nil {
		fmt.Println(err.Error())
		return supplyOps, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(investmentId, month)

	if err != nil {
		fmt.Println(err.Error())
		return supplyOps, 500, err
	}

	for rows.Next() {
		var supplyOp entity.SupplyOperation

		err = rows.Scan(&supplyOp.Category, &supplyOp.StockCode, &supplyOp.Quantity, &supplyOp.SupplyPrice, &supplyOp.Value, &supplyOp.SupplyDate)

		if err != nil {
			fmt.Println(err.Error())
			return supplyOps, 404, nil
		}

		supplyOps = append(supplyOps, supplyOp)
	}

	if len(supplyOps) == 0 {
		return supplyOps, 404, fmt.Errorf("nenhuma operação encontrada")
	}
	return supplyOps, 200, nil
}
