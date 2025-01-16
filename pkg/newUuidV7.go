package pkg

import "github.com/google/uuid"

func NewUUIDV7() string {
	uuidv7, err := uuid.NewV7()
	if err != nil {
		return ""
	}
	return uuidv7.String()
}
