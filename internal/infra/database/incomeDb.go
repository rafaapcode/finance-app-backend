package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
)

type IncomeDb struct {
	DB *sql.DB
}

func NewIncomeDB(db *sql.DB) *IncomeDb {
	return &IncomeDb{
		DB: db,
	}
}

func (inc *IncomeDb) CreateIncome(income *entity.Income) (int, error) {
	stmt, err := inc.DB.Prepare(`
		INSERT INTO income VALUES(
			$1,
			$2,
			$3,
			$4
		);
	`)

	if err != nil {
		fmt.Println(err.Error())
		return 500, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(income.Id.String(), income.Userid, income.Value, income.CreatedAt)

	if err != nil {
		fmt.Println(err.Error())
		return 500, err
	}

	return 201, nil
}

func (inc *IncomeDb) GetIncomeValueByUserId(userId string) (float64, int, error) {
	var value float64
	stmt, err := inc.DB.Prepare("SELECT value FROM income WHERE userid = $1")

	if err != nil {
		fmt.Println(err.Error())
		return 0.0, 500, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(userId).Scan(&value)

	if err != nil {
		return 0.0, 404, errors.New("user not found")
	}

	return value, 200, nil
}

func (inc *IncomeDb) DeleteIncome(id string) (int, error) {
	stmt, err := inc.DB.Prepare("DELETE FROM income WHERE id = $1")

	if err != nil {
		fmt.Println(err.Error())
		return 500, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(id)

	if err != nil {
		fmt.Println(err.Error())
		return 404, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil || rowsAffected == 0 {
		return 404, errors.New("income not found")
	}

	return 200, nil
}

func (inc *IncomeDb) UpdateIncome(userId string, newValue float64) (int, error) {
	stmt, err := inc.DB.Prepare("UPDATE income SET value = $1 WHERE userid = $2")

	if err != nil {
		fmt.Println(err.Error())
		return 500, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(newValue, userId)

	if err != nil {
		fmt.Println(err.Error())
		return 404, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil || rowsAffected == 0 {
		return 404, errors.New("income of user not found")
	}

	return 200, nil
}

func (inc *IncomeDb) GetIncomeByUserId(userId string) (bool, int, error) {
	var userid string
	stmt, err := inc.DB.Prepare("SELECT userid FROM income WHERE userid = $1")

	if err != nil {
		fmt.Println(err.Error())
		return false, 500, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(userId).Scan(&userid)

	if err != nil {
		return false, 404, errors.New("income of user not found")
	}

	return true, 200, nil
}
