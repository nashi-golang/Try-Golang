package services

import (
	"errors"
	"github.com/google/uuid"
	"try-golang/internal/models"
	"try-golang/internal/repository"
	"try-golang/pkg/dto"
)

type WeddingService struct {
	weddingRepo *repository.WeddingRepository
	peopleRepo  *repository.PeopleRepository
	photoRepo   *repository.PhotoRepository
}

func NewWeddingService(weddingRepo *repository.WeddingRepository, peopleRepo *repository.PeopleRepository, photoRepo *repository.PhotoRepository) *WeddingService {
	return &WeddingService{
		weddingRepo: weddingRepo,
		peopleRepo:  peopleRepo,
		photoRepo:   photoRepo,
	}
}

func (s *WeddingService) CreateWedding(dto dto.WeddingDto) (string, error) {
	newWedding := NewWeddingFromDto(dto)
	err := s.weddingRepo.CreateWedding(newWedding)
	if err != nil {
		return "", err
	}
	return "Wedding created: " + newWedding.ID.String(), nil
}

func (s *WeddingService) GetWeddings() []dto.WeddingDto {
	var weddingDtos []dto.WeddingDto

	weddings, err := s.weddingRepo.GetAllWeddings()
	if err != nil {
		return nil
	}

	for i := range weddings {
		weddingDto := NewWeddingDtoFromModel(weddings[i])
		weddingDtos = append(weddingDtos, *weddingDto)
	}
	return weddingDtos
}

func (s *WeddingService) GetWeddingById(id uuid.UUID) (*dto.WeddingDto, error) {

	wedding, err := s.weddingRepo.GetWeddingByID(id)
	if err != nil {
		return nil, err
	}
	weddingDto := NewWeddingDtoFromModel(*wedding)

	return weddingDto, err
}

func (s *WeddingService) UpdateWedding(id uuid.UUID, updatedWeddingData dto.WeddingDto) (string, error) {
	wedding, err := s.weddingRepo.GetWeddingByID(id)
	if err != nil {
		return "", errors.New("wedding not found")
	}

	s.updateWeddingFields(wedding, updatedWeddingData)

	err = s.weddingRepo.UpdateWedding(wedding)
	if err != nil {
		return "", errors.New("failed to update wedding")
	}

	return "Wedding updated successfully", nil
}

func (s *WeddingService) DeleteWedding(id uuid.UUID) error {
	return s.weddingRepo.DeleteWedding(id)
}

func (s *WeddingService) CreatePeople(weddingId uuid.UUID, dto dto.PeopleDto) (string, error) {
	newPeople := NewPeopleFromDto(dto)
	newPeople.WeddingID = weddingId
	err := s.peopleRepo.CreatePeople(newPeople)
	if err != nil {
		return "", err
	}
	return "People created: " + newPeople.ID.String(), nil
}

func (s *WeddingService) CreatePhoto(weddingId uuid.UUID, photoDto dto.PhotoDto) (string, error) {
	newPhoto := NewPhotoFromDto(photoDto)
	newPhoto.WeddingID = weddingId
	err := s.photoRepo.CreatePhoto(newPhoto)
	if err != nil {
		return "", err
	}
	return "Photo created: " + newPhoto.ID.String(), nil
}

func (s *WeddingService) GetPhotosByWeddingId(weddingId uuid.UUID) ([]dto.PhotoDto, error) {
	var photos []models.Photo
	var photoDtos []dto.PhotoDto
	photos, err := s.photoRepo.GetPhotosByWeddingID(weddingId)
	if err != nil {
		return photoDtos, err
	}
	photoDtos = NewPhotoDtoListFromModels(photos)
	return photoDtos, nil
}

func NewWeddingDtoFromModel(wedding models.Wedding) *dto.WeddingDto {
	return &dto.WeddingDto{
		ID:            wedding.ID,
		StartDatetime: wedding.StartDateTime,
		Location:      wedding.Location,
		Groom:         wedding.Groom,
		Bride:         wedding.Bride,
		Peoples:       NewPeopleDtosFromModels(wedding.Peoples),
		Photos:        NewPhotoDtoListFromModels(wedding.Photos),
	}
}

func NewWeddingFromDto(weddingDto dto.WeddingDto) *models.Wedding {
	return &models.Wedding{
		StartDateTime: weddingDto.StartDatetime,
		Location:      weddingDto.Location,
	}
}

func NewPeopleFromDto(peopleDto dto.PeopleDto) *models.People {
	return &models.People{
		Name:  peopleDto.Name,
		Phone: peopleDto.Phone,
	}
}
func NewPeopleDtoFromModel(people models.People) *dto.PeopleDto {
	return &dto.PeopleDto{
		ID:    people.ID,
		Name:  people.Name,
		Phone: people.Phone,
	}
}

func NewPeopleDtosFromModels(peoples []models.People) []dto.PeopleDto {
	result := make([]dto.PeopleDto, len(peoples))
	for i, p := range peoples {
		result[i] = *NewPeopleDtoFromModel(p)
	}
	return result
}

func NewPhotoFromDto(photoDto dto.PhotoDto) *models.Photo {
	return &models.Photo{
		Url:    photoDto.Url,
		Status: "Not Started",
	}
}
func NewPhotoDtoFromModel(photo models.Photo) *dto.PhotoDto {
	return &dto.PhotoDto{
		ID:     photo.ID,
		Url:    photo.Url,
		Status: photo.Status,
	}
}
func NewPhotoDtoListFromModels(photos []models.Photo) []dto.PhotoDto {
	result := make([]dto.PhotoDto, len(photos))
	for i, p := range photos {
		result[i] = *NewPhotoDtoFromModel(p)
	}
	return result
}

func (s *WeddingService) updateWeddingFields(wedding *models.Wedding, updatedWeddingData dto.WeddingDto) {
	wedding.StartDateTime = updatedWeddingData.StartDatetime
	wedding.Location = updatedWeddingData.Location
	wedding.Groom = updatedWeddingData.Groom
	wedding.Bride = updatedWeddingData.Bride
}
