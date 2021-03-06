// Code generated by MockGen. DO NOT EDIT.
// Source: ./contract.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	shopv1 "github.com/libmonsoon-dev/shop/pkg/v1"
)

// MockItemRepository is a mock of ItemRepository interface.
type MockItemRepository struct {
	ctrl     *gomock.Controller
	recorder *MockItemRepositoryMockRecorder
}

// MockItemRepositoryMockRecorder is the mock recorder for MockItemRepository.
type MockItemRepositoryMockRecorder struct {
	mock *MockItemRepository
}

// NewMockItemRepository creates a new mock instance.
func NewMockItemRepository(ctrl *gomock.Controller) *MockItemRepository {
	mock := &MockItemRepository{ctrl: ctrl}
	mock.recorder = &MockItemRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemRepository) EXPECT() *MockItemRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockItemRepository) Create(arg0 context.Context, arg1 *shopv1.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockItemRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockItemRepository)(nil).Create), arg0, arg1)
}

// MockAttributeRepository is a mock of AttributeRepository interface.
type MockAttributeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAttributeRepositoryMockRecorder
}

// MockAttributeRepositoryMockRecorder is the mock recorder for MockAttributeRepository.
type MockAttributeRepositoryMockRecorder struct {
	mock *MockAttributeRepository
}

// NewMockAttributeRepository creates a new mock instance.
func NewMockAttributeRepository(ctrl *gomock.Controller) *MockAttributeRepository {
	mock := &MockAttributeRepository{ctrl: ctrl}
	mock.recorder = &MockAttributeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAttributeRepository) EXPECT() *MockAttributeRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAttributeRepository) Create(arg0 context.Context, arg1 *shopv1.Attribute) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAttributeRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAttributeRepository)(nil).Create), arg0, arg1)
}

// MockValueTypeRepository is a mock of ValueTypeRepository interface.
type MockValueTypeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockValueTypeRepositoryMockRecorder
}

// MockValueTypeRepositoryMockRecorder is the mock recorder for MockValueTypeRepository.
type MockValueTypeRepositoryMockRecorder struct {
	mock *MockValueTypeRepository
}

// NewMockValueTypeRepository creates a new mock instance.
func NewMockValueTypeRepository(ctrl *gomock.Controller) *MockValueTypeRepository {
	mock := &MockValueTypeRepository{ctrl: ctrl}
	mock.recorder = &MockValueTypeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValueTypeRepository) EXPECT() *MockValueTypeRepositoryMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockValueTypeRepository) Save(arg0 context.Context, arg1 *shopv1.ValueType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockValueTypeRepositoryMockRecorder) Save(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockValueTypeRepository)(nil).Save), arg0, arg1)
}
