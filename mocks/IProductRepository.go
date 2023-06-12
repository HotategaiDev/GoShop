// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/repositories/product.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	models "goshop/app/models"
	serializers "goshop/app/serializers"
	paging "goshop/pkg/paging"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIProductRepository is a mock of IProductRepository interface.
type MockIProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIProductRepositoryMockRecorder
}

// MockIProductRepositoryMockRecorder is the mock recorder for MockIProductRepository.
type MockIProductRepositoryMockRecorder struct {
	mock *MockIProductRepository
}

// NewMockIProductRepository creates a new mock instance.
func NewMockIProductRepository(ctrl *gomock.Controller) *MockIProductRepository {
	mock := &MockIProductRepository{ctrl: ctrl}
	mock.recorder = &MockIProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProductRepository) EXPECT() *MockIProductRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIProductRepository) Create(ctx context.Context, product *models.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIProductRepositoryMockRecorder) Create(ctx, product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIProductRepository)(nil).Create), ctx, product)
}

// GetProductByID mocks base method.
func (m *MockIProductRepository) GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductByID", ctx, id)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductByID indicates an expected call of GetProductByID.
func (mr *MockIProductRepositoryMockRecorder) GetProductByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductByID", reflect.TypeOf((*MockIProductRepository)(nil).GetProductByID), ctx, id)
}

// ListProducts mocks base method.
func (m *MockIProductRepository) ListProducts(ctx context.Context, req serializers.ListProductReq) ([]*models.Product, *paging.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducts", ctx, req)
	ret0, _ := ret[0].([]*models.Product)
	ret1, _ := ret[1].(*paging.Pagination)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockIProductRepositoryMockRecorder) ListProducts(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockIProductRepository)(nil).ListProducts), ctx, req)
}

// Update mocks base method.
func (m *MockIProductRepository) Update(ctx context.Context, product *models.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIProductRepositoryMockRecorder) Update(ctx, product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIProductRepository)(nil).Update), ctx, product)
}
