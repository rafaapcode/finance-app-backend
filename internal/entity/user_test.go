package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAUser(t *testing.T) {
	user, err := NewUser("Rafae", "rafa@gmail.com", "http://rafa.photo.com")

	assert.Nil(t, err)

	err = user.Validate()

	assert.Nil(t, err)
	assert.Equal(t, "Rafae", user.Nome)
	assert.Equal(t, "rafa@gmail.com", user.Email)
	assert.Equal(t, "http://rafa.photo.com", user.PhotoUrl)
}

func TestNameIsRequired(t *testing.T) {
	user, err := NewUser("", "rafa@gmail.com", "http://rafa.photo.com")
	assert.Nil(t, err)

	err = user.Validate()

	assert.Equal(t, ErrNameIsRequired, err)
}

func TestNameIsInvalid(t *testing.T) {
	user, err := NewUser("ra", "rafa@gmail.com", "http://rafa.photo.com")
	assert.Nil(t, err)

	err = user.Validate()

	assert.Equal(t, ErrNameIsInvalid, err)
}

func TestEmailIsRequired(t *testing.T) {
	user, err := NewUser("rafael", "", "http://rafa.photo.com")
	assert.Nil(t, err)

	err = user.Validate()

	assert.Equal(t, ErrEmailIsRequired, err)
}

func TestEmailIsInvalid(t *testing.T) {
	user, err := NewUser("rafael", "rafa.com", "http://rafa.photo.com")
	assert.Nil(t, err)

	err = user.Validate()

	assert.Equal(t, ErrEmailIsInvalid, err)
}

func TestPhotoUrlIsRequired(t *testing.T) {
	user, err := NewUser("rafael", "rafa@gmail.com", "")
	assert.Nil(t, err)

	err = user.Validate()

	assert.Equal(t, ErrPhotoUrlIsRequired, err)
}

func TestPhotoUrlIsInvalid(t *testing.T) {
	user, err := NewUser("rafael", "rafa@gmail.com", "rafa.photo.com")
	assert.Nil(t, err)

	err = user.Validate()

	assert.Equal(t, ErrPhotoUrlIsInvalid, err)
}
