package services

import (
	"errors"
	"github.com/google/uuid"
	"try-golang/internal/models"
	"try-golang/internal/repository"
	"try-golang/pkg/dto"
)

type WeddingService struct {
	repo *repository.WeddingRepository
}

func NewWeddingService(repo *repository.WeddingRepository) *WeddingService {
	return &WeddingService{
		repo: repo,
	}
}

func (s *WeddingService) CreateWedding(dto dto.WeddingDto) (string, error) {
	newWedding := NewWeddingFromDto(dto)
	err := s.repo.CreateWedding(newWedding)
	if err != nil {
		return "", err
	}
	return "Wedding created: " + newWedding.ID.String(), nil
}

func (s *WeddingService) GetWeddings() []dto.WeddingDto {

	var weddingDtos []dto.WeddingDto

	weddings, err := s.repo.GetAllWeddings()
	if err != nil {
		return nil
	}

	for i := range weddings {
		weddingDto := NewWeddingDtoFromModel(weddings[i])
		weddingDtos = append(weddingDtos, *weddingDto)
	}
	return weddingDtos
}

func (s *WeddingService) GetWeddingById(id uuid.UUID) (dto.WeddingDto, error) {

	wedding, err := s.repo.GetWeddingByID(id)
	if err != nil {
		return dto.WeddingDto{}, err
	}
	weddingDto := NewWeddingDtoFromModel(*wedding)

	return *weddingDto, err
}

func (s *WeddingService) UpdateWedding(id uuid.UUID, updatedWeddingData dto.WeddingDto) (string, error) {
	wedding, err := s.repo.GetWeddingByID(id)
	if err != nil {
		return "", errors.New("wedding not found")
	}

	s.updateWeddingFields(wedding, updatedWeddingData)

	err = s.repo.UpdateWedding(wedding)
	if err != nil {
		return "", errors.New("failed to update wedding")
	}

	return "Wedding updated successfully", nil
}

func (s *WeddingService) DeleteWedding(id uuid.UUID) error {
	return s.repo.DeleteWedding(id)
}

func NewWeddingDtoFromModel(wedding models.Wedding) *dto.WeddingDto {
	return &dto.WeddingDto{
		ID:            wedding.ID,
		StartDatetime: wedding.StartDateTime,
		Location:      wedding.Location,
		Groom:         wedding.Groom,
		Bride:         wedding.Bride,
	}
}
func NewWeddingFromDto(weddingDto dto.WeddingDto) *models.Wedding {
	return &models.Wedding{
		StartDateTime: weddingDto.StartDatetime,
		Location:      weddingDto.Location,
	}
}
func (s *WeddingService) updateWeddingFields(wedding *models.Wedding, updatedWeddingData dto.WeddingDto) {
	wedding.StartDateTime = updatedWeddingData.StartDatetime
	wedding.Location = updatedWeddingData.Location
	wedding.Groom = updatedWeddingData.Groom
	wedding.Bride = updatedWeddingData.Bride
}
