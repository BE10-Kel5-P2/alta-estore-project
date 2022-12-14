// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "altaproject2/domain"

	mock "github.com/stretchr/testify/mock"
)

// CartData is an autogenerated mock type for the CartData type
type CartData struct {
	mock.Mock
}

// CheckDuplicate provides a mock function with given fields: newcart
func (_m *CartData) CheckDuplicate(newcart domain.Cart) bool {
	ret := _m.Called(newcart)

	var r0 bool
	if rf, ok := ret.Get(0).(func(domain.Cart) bool); ok {
		r0 = rf(newcart)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DeleteData provides a mock function with given fields: userID, productID
func (_m *CartData) DeleteData(userID int, productID int) bool {
	ret := _m.Called(userID, productID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int, int) bool); ok {
		r0 = rf(userID, productID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetData provides a mock function with given fields: userid
func (_m *CartData) GetData(userid int) domain.Cart {
	ret := _m.Called(userid)

	var r0 domain.Cart
	if rf, ok := ret.Get(0).(func(int) domain.Cart); ok {
		r0 = rf(userid)
	} else {
		r0 = ret.Get(0).(domain.Cart)
	}

	return r0
}

// GetDataProduct provides a mock function with given fields: userid
func (_m *CartData) GetDataProduct(userid int) []domain.CartProduct {
	ret := _m.Called(userid)

	var r0 []domain.CartProduct
	if rf, ok := ret.Get(0).(func(int) []domain.CartProduct); ok {
		r0 = rf(userid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.CartProduct)
		}
	}

	return r0
}

// PostData provides a mock function with given fields: newcart
func (_m *CartData) PostData(newcart domain.Cart) domain.Cart {
	ret := _m.Called(newcart)

	var r0 domain.Cart
	if rf, ok := ret.Get(0).(func(domain.Cart) domain.Cart); ok {
		r0 = rf(newcart)
	} else {
		r0 = ret.Get(0).(domain.Cart)
	}

	return r0
}

// UpdateData provides a mock function with given fields: newUpdate, productId
func (_m *CartData) UpdateData(newUpdate domain.Cart, productId int) domain.Cart {
	ret := _m.Called(newUpdate, productId)

	var r0 domain.Cart
	if rf, ok := ret.Get(0).(func(domain.Cart, int) domain.Cart); ok {
		r0 = rf(newUpdate, productId)
	} else {
		r0 = ret.Get(0).(domain.Cart)
	}

	return r0
}

type mockConstructorTestingTNewCartData interface {
	mock.TestingT
	Cleanup(func())
}

// NewCartData creates a new instance of CartData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCartData(t mockConstructorTestingTNewCartData) *CartData {
	mock := &CartData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
