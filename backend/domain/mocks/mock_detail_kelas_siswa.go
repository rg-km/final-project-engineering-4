// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rg-km/final-project-engineering-4/backend/domain (interfaces: DetailKelasSiswaRepository,DetailKelasSiswaUseCase)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
)

// MockDetailKelasSiswaRepository is a mock of DetailKelasSiswaRepository interface.
type MockDetailKelasSiswaRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDetailKelasSiswaRepositoryMockRecorder
}

// MockDetailKelasSiswaRepositoryMockRecorder is the mock recorder for MockDetailKelasSiswaRepository.
type MockDetailKelasSiswaRepositoryMockRecorder struct {
	mock *MockDetailKelasSiswaRepository
}

// NewMockDetailKelasSiswaRepository creates a new mock instance.
func NewMockDetailKelasSiswaRepository(ctrl *gomock.Controller) *MockDetailKelasSiswaRepository {
	mock := &MockDetailKelasSiswaRepository{ctrl: ctrl}
	mock.recorder = &MockDetailKelasSiswaRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDetailKelasSiswaRepository) EXPECT() *MockDetailKelasSiswaRepositoryMockRecorder {
	return m.recorder
}

// MockDetailKelasSiswaUseCase is a mock of DetailKelasSiswaUseCase interface.
type MockDetailKelasSiswaUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockDetailKelasSiswaUseCaseMockRecorder
}

// MockDetailKelasSiswaUseCaseMockRecorder is the mock recorder for MockDetailKelasSiswaUseCase.
type MockDetailKelasSiswaUseCaseMockRecorder struct {
	mock *MockDetailKelasSiswaUseCase
}

// NewMockDetailKelasSiswaUseCase creates a new mock instance.
func NewMockDetailKelasSiswaUseCase(ctrl *gomock.Controller) *MockDetailKelasSiswaUseCase {
	mock := &MockDetailKelasSiswaUseCase{ctrl: ctrl}
	mock.recorder = &MockDetailKelasSiswaUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDetailKelasSiswaUseCase) EXPECT() *MockDetailKelasSiswaUseCaseMockRecorder {
	return m.recorder
}
