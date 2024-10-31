package dto

import (
	"github.com/google/uuid"
)

type PhotoDto struct {
	ID     uuid.UUID `json:"id"`
	Status string
	Url    string
}
