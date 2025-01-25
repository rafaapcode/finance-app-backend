package pkg

import "github.com/google/uuid"

type ID = uuid.UUID

func NewUUID() (ID, error) {
	uuidv7, err := uuid.NewV7()
	return ID(uuidv7), err
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}
