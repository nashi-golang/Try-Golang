package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"try-golang/internal/models"
)

type PeopleRepository struct {
	db *gorm.DB
}

func NewPeopleRepository(db *gorm.DB) *PeopleRepository {
	return &PeopleRepository{
		db: db,
	}
}
func (r *PeopleRepository) CreatePeople(People *models.People) error {
	return r.db.Create(People).Error
}
func (r *PeopleRepository) GetPeopleByID(id uuid.UUID) (*models.People, error) {
	var people models.People
	err := r.db.First(&people, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &people, nil
}
func (r *PeopleRepository) GetPeoplesByWeddingID(weddingId uuid.UUID) (*models.People, error) {
	var people models.People
	err := r.db.Find(&people, "weddingId = ?", weddingId).Error
	if err != nil {
		return nil, err
	}
	return &people, nil
}

func (r *PeopleRepository) UpdatePeople(people *models.People) error {
	return r.db.Save(people).Error
}

func (r *PeopleRepository) DeletePeople(id uuid.UUID) error {
	return r.db.Delete(&models.People{}, id).Error
}
