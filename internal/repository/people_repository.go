package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"try-golang/internal/models"
)

type PeopleRepository interface {
	CreatePeople(People *models.People) error
	GetPeopleByID(id uuid.UUID) (*models.People, error)
	UpdatePeople(People *models.People) error
	DeletePeople(id uuid.UUID) error
}

type peopleRepositoryImpl struct {
	db *gorm.DB
}

func NewPeopleRepository(db *gorm.DB) PeopleRepository {
	return &peopleRepositoryImpl{
		db: db,
	}
}
func (r *peopleRepositoryImpl) CreatePeople(People *models.People) error {
	return r.db.Create(People).Error
}
func (r *peopleRepositoryImpl) GetPeopleByID(id uuid.UUID) (*models.People, error) {
	var people models.People
	err := r.db.First(&people, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &people, nil
}
func (r *peopleRepositoryImpl) GetPeoplesByWeddingID(weddingId uuid.UUID) (*models.People, error) {
	var people models.People
	err := r.db.Find(&people, "weddingId = ?", weddingId).Error
	if err != nil {
		return nil, err
	}
	return &people, nil
}

func (r *peopleRepositoryImpl) UpdatePeople(people *models.People) error {
	return r.db.Save(people).Error
}

func (r *peopleRepositoryImpl) DeletePeople(id uuid.UUID) error {
	return r.db.Delete(&models.People{}, id).Error
}
