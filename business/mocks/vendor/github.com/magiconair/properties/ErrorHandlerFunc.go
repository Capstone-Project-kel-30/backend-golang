// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ErrorHandlerFunc is an autogenerated mock type for the ErrorHandlerFunc type
type ErrorHandlerFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *ErrorHandlerFunc) Execute(_a0 error) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewErrorHandlerFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewErrorHandlerFunc creates a new instance of ErrorHandlerFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewErrorHandlerFunc(t mockConstructorTestingTNewErrorHandlerFunc) *ErrorHandlerFunc {
	mock := &ErrorHandlerFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}