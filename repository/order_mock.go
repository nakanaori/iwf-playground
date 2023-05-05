// Code generated by MockGen. DO NOT EDIT.
// Source: repository/order.go

// Package repository is a generated GoMock package.
package repository

import (
	model "iwf-playground/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOrderRepository is a mock of OrderRepository interface.
type MockOrderRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepositoryMockRecorder
}

// MockOrderRepositoryMockRecorder is the mock recorder for MockOrderRepository.
type MockOrderRepositoryMockRecorder struct {
	mock *MockOrderRepository
}

// NewMockOrderRepository creates a new mock instance.
func NewMockOrderRepository(ctrl *gomock.Controller) *MockOrderRepository {
	mock := &MockOrderRepository{ctrl: ctrl}
	mock.recorder = &MockOrderRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepository) EXPECT() *MockOrderRepositoryMockRecorder {
	return m.recorder
}

// GetOrderById mocks base method.
func (m *MockOrderRepository) GetOrderById(orderId string) (model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderById", orderId)
	ret0, _ := ret[0].(model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderById indicates an expected call of GetOrderById.
func (mr *MockOrderRepositoryMockRecorder) GetOrderById(orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderById", reflect.TypeOf((*MockOrderRepository)(nil).GetOrderById), orderId)
}

// SetOrderStatus mocks base method.
func (m *MockOrderRepository) SetOrderStatus(orderId, state string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetOrderStatus", orderId, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetOrderStatus indicates an expected call of SetOrderStatus.
func (mr *MockOrderRepositoryMockRecorder) SetOrderStatus(orderId, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOrderStatus", reflect.TypeOf((*MockOrderRepository)(nil).SetOrderStatus), orderId, state)
}