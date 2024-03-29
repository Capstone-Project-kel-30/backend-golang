// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// OrderOption is an autogenerated mock type for the OrderOption type
type OrderOption struct {
	mock.Mock
}

// privateOrderOpt provides a mock function with given fields:
func (_m *OrderOption) privateOrderOpt() {
	_m.Called()
}

type mockConstructorTestingTNewOrderOption interface {
	mock.TestingT
	Cleanup(func())
}

// NewOrderOption creates a new instance of OrderOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrderOption(t mockConstructorTestingTNewOrderOption) *OrderOption {
	mock := &OrderOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
