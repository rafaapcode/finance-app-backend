package database

import (
	"context"
	"database/sql"
	"testing"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
	"github.com/rafaapcode/finance-app-backend/test"
	"github.com/stretchr/testify/assert"
)

var db *sql.DB

var defaultContext context.Context = context.Background()
var testsCtx context.Context

type UserId string

var idUser UserId = "userId"

func init() {
	database, err := test.Initialize()
	if err != nil {
		panic(err)
	}

	db = database
}

func TestCreateUser(t *testing.T) {
	user, err := entity.NewUser("Rafael", "rafa@gmail.com", "http://localhost.com")
	assert.NoError(t, err)

	userDb := NewUserDb(db)
	status, err := userDb.CreateUser(user)
	assert.NoError(t, err)
	assert.Equal(t, 201, status)

	testsCtx = context.WithValue(defaultContext, idUser, user.Id.String())
}

func TestGetUserById(t *testing.T) {
	userDb := NewUserDb(db)
	idOfUser := testsCtx.Value(idUser).(string)
	user, status, err := userDb.GetUser(idOfUser)

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEmpty(t, user)
	assert.Equal(t, "Rafael", user.Nome)
	assert.Equal(t, "rafa@gmail.com", user.Email)
	assert.Equal(t, "http://localhost.com", user.PhotoUrl)
}

func TestGetUserByEmail(t *testing.T) {
	userDb := NewUserDb(db)
	idOfUser := testsCtx.Value(idUser).(string)
	user, status, err := userDb.GetUser(idOfUser)
	assert.NoError(t, err)
	assert.Equal(t, 200, status)

	user, status, err = userDb.GetUserByEmail(user.Email)

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotEmpty(t, user)
	assert.Equal(t, "Rafael", user.Nome)
	assert.Equal(t, "rafa@gmail.com", user.Email)
	assert.Equal(t, "http://localhost.com", user.PhotoUrl)
}

func TestGetUserByInvalidEmail(t *testing.T) {
	userDb := NewUserDb(db)
	idOfUser := testsCtx.Value(idUser).(string)
	user, status, err := userDb.GetUser(idOfUser)
	assert.NoError(t, err)
	assert.Equal(t, 200, status)

	user, status, err = userDb.GetUserByEmail("rafa@@gmail.com")

	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Nil(t, user)
}

func TestGetUserByInvalidId(t *testing.T) {
	userDb := NewUserDb(db)
	user, status, err := userDb.GetUser("0194db6f-2570-716e-b9f0-1af62eddf56")
	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Nil(t, user)
}

func TestUpdateUser(t *testing.T) {
	userDb := NewUserDb(db)
	idOfUser := testsCtx.Value(idUser).(string)
	user, _, err := userDb.GetUser(idOfUser)
	assert.NoError(t, err)

	user.PhotoUrl = "http://novaUrl.com.br"

	status, err := userDb.UpdateUser(user)

	assert.NoError(t, err)
	assert.Equal(t, 200, status)

	user, _, err = userDb.GetUser(idOfUser)
	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, "http://novaUrl.com.br", user.PhotoUrl)
}

func TestDeleteUser(t *testing.T) {
	userDb := NewUserDb(db)
	idOfUser := testsCtx.Value(idUser).(string)
	res, status, err := userDb.DeleteUser(idOfUser)
	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, "User deleted with success", res)

	user, status, err := userDb.GetUser(idOfUser)
	assert.Error(t, err)
	assert.Equal(t, 404, status)
	assert.Nil(t, user)
}
