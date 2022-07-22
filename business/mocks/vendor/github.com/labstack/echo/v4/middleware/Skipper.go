// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// Skipper is an autogenerated mock type for the Skipper type
type Skipper struct {
	mock.Mock
}

// Execute provides a mock function with given fields: c
func (_m *Skipper) Execute(c echo.Context) bool {
	ret := _m.Called(c)

	var r0 bool
	if rf, ok := ret.Get(0).(func(echo.Context) bool); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewSkipper interface {
	mock.TestingT
	Cleanup(func())
}

// NewSkipper creates a new instance of Skipper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSkipper(t mockConstructorTestingTNewSkipper) *Skipper {
	mock := &Skipper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
