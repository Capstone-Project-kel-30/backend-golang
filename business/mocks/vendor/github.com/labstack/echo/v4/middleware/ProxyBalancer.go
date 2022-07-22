// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"

	mock "github.com/stretchr/testify/mock"
)

// ProxyBalancer is an autogenerated mock type for the ProxyBalancer type
type ProxyBalancer struct {
	mock.Mock
}

// AddTarget provides a mock function with given fields: _a0
func (_m *ProxyBalancer) AddTarget(_a0 *middleware.ProxyTarget) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*middleware.ProxyTarget) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Next provides a mock function with given fields: _a0
func (_m *ProxyBalancer) Next(_a0 echo.Context) *middleware.ProxyTarget {
	ret := _m.Called(_a0)

	var r0 *middleware.ProxyTarget
	if rf, ok := ret.Get(0).(func(echo.Context) *middleware.ProxyTarget); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*middleware.ProxyTarget)
		}
	}

	return r0
}

// RemoveTarget provides a mock function with given fields: _a0
func (_m *ProxyBalancer) RemoveTarget(_a0 string) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewProxyBalancer interface {
	mock.TestingT
	Cleanup(func())
}

// NewProxyBalancer creates a new instance of ProxyBalancer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProxyBalancer(t mockConstructorTestingTNewProxyBalancer) *ProxyBalancer {
	mock := &ProxyBalancer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}