package database

import (
	"database/sql"
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
		return nil, 404, err
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
		return 404, err
	}

	return 200, nil
}
