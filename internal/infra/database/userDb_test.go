package database

import (
	"database/sql"
	"testing"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
	"github.com/rafaapcode/finance-app-backend/test"
	"github.com/stretchr/testify/assert"
)

var db *sql.DB

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

	userDb.CreateUser(user)
}
