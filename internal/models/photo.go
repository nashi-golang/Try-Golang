package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Photo struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	WeddingID uuid.UUID
	Status    string
	Url       string
}

func (photo *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	photo.ID = uuid.New()
	return
}
