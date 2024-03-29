// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Marshaler is an autogenerated mock type for the Marshaler type
type Marshaler struct {
	mock.Mock
}

// MarshalTOML provides a mock function with given fields:
func (_m *Marshaler) MarshalTOML() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMarshaler interface {
	mock.TestingT
	Cleanup(func())
}

// NewMarshaler creates a new instance of Marshaler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMarshaler(t mockConstructorTestingTNewMarshaler) *Marshaler {
	mock := &Marshaler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
