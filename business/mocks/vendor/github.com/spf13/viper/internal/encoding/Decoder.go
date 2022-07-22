// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Decoder is an autogenerated mock type for the Decoder type
type Decoder struct {
	mock.Mock
}

// Decode provides a mock function with given fields: b, v
func (_m *Decoder) Decode(b []byte, v map[string]interface{}) error {
	ret := _m.Called(b, v)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, map[string]interface{}) error); ok {
		r0 = rf(b, v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDecoder interface {
	mock.TestingT
	Cleanup(func())
}

// NewDecoder creates a new instance of Decoder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDecoder(t mockConstructorTestingTNewDecoder) *Decoder {
	mock := &Decoder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
