// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// JWTErrorHandlerWithContext is an autogenerated mock type for the JWTErrorHandlerWithContext type
type JWTErrorHandlerWithContext struct {
	mock.Mock
}

// Execute provides a mock function with given fields: err, c
func (_m *JWTErrorHandlerWithContext) Execute(err error, c echo.Context) error {
	ret := _m.Called(err, c)

	var r0 error
	if rf, ok := ret.Get(0).(func(error, echo.Context) error); ok {
		r0 = rf(err, c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewJWTErrorHandlerWithContext interface {
	mock.TestingT
	Cleanup(func())
}

// NewJWTErrorHandlerWithContext creates a new instance of JWTErrorHandlerWithContext. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewJWTErrorHandlerWithContext(t mockConstructorTestingTNewJWTErrorHandlerWithContext) *JWTErrorHandlerWithContext {
	mock := &JWTErrorHandlerWithContext{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
