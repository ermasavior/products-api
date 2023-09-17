// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interfaces.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	model "github.com/ProductsAPI/model"
	sqlwrapper "github.com/ProductsAPI/utils/sqlwrapper"
	gomock "github.com/golang/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// GetDatabase mocks base method.
func (m *MockRepositoryInterface) GetDatabase() *sql.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDatabase")
	ret0, _ := ret[0].(*sql.DB)
	return ret0
}

// GetDatabase indicates an expected call of GetDatabase.
func (mr *MockRepositoryInterfaceMockRecorder) GetDatabase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatabase", reflect.TypeOf((*MockRepositoryInterface)(nil).GetDatabase))
}

// GetProductByProductID mocks base method.
func (m *MockRepositoryInterface) GetProductByProductID(ctx context.Context, productID int) (model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductByProductID", ctx, productID)
	ret0, _ := ret[0].(model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductByProductID indicates an expected call of GetProductByProductID.
func (mr *MockRepositoryInterfaceMockRecorder) GetProductByProductID(ctx, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductByProductID", reflect.TypeOf((*MockRepositoryInterface)(nil).GetProductByProductID), ctx, productID)
}

// GetProductVarietiesByProductID mocks base method.
func (m *MockRepositoryInterface) GetProductVarietiesByProductID(ctx context.Context, productID int) ([]model.ProductVariety, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductVarietiesByProductID", ctx, productID)
	ret0, _ := ret[0].([]model.ProductVariety)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductVarietiesByProductID indicates an expected call of GetProductVarietiesByProductID.
func (mr *MockRepositoryInterfaceMockRecorder) GetProductVarietiesByProductID(ctx, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductVarietiesByProductID", reflect.TypeOf((*MockRepositoryInterface)(nil).GetProductVarietiesByProductID), ctx, productID)
}

// InsertProduct mocks base method.
func (m *MockRepositoryInterface) InsertProduct(ctx context.Context, product model.Product, tx sqlwrapper.Transaction) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertProduct", ctx, product, tx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertProduct indicates an expected call of InsertProduct.
func (mr *MockRepositoryInterfaceMockRecorder) InsertProduct(ctx, product, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertProduct", reflect.TypeOf((*MockRepositoryInterface)(nil).InsertProduct), ctx, product, tx)
}

// InsertProductVarieties mocks base method.
func (m *MockRepositoryInterface) InsertProductVarieties(ctx context.Context, productID int, varieties []model.ProductVariety, tx sqlwrapper.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertProductVarieties", ctx, productID, varieties, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertProductVarieties indicates an expected call of InsertProductVarieties.
func (mr *MockRepositoryInterfaceMockRecorder) InsertProductVarieties(ctx, productID, varieties, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertProductVarieties", reflect.TypeOf((*MockRepositoryInterface)(nil).InsertProductVarieties), ctx, productID, varieties, tx)
}
