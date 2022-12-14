// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "altaproject2/domain"

	mock "github.com/stretchr/testify/mock"
)

// OrderUseCase is an autogenerated mock type for the OrderUseCase type
type OrderUseCase struct {
	mock.Mock
}

// DeleteOrder provides a mock function with given fields: userID, productID
func (_m *OrderUseCase) DeleteOrder(userID int, productID int) (bool, error) {
	ret := _m.Called(userID, productID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int, int) bool); ok {
		r0 = rf(userID, productID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(userID, productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrder provides a mock function with given fields: orderId
func (_m *OrderUseCase) GetOrder(orderId int) (domain.ProductOrders, error) {
	ret := _m.Called(orderId)

	var r0 domain.ProductOrders
	if rf, ok := ret.Get(0).(func(int) domain.ProductOrders); ok {
		r0 = rf(orderId)
	} else {
		r0 = ret.Get(0).(domain.ProductOrders)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(orderId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostOrder provides a mock function with given fields: neworder, data
func (_m *OrderUseCase) PostOrder(neworder domain.Order, data []domain.ProductOrders) (int, string) {
	ret := _m.Called(neworder, data)

	var r0 int
	if rf, ok := ret.Get(0).(func(domain.Order, []domain.ProductOrders) int); ok {
		r0 = rf(neworder, data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(domain.Order, []domain.ProductOrders) string); ok {
		r1 = rf(neworder, data)
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// Sum provides a mock function with given fields: neworder
func (_m *OrderUseCase) Sum(neworder domain.Order) int {
	ret := _m.Called(neworder)

	var r0 int
	if rf, ok := ret.Get(0).(func(domain.Order) int); ok {
		r0 = rf(neworder)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type mockConstructorTestingTNewOrderUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewOrderUseCase creates a new instance of OrderUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrderUseCase(t mockConstructorTestingTNewOrderUseCase) *OrderUseCase {
	mock := &OrderUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
