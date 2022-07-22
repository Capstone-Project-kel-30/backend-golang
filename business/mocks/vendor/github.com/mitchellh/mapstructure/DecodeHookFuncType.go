// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	reflect "reflect"

	mock "github.com/stretchr/testify/mock"
)

// DecodeHookFuncType is an autogenerated mock type for the DecodeHookFuncType type
type DecodeHookFuncType struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1, _a2
func (_m *DecodeHookFuncType) Execute(_a0 reflect.Type, _a1 reflect.Type, _a2 interface{}) (interface{}, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(reflect.Type, reflect.Type, interface{}) interface{}); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(reflect.Type, reflect.Type, interface{}) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDecodeHookFuncType interface {
	mock.TestingT
	Cleanup(func())
}

// NewDecodeHookFuncType creates a new instance of DecodeHookFuncType. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDecodeHookFuncType(t mockConstructorTestingTNewDecodeHookFuncType) *DecodeHookFuncType {
	mock := &DecodeHookFuncType{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
