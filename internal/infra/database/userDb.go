package database

import (
	"database/sql"
	"fmt"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
)

type UserDb struct {
	DB *sql.DB
}

func NewUserDb(db *sql.DB) *UserDb {
	return &UserDb{
		DB: db,
	}
}

func (userDb *UserDb) CreateUser(user *entity.User) (int, error) {
	stmt, err := userDb.DB.Prepare("INSERT INTO users VALUES ($1, $2, $3, $4)")

	if err != nil {
		fmt.Println(err.Error())
		return 500, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id.String(), user.Nome, user.Email, user.PhotoUrl)

	if err != nil {
		fmt.Println(err.Error())
		return 500, err
	}

	return 201, nil
}

func (userDb *UserDb) GetUser(id string) (*entity.User, int, error) {
	var user entity.User
	stmt, err := userDb.DB.Prepare("SELECT id, nome, email, photourl FROM users WHERE id = $1")

	if err != nil {
		fmt.Println(err.Error())
		return nil, 500, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&user.Id, &user.Nome, &user.Email, &user.PhotoUrl)

	if err != nil {
		fmt.Println(err.Error())
		return nil, 404, err
	}

	return &user, 200, nil
}

func (userDb *UserDb) DeleteUser(id string) (string, int, error) {
	stmt, err := userDb.DB.Prepare("DELETE FROM users where id = $1")

	if err != nil {
		fmt.Println(err.Error())
		return "", 500, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		fmt.Println(err.Error())
		return "", 404, err
	}

	return "User deleted with success", 200, nil
}

func (userDb *UserDb) UpdateUser(newUserData *entity.User) (int, error) {
	stmt, err := userDb.DB.Prepare("UPDATE users SET photourl = $1 WHERE id = $2")
	if err != nil {
		return 500, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(newUserData.PhotoUrl, newUserData.Id)

	if err != nil {
		return 404, err
	}

	return 200, err
}
