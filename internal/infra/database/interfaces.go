package database

import "github.com/rafaapcode/finance-app-backend/internal/entity"

type UserInterface interface {
	CreateUser(user *entity.User) (int, error)
	GetUser(id string) (*entity.User, int, error)
	DeleteUser(id string) (string, int, error)
	UpdateUser(newUserData *entity.User) (int, error)
}

// type ProductInterface interface {
// 	Create(product *entity.Product) error
// 	FindAll(page, limit int, sort string) ([]*entity.Product, error)
// 	FindById(id string) (*entity.Product, error)
// 	Update(prduct *entity.Product) error
// 	Delete(id string) error
// }
