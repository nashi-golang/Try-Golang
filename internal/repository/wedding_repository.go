package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"try-golang/internal/models"
)

type WeddingRepository interface {
	CreateWedding(wedding *models.Wedding) error
	GetWeddingByID(id uuid.UUID) (*models.Wedding, error)
	GetAllWeddings() ([]models.Wedding, error)
	UpdateWedding(wedding *models.Wedding) error
	DeleteWedding(id uuid.UUID) error
}

type weddingRepositoryImpl struct {
	db *gorm.DB
}

func NewWeddingRepository(db *gorm.DB) WeddingRepository {
	return &weddingRepositoryImpl{
		db: db,
	}
}
func (r *weddingRepositoryImpl) CreateWedding(wedding *models.Wedding) error {
	return r.db.Create(wedding).Error
}

func (r *weddingRepositoryImpl) GetWeddingByID(id uuid.UUID) (*models.Wedding, error) {
	var wedding models.Wedding
	err := r.db.
		Preload("Peoples").
		Preload("Photos").
		First(&wedding, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &wedding, nil
}

func (r *weddingRepositoryImpl) GetAllWeddings() ([]models.Wedding, error) {
	var weddings []models.Wedding
	err := r.db.
		Preload("Peoples").
		Preload("Photos").
		Find(&weddings).Error
	if err != nil {
		return nil, err
	}
	return weddings, nil
}

func (r *weddingRepositoryImpl) UpdateWedding(wedding *models.Wedding) error {
	return r.db.Save(wedding).Error
}

func (r *weddingRepositoryImpl) DeleteWedding(id uuid.UUID) error {
	return r.db.Delete(&models.Wedding{}, id).Error
}
