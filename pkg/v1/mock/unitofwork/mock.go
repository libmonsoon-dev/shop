// Code generated by MockGen. DO NOT EDIT.
// Source: ./contract.go

// Package mock_unitofwork is a generated GoMock package.
package mock_unitofwork

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	repository "github.com/libmonsoon-dev/shop/pkg/v1/repository"
	unitofwork "github.com/libmonsoon-dev/shop/pkg/v1/unitofwork"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Complete mocks base method.
func (m *MockInterface) Complete(err *error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Complete", err)
}

// Complete indicates an expected call of Complete.
func (mr *MockInterfaceMockRecorder) Complete(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Complete", reflect.TypeOf((*MockInterface)(nil).Complete), err)
}

// MockItemCreatorFactory is a mock of ItemCreatorFactory interface.
type MockItemCreatorFactory struct {
	ctrl     *gomock.Controller
	recorder *MockItemCreatorFactoryMockRecorder
}

// MockItemCreatorFactoryMockRecorder is the mock recorder for MockItemCreatorFactory.
type MockItemCreatorFactoryMockRecorder struct {
	mock *MockItemCreatorFactory
}

// NewMockItemCreatorFactory creates a new mock instance.
func NewMockItemCreatorFactory(ctrl *gomock.Controller) *MockItemCreatorFactory {
	mock := &MockItemCreatorFactory{ctrl: ctrl}
	mock.recorder = &MockItemCreatorFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemCreatorFactory) EXPECT() *MockItemCreatorFactoryMockRecorder {
	return m.recorder
}

// CreateContext mocks base method.
func (m *MockItemCreatorFactory) CreateContext(ctx context.Context) (unitofwork.ItemCreator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContext", ctx)
	ret0, _ := ret[0].(unitofwork.ItemCreator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContext indicates an expected call of CreateContext.
func (mr *MockItemCreatorFactoryMockRecorder) CreateContext(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContext", reflect.TypeOf((*MockItemCreatorFactory)(nil).CreateContext), ctx)
}

// MockItemCreator is a mock of ItemCreator interface.
type MockItemCreator struct {
	ctrl     *gomock.Controller
	recorder *MockItemCreatorMockRecorder
}

// MockItemCreatorMockRecorder is the mock recorder for MockItemCreator.
type MockItemCreatorMockRecorder struct {
	mock *MockItemCreator
}

// NewMockItemCreator creates a new mock instance.
func NewMockItemCreator(ctrl *gomock.Controller) *MockItemCreator {
	mock := &MockItemCreator{ctrl: ctrl}
	mock.recorder = &MockItemCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemCreator) EXPECT() *MockItemCreatorMockRecorder {
	return m.recorder
}

// AttributeRepository mocks base method.
func (m *MockItemCreator) AttributeRepository() repository.AttributeRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttributeRepository")
	ret0, _ := ret[0].(repository.AttributeRepository)
	return ret0
}

// AttributeRepository indicates an expected call of AttributeRepository.
func (mr *MockItemCreatorMockRecorder) AttributeRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttributeRepository", reflect.TypeOf((*MockItemCreator)(nil).AttributeRepository))
}

// Complete mocks base method.
func (m *MockItemCreator) Complete(err *error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Complete", err)
}

// Complete indicates an expected call of Complete.
func (mr *MockItemCreatorMockRecorder) Complete(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Complete", reflect.TypeOf((*MockItemCreator)(nil).Complete), err)
}

// ItemRepository mocks base method.
func (m *MockItemCreator) ItemRepository() repository.ItemRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ItemRepository")
	ret0, _ := ret[0].(repository.ItemRepository)
	return ret0
}

// ItemRepository indicates an expected call of ItemRepository.
func (mr *MockItemCreatorMockRecorder) ItemRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ItemRepository", reflect.TypeOf((*MockItemCreator)(nil).ItemRepository))
}

// ValueTypeRepository mocks base method.
func (m *MockItemCreator) ValueTypeRepository() repository.ValueTypeRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValueTypeRepository")
	ret0, _ := ret[0].(repository.ValueTypeRepository)
	return ret0
}

// ValueTypeRepository indicates an expected call of ValueTypeRepository.
func (mr *MockItemCreatorMockRecorder) ValueTypeRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValueTypeRepository", reflect.TypeOf((*MockItemCreator)(nil).ValueTypeRepository))
}