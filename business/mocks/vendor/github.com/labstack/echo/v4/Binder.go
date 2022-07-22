// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// Binder is an autogenerated mock type for the Binder type
type Binder struct {
	mock.Mock
}

// Bind provides a mock function with given fields: i, c
func (_m *Binder) Bind(i interface{}, c echo.Context) error {
	ret := _m.Called(i, c)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, echo.Context) error); ok {
		r0 = rf(i, c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewBinder interface {
	mock.TestingT
	Cleanup(func())
}

// NewBinder creates a new instance of Binder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBinder(t mockConstructorTestingTNewBinder) *Binder {
	mock := &Binder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
