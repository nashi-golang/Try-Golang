package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type People struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	WeddingID uuid.UUID
	Name      string
	Phone     string
}

func (people *People) BeforeCreate(tx *gorm.DB) (err error) {
	people.ID = uuid.New()
	return
}
