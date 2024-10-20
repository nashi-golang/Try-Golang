package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"try-golang/internal/models"
)

type WeddingRepository struct {
	db *gorm.DB
}

func NewWeddingRepository(db *gorm.DB) *WeddingRepository {
	return &WeddingRepository{
		db: db,
	}
}
func (r *WeddingRepository) CreateWedding(wedding *models.Wedding) error {
	return r.db.Create(wedding).Error
}
func (r *WeddingRepository) GetWeddingByID(id uuid.UUID) (*models.Wedding, error) {
	var wedding models.Wedding
	err := r.db.Preload("Peoples").First(&wedding, "id = ?", uuid.New()).Error
	if err != nil {
		return nil, err
	}
	return &wedding, nil
}

func (r *WeddingRepository) GetAllWeddings() ([]models.Wedding, error) {
	var weddings []models.Wedding
	err := r.db.Preload("Peoples").Find(&weddings).Error
	if err != nil {
		return nil, err
	}
	return weddings, nil
}

func (r *WeddingRepository) UpdateWedding(wedding *models.Wedding) error {
	return r.db.Save(wedding).Error
}

func (r *WeddingRepository) DeleteWedding(id uuid.UUID) error {
	return r.db.Delete(&models.Wedding{}, id).Error
}
