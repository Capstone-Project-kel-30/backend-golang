// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// headersEnder is an autogenerated mock type for the headersEnder type
type headersEnder struct {
	mock.Mock
}

// HeadersEnded provides a mock function with given fields:
func (_m *headersEnder) HeadersEnded() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTnewHeadersEnder interface {
	mock.TestingT
	Cleanup(func())
}

// newHeadersEnder creates a new instance of headersEnder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newHeadersEnder(t mockConstructorTestingTnewHeadersEnder) *headersEnder {
	mock := &headersEnder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
