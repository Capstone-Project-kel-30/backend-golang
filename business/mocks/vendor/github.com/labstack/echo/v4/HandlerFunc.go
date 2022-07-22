// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// HandlerFunc is an autogenerated mock type for the HandlerFunc type
type HandlerFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: c
func (_m *HandlerFunc) Execute(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewHandlerFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewHandlerFunc creates a new instance of HandlerFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHandlerFunc(t mockConstructorTestingTNewHandlerFunc) *HandlerFunc {
	mock := &HandlerFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
