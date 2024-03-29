// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// headersOrContinuation is an autogenerated mock type for the headersOrContinuation type
type headersOrContinuation struct {
	mock.Mock
}

// HeaderBlockFragment provides a mock function with given fields:
func (_m *headersOrContinuation) HeaderBlockFragment() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// HeadersEnded provides a mock function with given fields:
func (_m *headersOrContinuation) HeadersEnded() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTnewHeadersOrContinuation interface {
	mock.TestingT
	Cleanup(func())
}

// newHeadersOrContinuation creates a new instance of headersOrContinuation. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newHeadersOrContinuation(t mockConstructorTestingTnewHeadersOrContinuation) *headersOrContinuation {
	mock := &headersOrContinuation{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
