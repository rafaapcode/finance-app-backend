package database

import (
	"database/sql"
	"fmt"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
)

type GoalsDb struct {
	DB *sql.DB
}

func NewGoalsDB(db *sql.DB) *GoalsDb {
	return &GoalsDb{
		DB: db,
	}
}

func (g *GoalsDb) CreateGoal(goals *entity.Goals) (int, error) {
	stmt, err := g.DB.Prepare(`
		INSERT INTO goals VALUES (
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

	_, err = stmt.Exec(goals.Id.String(), goals.Userid, goals.Category, goals.Percentage, goals.CreatedAt)

	if err != nil {
		fmt.Println(err.Error())
		return 500, err
	}

	return 201, nil
}

func (g *GoalsDb) UpdateGoal(id string, newPercentage float64) (int, error) {
	stmt, err := g.DB.Prepare("UPDATE goals SET percentage = $1 WHERE id = $2")

	if err != nil {
		fmt.Println(err.Error())
		return 500, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newPercentage, id)

	if err != nil {
		fmt.Println(err.Error())
		return 404, err
	}

	return 200, nil
}

func (g *GoalsDb) DeleteGoal(goalId string) (int, error) {
	stmt, err := g.DB.Prepare("DELETE FROM goals WHERE id = $1")

	if err != nil {
		fmt.Println(err.Error())
		return 500, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(goalId)

	if err != nil {
		fmt.Println(err.Error())
		return 404, err
	}

	return 200, nil
}

func (g *GoalsDb) ListAllGoals(userId string) ([]entity.Goals, int, error) {
	var allGoals []entity.Goals
	stmt, err := g.DB.Prepare("SELECT category, percentage FROM goals WHERE userid = $1 LIMIT 10")

	if err != nil {
		fmt.Println(err.Error())
		return allGoals, 500, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(userId)

	if err != nil {
		fmt.Println(err.Error())
		return allGoals, 404, err
	}

	for rows.Next() {
		var goal entity.Goals

		err = rows.Scan(&goal.Category, &goal.Percentage)
		if err != nil {
			return allGoals, 404, err
		}

		allGoals = append(allGoals, goal)
	}

	if len(allGoals) == 0 {
		return allGoals, 404, fmt.Errorf("nenhum objetivo encontrado")
	}

	return allGoals, 200, nil
}
