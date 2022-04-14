package entity

import (
	uuid "github.com/google/uuid"
)

//ID entity ID
type ID = uuid.UUID

//NewID create a new entity ID
func NewID() string {
	return uuid.New().String()
}
