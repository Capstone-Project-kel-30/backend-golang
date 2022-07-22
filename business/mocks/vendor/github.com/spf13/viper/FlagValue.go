// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FlagValue is an autogenerated mock type for the FlagValue type
type FlagValue struct {
	mock.Mock
}

// HasChanged provides a mock function with given fields:
func (_m *FlagValue) HasChanged() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *FlagValue) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ValueString provides a mock function with given fields:
func (_m *FlagValue) ValueString() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ValueType provides a mock function with given fields:
func (_m *FlagValue) ValueType() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewFlagValue interface {
	mock.TestingT
	Cleanup(func())
}

// NewFlagValue creates a new instance of FlagValue. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFlagValue(t mockConstructorTestingTNewFlagValue) *FlagValue {
	mock := &FlagValue{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
