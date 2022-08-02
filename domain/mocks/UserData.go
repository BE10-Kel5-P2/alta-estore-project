// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "altaproject2/domain"

	mock "github.com/stretchr/testify/mock"
)

// UserData is an autogenerated mock type for the UserData type
type UserData struct {
	mock.Mock
}

// CheckDuplicate provides a mock function with given fields: newuser
func (_m *UserData) CheckDuplicate(newuser domain.User) bool {
	ret := _m.Called(newuser)

	var r0 bool
	if rf, ok := ret.Get(0).(func(domain.User) bool); ok {
		r0 = rf(newuser)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Delete provides a mock function with given fields: userID
func (_m *UserData) Delete(userID int) bool {
	ret := _m.Called(userID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetPasswordData provides a mock function with given fields: name
func (_m *UserData) GetPasswordData(name string) string {
	ret := _m.Called(name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetProfile provides a mock function with given fields: userID
func (_m *UserData) GetProfile(userID int) (domain.User, error) {
	ret := _m.Called(userID)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(int) domain.User); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: userdata
func (_m *UserData) Login(userdata domain.User) domain.User {
	ret := _m.Called(userdata)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(userdata)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	return r0
}

// RegisterData provides a mock function with given fields: newuser
func (_m *UserData) RegisterData(newuser domain.User) domain.User {
	ret := _m.Called(newuser)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(newuser)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	return r0
}

// UpdateUserData provides a mock function with given fields: newuser
func (_m *UserData) UpdateUserData(newuser domain.User) domain.User {
	ret := _m.Called(newuser)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(newuser)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	return r0
}

type mockConstructorTestingTNewUserData interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserData creates a new instance of UserData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserData(t mockConstructorTestingTNewUserData) *UserData {
	mock := &UserData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
