package database

import (
	"log"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
	"gorm.io/gorm"
)

type UserDb struct {
	DB *gorm.DB
}

func NewUserDb(db *gorm.DB) *UserDb {
	return &UserDb{
		DB: db,
	}
}

func (userDb *UserDb) CreateUser(user *entity.User) (int, error) {
	err := userDb.DB.Create(user).Error
	if err != nil {
		log.Fatal(err.Error())
		return 500, err
	}
	return 201, nil
}

func (userDb *UserDb) GetUser(id string) (*entity.User, int, error) {
	var user entity.User
	err := userDb.DB.First(&user, "id = ?", id).Error

	if err != nil {
		log.Fatal(err.Error())
		return nil, 404, err
	}

	return &user, 200, nil
}

func (userDb *UserDb) DeleteUser(id string) (string, int, error) {
	_, status, err := userDb.GetUser(id)
	if status != 200 {
		return "", 404, err
	}

	return "User deleted with success", 200, nil
}

func (userDb *UserDb) UpdateUser(newUserData *entity.User) (int, error) {
	_, status, err := userDb.GetUser(newUserData.Id.String())

	if status != 200 {
		return 404, err
	}

	err = userDb.DB.Save(newUserData).Error

	return 200, err
}
