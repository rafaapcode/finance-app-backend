package database

import (
	"database/sql"
	"fmt"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
)

type OutcomeDb struct {
	DB *sql.DB
}

func NewOutcomeDb(db *sql.DB) *OutcomeDb {
	return &OutcomeDb{
		DB: db,
	}
}

func (outDb *OutcomeDb) CreateOutcome(outcome *entity.Outcome) (int, error) {
	stmt, err := outDb.DB.Prepare(`
		INSERT INTO outcome VALUES(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7, 
			$8,
			$9
		)
	`)

	if err != nil {
		fmt.Println(err.Error())
		return 500, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(outcome.Id.String(), outcome.Userid, outcome.Type, outcome.Category, outcome.Value, outcome.PaymentMethod, outcome.Notification, outcome.ExpireDate, outcome.CreatedAt)

	if err != nil {
		fmt.Println(err.Error())
		return 500, err
	}

	return 201, nil
}

func (outDb *OutcomeDb) GetOutcomeById(id string) (*entity.Outcome, int, error) {
	var outcome entity.Outcome

	stmt, err := outDb.DB.Prepare("SELECT id, userid, type, category, value, paymentmethod, notification, expiredate, createdat FROM outcome WHERE id = $1")

	if err != nil {
		fmt.Println(err.Error())
		return nil, 500, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&outcome.Id, &outcome.Userid, &outcome.Type, &outcome.Category, &outcome.Value, &outcome.PaymentMethod, &outcome.Notification, &outcome.ExpireDate, &outcome.CreatedAt)

	if err != nil {
		fmt.Println(err.Error())
		return nil, 404, err
	}

	return &outcome, 200, nil
}

// A testar
func (outDb *OutcomeDb) GetAllOutcomeOfMonth(month int, userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome
	stmt, err := outDb.DB.Prepare("SELECT id, userid, type, category, value, paymentmethod, notification,expiredate, createdat FROM outcome WHERE date_part('month', createdat) = $1 AND userid = $2")

	if err != nil {
		fmt.Println(err.Error())
		return nil, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(month, userId)

	if err != nil {
		fmt.Println(err.Error())
		return nil, 404, err
	}

	defer rows.Close()

	for rows.Next() {
		var outcome entity.Outcome
		err = rows.Scan(&outcome.Id, &outcome.Userid, &outcome.Type, &outcome.Category, &outcome.Value, &outcome.PaymentMethod, &outcome.Notification, &outcome.ExpireDate, &outcome.CreatedAt)
		if err != nil {
			fmt.Println(err.Error())
			return nil, 500, err
		}
		outcomes = append(outcomes, outcome)
	}

	return outcomes, 200, nil
}

func (outDb *OutcomeDb) GetAllFixedOutcome(userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome
	stmt, err := outDb.DB.Prepare("SELECT id, userid, type, category, value, paymentmethod, notification,expiredate, createdat FROM outcome WHERE type = 'fixo' AND userid = $1")

	if err != nil {
		fmt.Println(err.Error())
		return outcomes, 500, nil
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId)

	if err != nil {
		fmt.Println(err.Error())
		return outcomes, 404, nil
	}
	defer rows.Close()

	for rows.Next() {
		var outcome entity.Outcome

		err = rows.Scan(&outcome.Id, &outcome.Userid, &outcome.Type, &outcome.Category, &outcome.Value, &outcome.PaymentMethod, &outcome.Notification, &outcome.ExpireDate, &outcome.CreatedAt)
		if err != nil {
			fmt.Println(err.Error())
			return nil, 404, err
		}

		outcomes = append(outcomes, outcome)
	}

	if len(outcomes) == 0 {
		return nil, 404, fmt.Errorf("nenhuma saída encontrada")
	}

	return outcomes, 200, nil
}

func (outDb *OutcomeDb) GetAllOutcomeByCategory(category string, userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome

	stmt, err := outDb.DB.Prepare("SELECT id, userid, type, category, value, paymentmethod, notification,expiredate, createdat FROM outcome WHERE category = $1 AND userid = $2")

	if err != nil {
		fmt.Println(err.Error())
		return outcomes, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(category, userId)

	if err != nil {
		fmt.Println(err.Error())
		return outcomes, 404, err
	}
	defer rows.Close()

	for rows.Next() {
		var outcome entity.Outcome
		err = rows.Scan(&outcome.Id, &outcome.Userid, &outcome.Type, &outcome.Category, &outcome.Value, &outcome.PaymentMethod, &outcome.Notification, &outcome.ExpireDate, &outcome.CreatedAt)
		if err != nil {
			fmt.Println(err.Error())
			return outcomes, 500, err
		}
		outcomes = append(outcomes, outcome)
	}

	if len(outcomes) == 0 {
		return outcomes, 404, fmt.Errorf("nenhuma saída encontrada")
	}

	return outcomes, 200, nil
}

func (outDb *OutcomeDb) GetAllOutcomeByPaymentMethod(method string, userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome

	stmt, err := outDb.DB.Prepare("SELECT id, userid, type, category, value, paymentmethod, notification,expiredate, createdat FROM outcome WHERE paymentmethod = $1 AND userid = $2")

	if err != nil {
		fmt.Println(err.Error())
		return outcomes, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(method, userId)

	if err != nil {
		fmt.Println(err.Error())
		return outcomes, 404, err
	}
	defer rows.Close()

	for rows.Next() {
		var outcome entity.Outcome
		err = rows.Scan(&outcome.Id, &outcome.Userid, &outcome.Type, &outcome.Category, &outcome.Value, &outcome.PaymentMethod, &outcome.Notification, &outcome.ExpireDate, &outcome.CreatedAt)
		if err != nil {
			fmt.Println(err.Error())
			return outcomes, 500, err
		}
		outcomes = append(outcomes, outcome)
	}

	if len(outcomes) == 0 {
		return outcomes, 404, fmt.Errorf("nenhuma saída encontrada")
	}

	return outcomes, 200, nil
}

func (outDb *OutcomeDb) GetAllOutcomeByType(typeOutcome string, userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome

	stmt, err := outDb.DB.Prepare("SELECT id, userid, type, category, value, paymentmethod, notification,expiredate, createdat FROM outcome WHERE type = $1 AND userid = $2")

	if err != nil {
		fmt.Println(err.Error())
		return outcomes, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(typeOutcome, userId)

	if err != nil {
		fmt.Println(err.Error())
		return outcomes, 404, err
	}
	defer rows.Close()

	for rows.Next() {
		var outcome entity.Outcome
		err = rows.Scan(&outcome.Id, &outcome.Userid, &outcome.Type, &outcome.Category, &outcome.Value, &outcome.PaymentMethod, &outcome.Notification, &outcome.ExpireDate, &outcome.CreatedAt)
		if err != nil {
			fmt.Println(err.Error())
			return outcomes, 500, err
		}
		outcomes = append(outcomes, outcome)
	}

	if len(outcomes) == 0 {
		return outcomes, 404, fmt.Errorf("nenhuma saída encontrada")
	}

	return outcomes, 200, nil
}

func (outDb *OutcomeDb) GetOutcomeAboutToExpire(daysToExpire int, userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome

	stmt, err := outDb.DB.Prepare("SELECT id, userid, type, category, value, paymentmethod, notification, expiredate, createdat FROM outcome WHERE expiredate BETWEEN CURRENT_DATE AND CURRENT_DATE + INTERVAL '$1 days' AND userid = $2")

	if err != nil {
		fmt.Println(err.Error())
		return outcomes, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(daysToExpire, userId)

	if err != nil {
		fmt.Println(err.Error())
		return outcomes, 404, err
	}
	defer rows.Close()

	for rows.Next() {
		var outcome entity.Outcome
		err = rows.Scan(&outcome.Id, &outcome.Userid, &outcome.Type, &outcome.Category, &outcome.Value, &outcome.PaymentMethod, &outcome.Notification, &outcome.ExpireDate, &outcome.CreatedAt)
		if err != nil {
			fmt.Println(err.Error())
			return outcomes, 500, err
		}
		outcomes = append(outcomes, outcome)
	}

	if len(outcomes) == 0 {
		return outcomes, 404, fmt.Errorf("nenhuma saída encontrada")
	}

	return outcomes, 200, nil
}

func (outDb *OutcomeDb) GetOutcomeLessThan(value float64, userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome

	stmt, err := outDb.DB.Prepare("SELECT id, userid, type, category, value, paymentmethod, notification, expiredate, createdat FROM outcome WHERE value < $1 AND userid = $2")

	if err != nil {
		fmt.Println(err.Error())
		return outcomes, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(value, userId)

	if err != nil {
		fmt.Println(err.Error())
		return outcomes, 404, err
	}
	defer rows.Close()

	for rows.Next() {
		var outcome entity.Outcome
		err = rows.Scan(&outcome.Id, &outcome.Userid, &outcome.Type, &outcome.Category, &outcome.Value, &outcome.PaymentMethod, &outcome.Notification, &outcome.ExpireDate, &outcome.CreatedAt)
		if err != nil {
			fmt.Println(err.Error())
			return outcomes, 500, err
		}
		outcomes = append(outcomes, outcome)
	}

	if len(outcomes) == 0 {
		return outcomes, 404, fmt.Errorf("nenhuma saída encontrada")
	}

	return outcomes, 200, nil
}

func (outDb *OutcomeDb) GetOutcomeHigherThan(value float64, userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome

	stmt, err := outDb.DB.Prepare("SELECT id, userid, type, category, value, paymentmethod, notification, expiredate, createdat FROM outcome WHERE value > $1 AND userid = $2")

	if err != nil {
		fmt.Println(err.Error())
		return outcomes, 500, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(value, userId)

	if err != nil {
		fmt.Println(err.Error())
		return outcomes, 404, err
	}
	defer rows.Close()

	for rows.Next() {
		var outcome entity.Outcome
		err = rows.Scan(&outcome.Id, &outcome.Userid, &outcome.Type, &outcome.Category, &outcome.Value, &outcome.PaymentMethod, &outcome.Notification, &outcome.ExpireDate, &outcome.CreatedAt)
		if err != nil {
			fmt.Println(err.Error())
			return outcomes, 500, err
		}
		outcomes = append(outcomes, outcome)
	}

	if len(outcomes) == 0 {
		return outcomes, 404, fmt.Errorf("nenhuma saída encontrada")
	}

	return outcomes, 200, nil
}

func (outDb *OutcomeDb) DeleteOutcome(id string) (string, int, error) {
	stmt, err := outDb.DB.Prepare("DELETE FROM outcome WHERE id = $1")

	if err != nil {
		fmt.Println(err.Error())
		return "Erro ao deletar a saída", 500, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Println(err.Error())
		return "Saída não encontrada", 404, err
	}

	return "Saída deletada com sucesso", 200, nil
}
