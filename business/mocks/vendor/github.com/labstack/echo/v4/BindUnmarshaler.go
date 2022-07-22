// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// BindUnmarshaler is an autogenerated mock type for the BindUnmarshaler type
type BindUnmarshaler struct {
	mock.Mock
}

// UnmarshalParam provides a mock function with given fields: param
func (_m *BindUnmarshaler) UnmarshalParam(param string) error {
	ret := _m.Called(param)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(param)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewBindUnmarshaler interface {
	mock.TestingT
	Cleanup(func())
}

// NewBindUnmarshaler creates a new instance of BindUnmarshaler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBindUnmarshaler(t mockConstructorTestingTNewBindUnmarshaler) *BindUnmarshaler {
	mock := &BindUnmarshaler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
