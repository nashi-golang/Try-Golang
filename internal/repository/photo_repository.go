package repository

import (
	"fmt"
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
func (r *PhotoRepository) GetPendingPhotosByWeddingID(weddingId uuid.UUID) ([]models.Photo, error) {
	var wedding models.Wedding

	if err := r.db.Preload("Photos", "status = ?", "Not Started").
		First(&wedding, "id = ?", weddingId).
		Error; err != nil {
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

func (r *PhotoRepository) UpdatePhotos(photos []models.Photo) error {
	if len(photos) == 0 {
		return nil // 업데이트할 레코드가 없으면 바로 반환
	}

	// SQL 쿼리 문자열과 파라미터 슬라이스 초기화
	sql := "UPDATE photos SET status = CASE"
	ids := make([]interface{}, 0, len(photos))

	for _, photo := range photos {
		if photo.ID == uuid.Nil {
			fmt.Println("Error: Photo ID is not set")
			continue
		}
		sql += " WHEN id = ? THEN ?"
		ids = append(ids, photo.ID, photo.Status)
	}

	sql += " ELSE status END WHERE id IN ("
	placeholders := ""
	for i, photo := range photos {
		if i > 0 {
			placeholders += ", "
		}
		placeholders += "?"
		ids = append(ids, photo.ID)
	}
	sql += placeholders + ")"

	// 업데이트 실행
	err := r.db.Exec(sql, ids...).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PhotoRepository) DeletePhoto(id uuid.UUID) error {
	return r.db.Delete(&models.Photo{}, id).Error
}
