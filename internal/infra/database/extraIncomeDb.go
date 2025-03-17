package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
)

type ExtraIncomeDb struct {
	DB *sql.DB
}

func NewExtraIncomeDB(db *sql.DB) *ExtraIncomeDb {
	return &ExtraIncomeDb{
		DB: db,
	}
}

func (extInc *ExtraIncomeDb) CreateExtraIncome(extraIncome *entity.ExtraIncome) (int, error) {
	stmt, err := extInc.DB.Prepare(`
		INSERT INTO extraincome VALUES(
			$1,
			$2,
			$3,
			$4,
			$5
		);
	`)

	if err != nil {
		fmt.Println(err.Error())
		return 500, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(extraIncome.Id.String(), extraIncome.Userid, extraIncome.Category, extraIncome.Value, extraIncome.Date)

	if err != nil {
		fmt.Println(err.Error())
		return 500, err
	}

	return 201, nil
}

func (extInc *ExtraIncomeDb) GetExtraIncomeById(id string) (*entity.ExtraIncome, int, error) {
	var extIncome entity.ExtraIncome
	stmt, err := extInc.DB.Prepare("SELECT id, userid, category, value, date FROM extraincome WHERE id = $1")

	if err != nil {
		fmt.Println(err.Error())
		return nil, 500, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&extIncome.Id, &extIncome.Userid, &extIncome.Category, &extIncome.Value, &extIncome.Date)

	if err != nil {
		fmt.Println(err.Error())
		return nil, 404, errors.New("Extra income not found")
	}

	return &extIncome, 200, nil
}

func (extInc *ExtraIncomeDb) DeleteExtraIncome(id string) (int, error) {
	stmt, err := extInc.DB.Prepare("DELETE FROM extraincome WHERE id = $1")

	if err != nil {
		fmt.Println(err.Error())
		return 500, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return 404, errors.New("Extra income not found")
	}

	return 200, nil
}

func (extInc *ExtraIncomeDb) GetAllExtraIncomeOfMonth(month int, userId string) ([]entity.ExtraIncome, int, error) {
	var extraIncomeOfMonth []entity.ExtraIncome
	stmt, err := extInc.DB.Prepare("SELECT id, userid, category, value, date FROM extraincome WHERE date_part('month', date) = $1 AND userid = $2")

	if err != nil {
		return extraIncomeOfMonth, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(month, userId)

	if err != nil {
		fmt.Println(err.Error())
		return extraIncomeOfMonth, 404, errors.New("extra income not found")
	}

	for rows.Next() {
		var extraInc entity.ExtraIncome

		err = rows.Scan(&extraInc.Id, &extraInc.Userid, &extraInc.Category, &extraInc.Value, &extraInc.Date)

		if err != nil {
			return extraIncomeOfMonth, 500, nil
		}

		extraIncomeOfMonth = append(extraIncomeOfMonth, extraInc)
	}

	if len(extraIncomeOfMonth) == 0 {
		return extraIncomeOfMonth, 404, errors.New("Extra income not found")
	}

	return extraIncomeOfMonth, 200, nil
}

func (extInc *ExtraIncomeDb) GetTotalValueOfExtracIncomeOfTheMonth(month int, userId string) (float64, int, error) {
	var total sql.NullFloat64
	stmt, err := extInc.DB.Prepare("SELECT sum(value) FROM extraincome WHERE date_part('month', date) = $1 AND userid = $2")

	if err != nil {
		return 0.0, 500, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(month, userId).Scan(&total)

	if err != nil {
		return 0.0, 404, errors.New("user not found")
	}

	fmt.Println("res", total)
	return total.Float64, 200, nil
}
