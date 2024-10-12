package services

import (
	"time"
	"try-golang/pkg/dto"
)

type WeddingService struct{}

func NewWeddingService() *WeddingService {
	return &WeddingService{}
}

func (s *WeddingService) CreateWedding(dto dto.WeddingDto) string {

	return "Wedding created" + dto.Location
}

func (s *WeddingService) GetWeddings() []dto.WeddingDto {
	var weddings []dto.WeddingDto
	weddings = append(weddings, dto.WeddingDto{
		ID:            "id",
		StartDatetime: time.Now(),
		Location:      "신림역",
		Groom:         "신랑",
		Bride:         "신부",
	})

	return weddings
}

func (s *WeddingService) GetWeddingById(id string) dto.WeddingDto {
	var wedding = dto.WeddingDto{
		ID:            "id",
		StartDatetime: time.Now(),
		Location:      "신림역",
		Groom:         "신랑",
		Bride:         "신부",
	}
	return wedding
}

func (s *WeddingService) UpdateWedding(id string) string {
	return "Updating wedding with ID: " + id
}

func (s *WeddingService) DeleteWedding(id string) string {
	return "Deleting wedding with ID: " + id
}
