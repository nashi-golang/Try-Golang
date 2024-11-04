package services

import (
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"testing"
	"try-golang/internal/models"
	"try-golang/internal/services"
	"try-golang/pkg/dto"
	"try-golang/test/mocks"
)

func TestWeddingService_CreateWedding(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWeddingRepo := mocks.NewMockWeddingRepository(ctrl)

	weddingService := services.NewWeddingService(mockWeddingRepo, nil, nil)
	testWeddingDto := dto.WeddingDto{
		Location: "Sample Location",
	}

	mockWeddingRepo.EXPECT().CreateWedding(gomock.Any()).Return(nil).Times(1)

	result, err := weddingService.CreateWedding(testWeddingDto)
	if err != nil || result == "" {
		t.Fatalf("expected successful wedding creation, got error: %v", err)
	}
}

func TestWeddingService_GetWeddings(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWeddingRepo := mocks.NewMockWeddingRepository(ctrl)
	weddingService := services.NewWeddingService(mockWeddingRepo, nil, nil)

	testID := uuid.New()

	weddings := []models.Wedding{
		{ID: testID, Location: "Sample Location"},
	}
	mockWeddingRepo.EXPECT().GetAllWeddings().Return(weddings, nil).Times(1)

	result := weddingService.GetWeddings()
	if len(result) != 1 {
		t.Fatalf("expected: 1 wedding, but got %v", len(result))
	}
	if result[0].ID != testID {
		t.Fatalf("expected id: %s , but got %v", testID, result[0].ID)
	}
}

func TestWeddingService_GetWeddingById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWeddingRepo := mocks.NewMockWeddingRepository(ctrl)
	weddingService := services.NewWeddingService(mockWeddingRepo, nil, nil)

	testID := uuid.New()
	testWedding := &models.Wedding{
		ID:       testID,
		Location: "Sample Location",
	}
	mockWeddingRepo.EXPECT().GetWeddingByID(testID).Return(testWedding, nil).Times(1)

	result, _ := weddingService.GetWeddingById(testID)
	if result.ID != testID {
		t.Fatalf("expected ID: %v, but got %v", testID, result.ID)
	}
}

func TestWeddingService_UpdateWedding(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWeddingRepo := mocks.NewMockWeddingRepository(ctrl)
	weddingService := services.NewWeddingService(mockWeddingRepo, nil, nil)

	testID := uuid.New()
	testWedding := &models.Wedding{
		ID:       testID,
		Location: "Sample Location",
	}
	updatedWeddingData := dto.WeddingDto{
		Location: "New Location",
	}

	mockWeddingRepo.EXPECT().GetWeddingByID(testID).Return(testWedding, nil).Times(1)
	mockWeddingRepo.EXPECT().UpdateWedding(gomock.Any()).Return(nil).Times(1)

	_, err := weddingService.UpdateWedding(testID, updatedWeddingData)
	if err != nil {
		t.Fatalf("expected successful wedding update, got error: %v", err)
	}

}

func TestWeddingService_DeleteWedding(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWeddingRepo := mocks.NewMockWeddingRepository(ctrl)
	weddingService := services.NewWeddingService(mockWeddingRepo, nil, nil)

	testID := uuid.New()
	mockWeddingRepo.EXPECT().DeleteWedding(testID).Return(nil)

	err := weddingService.DeleteWedding(testID)
	if err != nil {
		t.Fatalf("expected successful wedding deletion, got error: %v", err)
	}
}
