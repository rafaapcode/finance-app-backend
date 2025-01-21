package user

import (
	"fmt"

	"github.com/rafaapcode/finance-app-backend/api/model"
	"github.com/rafaapcode/finance-app-backend/pkg"
)

type UserController struct {
	Name     string
	Email    string
	PhotoUrl string
	Repo     UserRepository
}

func (controllerUser UserController) CreateUser() (model.User, int, error) {
	userId := pkg.NewUUIDV7()
	if userId == "" {
		return model.User{}, 500, fmt.Errorf("error to create an ID")
	}

	// Insere no DB
	return model.User{}, 201, nil
}

func (controllerUser UserController) GetUser() (model.User, int, error) {
	fmt.Print(controllerUser.Name)

	// Busca no DB
	return model.User{}, 201, nil
}

func (controllerUser UserController) DeleteUser() (string, int, error) {
	fmt.Print(controllerUser.Name)

	// Deleta no DB
	return "User deleted with success", 201, nil
}

func (controllerUser UserController) UpdateUser(newName, newEmail string) (model.User, int, error) {
	fmt.Print(newName)

	// Atualiza no DB
	return model.User{}, 201, nil
}
