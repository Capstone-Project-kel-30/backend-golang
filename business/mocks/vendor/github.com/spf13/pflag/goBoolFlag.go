// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// goBoolFlag is an autogenerated mock type for the goBoolFlag type
type goBoolFlag struct {
	mock.Mock
}

// IsBoolFlag provides a mock function with given fields:
func (_m *goBoolFlag) IsBoolFlag() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Set provides a mock function with given fields: _a0
func (_m *goBoolFlag) Set(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *goBoolFlag) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTnewGoBoolFlag interface {
	mock.TestingT
	Cleanup(func())
}

// newGoBoolFlag creates a new instance of goBoolFlag. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newGoBoolFlag(t mockConstructorTestingTnewGoBoolFlag) *goBoolFlag {
	mock := &goBoolFlag{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
