// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package repositories is a generated GoMock package.
package repositories

import (
	context "context"
	reflect "reflect"

	entities "github.com/RyaWcksn/fiber-restful/entities"
	gomock "github.com/golang/mock/gomock"
)

// MockCRepository is a mock of CRepository interface.
type MockCRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCRepositoryMockRecorder
}

// MockCRepositoryMockRecorder is the mock recorder for MockCRepository.
type MockCRepositoryMockRecorder struct {
	mock *MockCRepository
}

// NewMockCRepository creates a new mock instance.
func NewMockCRepository(ctrl *gomock.Controller) *MockCRepository {
	mock := &MockCRepository{ctrl: ctrl}
	mock.recorder = &MockCRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCRepository) EXPECT() *MockCRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockCRepository) Get(ctx context.Context, id int) (*entities.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*entities.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCRepositoryMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCRepository)(nil).Get), ctx, id)
}
