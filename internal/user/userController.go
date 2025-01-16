package user

import (
	"fmt"

	"github.com/rafaapcode/finance-app-backend/pkg"
	"gorm.io/gorm"
)

type UserController struct {
	Name  string
	Email string
	DB    *gorm.DB
}

func (controllerUser UserController) CreateUser() (string, int, error) {
	userId := pkg.NewUUIDV7()
	if userId == "" {
		return "", 500, fmt.Errorf("error to create an ID")
	}

	// Insere no DB
	return userId, 201, nil
}

func (controllerUser UserController) GetUser() (string, int, error) {
	fmt.Print(controllerUser.Name)

	// Busca no DB
	return "User returned", 201, nil
}

func (controllerUser UserController) DeleteUser() (string, int, error) {
	fmt.Print(controllerUser.Name)

	// Deleta no DB
	return "User deleted with success", 201, nil
}

func (controllerUser UserController) UpdateUser(newName, newEmail string) (string, int, error) {
	fmt.Print(newName)

	// Atualiza no DB
	return "User updated with success", 201, nil
}
