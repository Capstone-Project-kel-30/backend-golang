// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
	http2 "golang.org/x/net/http2"
)

// clientConnPoolIdleCloser is an autogenerated mock type for the clientConnPoolIdleCloser type
type clientConnPoolIdleCloser struct {
	mock.Mock
}

// GetClientConn provides a mock function with given fields: req, addr
func (_m *clientConnPoolIdleCloser) GetClientConn(req *http.Request, addr string) (*http2.ClientConn, error) {
	ret := _m.Called(req, addr)

	var r0 *http2.ClientConn
	if rf, ok := ret.Get(0).(func(*http.Request, string) *http2.ClientConn); ok {
		r0 = rf(req, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http2.ClientConn)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*http.Request, string) error); ok {
		r1 = rf(req, addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarkDead provides a mock function with given fields: _a0
func (_m *clientConnPoolIdleCloser) MarkDead(_a0 *http2.ClientConn) {
	_m.Called(_a0)
}

// closeIdleConnections provides a mock function with given fields:
func (_m *clientConnPoolIdleCloser) closeIdleConnections() {
	_m.Called()
}

type mockConstructorTestingTnewClientConnPoolIdleCloser interface {
	mock.TestingT
	Cleanup(func())
}

// newClientConnPoolIdleCloser creates a new instance of clientConnPoolIdleCloser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newClientConnPoolIdleCloser(t mockConstructorTestingTnewClientConnPoolIdleCloser) *clientConnPoolIdleCloser {
	mock := &clientConnPoolIdleCloser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
