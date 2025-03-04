// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/omkarbhostekar/brewgo/order/db/sqlc (interfaces: Store)

// Package mock_sqlc is a generated GoMock package.
package mock_sqlc

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	orders "github.com/omkarbhostekar/brewgo/order/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AreAllOrderItemsReady mocks base method.
func (m *MockStore) AreAllOrderItemsReady(arg0 context.Context, arg1 int32) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AreAllOrderItemsReady", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AreAllOrderItemsReady indicates an expected call of AreAllOrderItemsReady.
func (mr *MockStoreMockRecorder) AreAllOrderItemsReady(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AreAllOrderItemsReady", reflect.TypeOf((*MockStore)(nil).AreAllOrderItemsReady), arg0, arg1)
}

// CreateOrder mocks base method.
func (m *MockStore) CreateOrder(arg0 context.Context, arg1 orders.CreateOrderParams) (orders.CounterOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", arg0, arg1)
	ret0, _ := ret[0].(orders.CounterOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockStoreMockRecorder) CreateOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockStore)(nil).CreateOrder), arg0, arg1)
}

// CreateOrderItem mocks base method.
func (m *MockStore) CreateOrderItem(arg0 context.Context, arg1 orders.CreateOrderItemParams) (orders.CounterOrderItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrderItem", arg0, arg1)
	ret0, _ := ret[0].(orders.CounterOrderItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrderItem indicates an expected call of CreateOrderItem.
func (mr *MockStoreMockRecorder) CreateOrderItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderItem", reflect.TypeOf((*MockStore)(nil).CreateOrderItem), arg0, arg1)
}

// DeleteOrder mocks base method.
func (m *MockStore) DeleteOrder(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrder", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrder indicates an expected call of DeleteOrder.
func (mr *MockStoreMockRecorder) DeleteOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockStore)(nil).DeleteOrder), arg0, arg1)
}

// DeleteOrderItem mocks base method.
func (m *MockStore) DeleteOrderItem(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrderItem", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrderItem indicates an expected call of DeleteOrderItem.
func (mr *MockStoreMockRecorder) DeleteOrderItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrderItem", reflect.TypeOf((*MockStore)(nil).DeleteOrderItem), arg0, arg1)
}

// GetOrderById mocks base method.
func (m *MockStore) GetOrderById(arg0 context.Context, arg1 int32) (orders.CounterOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderById", arg0, arg1)
	ret0, _ := ret[0].(orders.CounterOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderById indicates an expected call of GetOrderById.
func (mr *MockStoreMockRecorder) GetOrderById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderById", reflect.TypeOf((*MockStore)(nil).GetOrderById), arg0, arg1)
}

// GetOrderItemsByOrderId mocks base method.
func (m *MockStore) GetOrderItemsByOrderId(arg0 context.Context, arg1 int32) ([]orders.GetOrderItemsByOrderIdRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderItemsByOrderId", arg0, arg1)
	ret0, _ := ret[0].([]orders.GetOrderItemsByOrderIdRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderItemsByOrderId indicates an expected call of GetOrderItemsByOrderId.
func (mr *MockStoreMockRecorder) GetOrderItemsByOrderId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderItemsByOrderId", reflect.TypeOf((*MockStore)(nil).GetOrderItemsByOrderId), arg0, arg1)
}

// GetOrdersByUserId mocks base method.
func (m *MockStore) GetOrdersByUserId(arg0 context.Context, arg1 int32) ([]orders.CounterOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrdersByUserId", arg0, arg1)
	ret0, _ := ret[0].([]orders.CounterOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrdersByUserId indicates an expected call of GetOrdersByUserId.
func (mr *MockStoreMockRecorder) GetOrdersByUserId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrdersByUserId", reflect.TypeOf((*MockStore)(nil).GetOrdersByUserId), arg0, arg1)
}

// GetTotalAmountByOrderId mocks base method.
func (m *MockStore) GetTotalAmountByOrderId(arg0 context.Context, arg1 int32) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalAmountByOrderId", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalAmountByOrderId indicates an expected call of GetTotalAmountByOrderId.
func (mr *MockStoreMockRecorder) GetTotalAmountByOrderId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalAmountByOrderId", reflect.TypeOf((*MockStore)(nil).GetTotalAmountByOrderId), arg0, arg1)
}

// PlaceOrderTx mocks base method.
func (m *MockStore) PlaceOrderTx(arg0 context.Context, arg1 orders.PlaceOrderTxParams) (orders.PlaceOrderTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlaceOrderTx", arg0, arg1)
	ret0, _ := ret[0].(orders.PlaceOrderTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PlaceOrderTx indicates an expected call of PlaceOrderTx.
func (mr *MockStoreMockRecorder) PlaceOrderTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlaceOrderTx", reflect.TypeOf((*MockStore)(nil).PlaceOrderTx), arg0, arg1)
}

// UpdateOrder mocks base method.
func (m *MockStore) UpdateOrder(arg0 context.Context, arg1 orders.UpdateOrderParams) (orders.CounterOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", arg0, arg1)
	ret0, _ := ret[0].(orders.CounterOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockStoreMockRecorder) UpdateOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockStore)(nil).UpdateOrder), arg0, arg1)
}

// UpdateOrderItemStatus mocks base method.
func (m *MockStore) UpdateOrderItemStatus(arg0 context.Context, arg1 orders.UpdateOrderItemStatusParams) (orders.CounterOrderItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderItemStatus", arg0, arg1)
	ret0, _ := ret[0].(orders.CounterOrderItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrderItemStatus indicates an expected call of UpdateOrderItemStatus.
func (mr *MockStoreMockRecorder) UpdateOrderItemStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderItemStatus", reflect.TypeOf((*MockStore)(nil).UpdateOrderItemStatus), arg0, arg1)
}
