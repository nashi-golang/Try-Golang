package services

type WeddingService struct{}

func NewWeddingService() *WeddingService {
	return &WeddingService{}
}

func (s *WeddingService) CreateWedding() string {
	return "Wedding created"
}

func (s *WeddingService) GetWeddings() string {
	return "Getting all weddings"
}

func (s *WeddingService) GetWeddingById(id string) string {
	return "Getting wedding with ID: " + id
}

func (s *WeddingService) UpdateWedding(id string) string {
	return "Updating wedding with ID: " + id
}

func (s *WeddingService) DeleteWedding(id string) string {
	return "Deleting wedding with ID: " + id
}
