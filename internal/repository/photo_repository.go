package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"try-golang/internal/models"
)

type PhotoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *PhotoRepository {
	return &PhotoRepository{
		db: db,
	}
}
func (r *PhotoRepository) CreatePhoto(Photo *models.Photo) error {
	return r.db.Create(Photo).Error
}

func (r *PhotoRepository) GetPhotosByWeddingID(weddingId uuid.UUID) ([]models.Photo, error) {
	var wedding models.Wedding

	if err := r.db.Preload("Photos").First(&wedding, "id = ?", weddingId).Error; err != nil {
		return nil, err
	}
	return wedding.Photos, nil
}
func (r *PhotoRepository) GetAllPhotos() ([]models.Photo, error) {
	var Photos []models.Photo
	err := r.db.Preload("Peoples").Find(&Photos).Error
	if err != nil {
		return nil, err
	}
	return Photos, nil
}

func (r *PhotoRepository) UpdatePhoto(Photo *models.Photo) error {
	return r.db.Save(Photo).Error
}

func (r *PhotoRepository) DeletePhoto(id uuid.UUID) error {
	return r.db.Delete(&models.Photo{}, id).Error
}
