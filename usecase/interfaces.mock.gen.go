// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/interfaces.go

// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	reflect "reflect"

	model "github.com/ProductsAPI/model"
	gomock "github.com/golang/mock/gomock"
)

// MockUsecaseInterface is a mock of UsecaseInterface interface.
type MockUsecaseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseInterfaceMockRecorder
}

// MockUsecaseInterfaceMockRecorder is the mock recorder for MockUsecaseInterface.
type MockUsecaseInterfaceMockRecorder struct {
	mock *MockUsecaseInterface
}

// NewMockUsecaseInterface creates a new mock instance.
func NewMockUsecaseInterface(ctrl *gomock.Controller) *MockUsecaseInterface {
	mock := &MockUsecaseInterface{ctrl: ctrl}
	mock.recorder = &MockUsecaseInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecaseInterface) EXPECT() *MockUsecaseInterfaceMockRecorder {
	return m.recorder
}

// AddProduct mocks base method.
func (m *MockUsecaseInterface) AddProduct(ctx context.Context, product model.Product) (int, model.ValidationProductResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProduct", ctx, product)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(model.ValidationProductResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddProduct indicates an expected call of AddProduct.
func (mr *MockUsecaseInterfaceMockRecorder) AddProduct(ctx, product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProduct", reflect.TypeOf((*MockUsecaseInterface)(nil).AddProduct), ctx, product)
}

// GetProduct mocks base method.
func (m *MockUsecaseInterface) GetProduct(ctx context.Context, productID int) (model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", ctx, productID)
	ret0, _ := ret[0].(model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockUsecaseInterfaceMockRecorder) GetProduct(ctx, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockUsecaseInterface)(nil).GetProduct), ctx, productID)
}
