package dto

import (
	"github.com/google/uuid"
)

type PeopleDto struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Phone string    `json:"phone"`
}
