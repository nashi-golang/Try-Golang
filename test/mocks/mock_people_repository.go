// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/people_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	models "try-golang/internal/models"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockPeopleRepository is a mock of PeopleRepository interface.
type MockPeopleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPeopleRepositoryMockRecorder
}

// MockPeopleRepositoryMockRecorder is the mock recorder for MockPeopleRepository.
type MockPeopleRepositoryMockRecorder struct {
	mock *MockPeopleRepository
}

// NewMockPeopleRepository creates a new mock instance.
func NewMockPeopleRepository(ctrl *gomock.Controller) *MockPeopleRepository {
	mock := &MockPeopleRepository{ctrl: ctrl}
	mock.recorder = &MockPeopleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeopleRepository) EXPECT() *MockPeopleRepositoryMockRecorder {
	return m.recorder
}

// CreatePeople mocks base method.
func (m *MockPeopleRepository) CreatePeople(People *models.People) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePeople", People)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePeople indicates an expected call of CreatePeople.
func (mr *MockPeopleRepositoryMockRecorder) CreatePeople(People interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePeople", reflect.TypeOf((*MockPeopleRepository)(nil).CreatePeople), People)
}

// DeletePeople mocks base method.
func (m *MockPeopleRepository) DeletePeople(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePeople", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePeople indicates an expected call of DeletePeople.
func (mr *MockPeopleRepositoryMockRecorder) DeletePeople(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePeople", reflect.TypeOf((*MockPeopleRepository)(nil).DeletePeople), id)
}

// GetPeopleByID mocks base method.
func (m *MockPeopleRepository) GetPeopleByID(id uuid.UUID) (*models.People, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeopleByID", id)
	ret0, _ := ret[0].(*models.People)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeopleByID indicates an expected call of GetPeopleByID.
func (mr *MockPeopleRepositoryMockRecorder) GetPeopleByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeopleByID", reflect.TypeOf((*MockPeopleRepository)(nil).GetPeopleByID), id)
}

// UpdatePeople mocks base method.
func (m *MockPeopleRepository) UpdatePeople(People *models.People) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePeople", People)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePeople indicates an expected call of UpdatePeople.
func (mr *MockPeopleRepositoryMockRecorder) UpdatePeople(People interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePeople", reflect.TypeOf((*MockPeopleRepository)(nil).UpdatePeople), People)
}
