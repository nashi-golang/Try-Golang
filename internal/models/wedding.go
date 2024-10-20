package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Wedding struct {
	ID            uuid.UUID `gorm:"type:char(36);primaryKey"`
	StartDateTime time.Time
	Location      string
	Bride         uuid.UUID
	Groom         uuid.UUID
	Peoples       []People `gorm:"foreignKey:WeddingID"`
}

func (wedding *Wedding) BeforeCreate(tx *gorm.DB) (err error) {
	wedding.ID = uuid.New()
	return
}
