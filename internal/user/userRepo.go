package user

import (
	"github.com/rafaapcode/finance-app-backend/api/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user model.User) (model.User, int, error)
	GetUser(name string) (model.User, int, error)
	DeleteUser(name string) (string, int, error)
	UpdateUser(name string, newData model.User) (model.User, int, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func (u UserRepo) CreateUser(user model.User) (model.User, int, error) {
	return model.User{}, 10, nil
}

func (u UserRepo) GetUser(name string) (model.User, int, error) {
	return model.User{}, 10, nil
}

func (u UserRepo) DeleteUser(name string) (string, int, error) {
	return "", 10, nil
}

func (u UserRepo) UpdateUser(newData model.User) (model.User, int, error) {
	return model.User{}, 10, nil
}
