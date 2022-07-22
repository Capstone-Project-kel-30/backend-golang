// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	jwt "github.com/golang-jwt/jwt"
	mock "github.com/stretchr/testify/mock"
)

// Keyfunc is an autogenerated mock type for the Keyfunc type
type Keyfunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *Keyfunc) Execute(_a0 *jwt.Token) (interface{}, error) {
	ret := _m.Called(_a0)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(*jwt.Token) interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*jwt.Token) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewKeyfunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewKeyfunc creates a new instance of Keyfunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewKeyfunc(t mockConstructorTestingTNewKeyfunc) *Keyfunc {
	mock := &Keyfunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
