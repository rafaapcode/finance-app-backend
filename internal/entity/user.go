package entity

import (
	"regexp"

	"github.com/rafaapcode/finance-app-backend/pkg"
)

type User struct {
	Id         pkg.ID       `json:"id"`
	Nome       string       `json:"nome"`
	Email      string       `json:"email"`
	PhotoUrl   string       `json:"photoUrl"`
	Income     []Income     `json:"income"`
	Outcome    []Outcome    `json:"outcome"`
	Investment []Investment `json:"investment"`
	Goals      []Goals      `json:"goals"`
}

func NewUser(nome, email, photoUrl string) (*User, error) {
	id, err := pkg.NewUUID()

	if err != nil {
		return nil, err
	}

	return &User{
		Id:       id,
		Nome:     nome,
		Email:    email,
		PhotoUrl: photoUrl,
	}, nil
}

func isValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(regex)

	return re.MatchString(email)
}

func isValidURL(url string) bool {
	regex := `^(http?|https?|ftp):\/\/[^\s/$.?#].[^\s]*$`

	re := regexp.MustCompile(regex)

	return re.MatchString(url)
}

func (u *User) Validate() error {
	if u.Id.String() == "" {
		return ErrIdIdRequired
	}
	if _, err := pkg.ParseID(u.Id.String()); err != nil {
		return ErrIdInvalidId
	}
	if u.Nome == "" {
		return ErrNameIsRequired
	}
	if len(u.Nome) <= 3 {
		return ErrNameIsInvalid
	}
	if u.Email == "" {
		return ErrEmailIsRequired
	}
	if !isValidEmail(u.Email) {
		return ErrEmailIsInvalid
	}
	if u.PhotoUrl == "" {
		return ErrPhotoUrlIsRequired
	}
	if !isValidURL(u.PhotoUrl) {
		return ErrPhotoUrlIsInvalid
	}
	return nil
}
