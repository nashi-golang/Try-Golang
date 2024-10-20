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
}

func NewWeddingService(weddingRepo *repository.WeddingRepository, peopleRepo *repository.PeopleRepository) *WeddingService {
	return &WeddingService{
		weddingRepo: weddingRepo,
		peopleRepo:  peopleRepo,
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

func NewWeddingDtoFromModel(wedding models.Wedding) *dto.WeddingDto {
	return &dto.WeddingDto{
		ID:            wedding.ID,
		StartDatetime: wedding.StartDateTime,
		Location:      wedding.Location,
		Groom:         wedding.Groom,
		Bride:         wedding.Bride,
		Peoples:       NewPeopleDtosFromModels(wedding.Peoples),
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
func (s *WeddingService) updateWeddingFields(wedding *models.Wedding, updatedWeddingData dto.WeddingDto) {
	wedding.StartDateTime = updatedWeddingData.StartDatetime
	wedding.Location = updatedWeddingData.Location
	wedding.Groom = updatedWeddingData.Groom
	wedding.Bride = updatedWeddingData.Bride
}
