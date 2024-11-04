package services

import (
	"github.com/google/uuid"
	"math/rand"
	"sync"
	"time"
	"try-golang/internal/models"
	"try-golang/internal/repository"
)

type ImageService struct {
	photoRepo repository.PhotoRepository
}
type ImageTask struct {
	Photo  models.Photo
	Status string
}

type StatusUpdate struct {
	Photo  models.Photo
	Status string
}

func NewImageService(photoRepo repository.PhotoRepository) *ImageService {
	return &ImageService{
		photoRepo: photoRepo,
	}
}

func (s *ImageService) StartPhotoProcessAllImage(weddingId uuid.UUID) error {
	photos, err := s.photoRepo.GetPhotosByWeddingID(weddingId)
	if err != nil {
		return err
	}

	go s.processPhotosInBackground(photos)

	return nil
}
func (s *ImageService) processPhotosInBackground(photos []models.Photo) {
	taskQueue := make(chan ImageTask, len(photos))
	statusUpdates := make(chan StatusUpdate, len(photos))
	done := make(chan bool)

	imageStatus := make(map[uuid.UUID]models.Photo)

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.worker(taskQueue, statusUpdates)
		}()
	}

	go s.dbUpdater(statusUpdates, imageStatus, done)

	for _, photo := range photos {
		taskQueue <- ImageTask{Photo: photo, Status: "Pending"}
	}

	close(taskQueue)

	wg.Wait()

	close(statusUpdates)

	<-done
}

func (s *ImageService) worker(taskQueue <-chan ImageTask, statusUpdates chan<- StatusUpdate) {
	for task := range taskQueue {

		statusUpdates <- StatusUpdate{Photo: task.Photo, Status: "In Progress"}

		randomDuration := time.Duration(rand.Int63n(int64(10 * time.Second)))
		time.Sleep(randomDuration)

		if time.Now().UnixNano()%9 != 0 {
			statusUpdates <- StatusUpdate{Photo: task.Photo, Status: "Success"}
		} else {
			statusUpdates <- StatusUpdate{Photo: task.Photo, Status: "Failure"}
		}
	}
}

func (s *ImageService) dbUpdater(statusUpdates <-chan StatusUpdate, imageStatus map[uuid.UUID]models.Photo, done chan<- bool) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case update, ok := <-statusUpdates:
			if !ok {
				s.saveToDB(imageStatus)
				done <- true
				return
			}
			update.Photo.Status = update.Status
			imageStatus[update.Photo.ID] = update.Photo
		case <-ticker.C:
			s.saveToDB(imageStatus)
		}
	}
}

func (s *ImageService) saveToDB(imageStatus map[uuid.UUID]models.Photo) {
	var photos []models.Photo

	for id, photo := range imageStatus {
		photos = append(photos, photo)
		delete(imageStatus, id)
	}

	_ = s.photoRepo.UpdatePhotos(photos)
}
