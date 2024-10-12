package dto

import "time"

type WeddingDto struct {
	ID            string    `json:"id"`
	StartDatetime time.Time `json:"start-datetime"`
	Location      string    `json:"location"`
	Groom         string    `json:"groom"`
	Bride         string    `json:"bridge"`
}
