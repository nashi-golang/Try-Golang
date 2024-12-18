// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/photo_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	models "try-golang/internal/models"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockPhotoRepository is a mock of PhotoRepository interface.
type MockPhotoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPhotoRepositoryMockRecorder
}

// MockPhotoRepositoryMockRecorder is the mock recorder for MockPhotoRepository.
type MockPhotoRepositoryMockRecorder struct {
	mock *MockPhotoRepository
}

// NewMockPhotoRepository creates a new mock instance.
func NewMockPhotoRepository(ctrl *gomock.Controller) *MockPhotoRepository {
	mock := &MockPhotoRepository{ctrl: ctrl}
	mock.recorder = &MockPhotoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhotoRepository) EXPECT() *MockPhotoRepositoryMockRecorder {
	return m.recorder
}

// CreatePhoto mocks base method.
func (m *MockPhotoRepository) CreatePhoto(Photo *models.Photo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePhoto", Photo)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePhoto indicates an expected call of CreatePhoto.
func (mr *MockPhotoRepositoryMockRecorder) CreatePhoto(Photo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePhoto", reflect.TypeOf((*MockPhotoRepository)(nil).CreatePhoto), Photo)
}

// DeletePhoto mocks base method.
func (m *MockPhotoRepository) DeletePhoto(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePhoto", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePhoto indicates an expected call of DeletePhoto.
func (mr *MockPhotoRepositoryMockRecorder) DeletePhoto(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePhoto", reflect.TypeOf((*MockPhotoRepository)(nil).DeletePhoto), id)
}

// GetAllPhotos mocks base method.
func (m *MockPhotoRepository) GetAllPhotos() ([]models.Photo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPhotos")
	ret0, _ := ret[0].([]models.Photo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPhotos indicates an expected call of GetAllPhotos.
func (mr *MockPhotoRepositoryMockRecorder) GetAllPhotos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPhotos", reflect.TypeOf((*MockPhotoRepository)(nil).GetAllPhotos))
}

// GetPhotosByWeddingID mocks base method.
func (m *MockPhotoRepository) GetPhotosByWeddingID(weddingId uuid.UUID) ([]models.Photo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhotosByWeddingID", weddingId)
	ret0, _ := ret[0].([]models.Photo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotosByWeddingID indicates an expected call of GetPhotosByWeddingID.
func (mr *MockPhotoRepositoryMockRecorder) GetPhotosByWeddingID(weddingId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotosByWeddingID", reflect.TypeOf((*MockPhotoRepository)(nil).GetPhotosByWeddingID), weddingId)
}

// UpdatePhoto mocks base method.
func (m *MockPhotoRepository) UpdatePhoto(Photo *models.Photo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePhoto", Photo)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePhoto indicates an expected call of UpdatePhoto.
func (mr *MockPhotoRepositoryMockRecorder) UpdatePhoto(Photo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePhoto", reflect.TypeOf((*MockPhotoRepository)(nil).UpdatePhoto), Photo)
}

// UpdatePhotos mocks base method.
func (m *MockPhotoRepository) UpdatePhotos(Photo []models.Photo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePhotos", Photo)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePhotos indicates an expected call of UpdatePhotos.
func (mr *MockPhotoRepositoryMockRecorder) UpdatePhotos(Photo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePhotos", reflect.TypeOf((*MockPhotoRepository)(nil).UpdatePhotos), Photo)
}
