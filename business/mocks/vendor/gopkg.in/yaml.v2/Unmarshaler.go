// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Unmarshaler is an autogenerated mock type for the Unmarshaler type
type Unmarshaler struct {
	mock.Mock
}

// UnmarshalYAML provides a mock function with given fields: unmarshal
func (_m *Unmarshaler) UnmarshalYAML(unmarshal func(interface{}) error) error {
	ret := _m.Called(unmarshal)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(interface{}) error) error); ok {
		r0 = rf(unmarshal)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUnmarshaler interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnmarshaler creates a new instance of Unmarshaler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnmarshaler(t mockConstructorTestingTNewUnmarshaler) *Unmarshaler {
	mock := &Unmarshaler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
