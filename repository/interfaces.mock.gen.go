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

// DeleteProduct mocks base method.
func (m *MockRepositoryInterface) DeleteProduct(ctx context.Context, productID int, tx sqlwrapper.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", ctx, productID, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockRepositoryInterfaceMockRecorder) DeleteProduct(ctx, productID, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockRepositoryInterface)(nil).DeleteProduct), ctx, productID, tx)
}

// DeleteProductVariety mocks base method.
func (m *MockRepositoryInterface) DeleteProductVariety(ctx context.Context, variety model.ProductVariety, tx sqlwrapper.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProductVariety", ctx, variety, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProductVariety indicates an expected call of DeleteProductVariety.
func (mr *MockRepositoryInterfaceMockRecorder) DeleteProductVariety(ctx, variety, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProductVariety", reflect.TypeOf((*MockRepositoryInterface)(nil).DeleteProductVariety), ctx, variety, tx)
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

// InsertProductVariety mocks base method.
func (m *MockRepositoryInterface) InsertProductVariety(ctx context.Context, variety model.ProductVariety, tx sqlwrapper.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertProductVariety", ctx, variety, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertProductVariety indicates an expected call of InsertProductVariety.
func (mr *MockRepositoryInterfaceMockRecorder) InsertProductVariety(ctx, variety, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertProductVariety", reflect.TypeOf((*MockRepositoryInterface)(nil).InsertProductVariety), ctx, variety, tx)
}

// UpdateProduct mocks base method.
func (m *MockRepositoryInterface) UpdateProduct(ctx context.Context, product model.Product, tx sqlwrapper.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", ctx, product, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockRepositoryInterfaceMockRecorder) UpdateProduct(ctx, product, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockRepositoryInterface)(nil).UpdateProduct), ctx, product, tx)
}

// UpdateProductVariety mocks base method.
func (m *MockRepositoryInterface) UpdateProductVariety(ctx context.Context, variety model.ProductVariety, tx sqlwrapper.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductVariety", ctx, variety, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProductVariety indicates an expected call of UpdateProductVariety.
func (mr *MockRepositoryInterfaceMockRecorder) UpdateProductVariety(ctx, variety, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductVariety", reflect.TypeOf((*MockRepositoryInterface)(nil).UpdateProductVariety), ctx, variety, tx)
}
