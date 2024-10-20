package dto

import (
	"github.com/google/uuid"
	"time"
)

type WeddingDto struct {
	ID            uuid.UUID   `json:"id"`
	StartDatetime time.Time   `json:"start-datetime"`
	Location      string      `json:"location"`
	Groom         uuid.UUID   `json:"groom"`
	Bride         uuid.UUID   `json:"bridge"`
	Peoples       []PeopleDto `json:"peoples"`
}
