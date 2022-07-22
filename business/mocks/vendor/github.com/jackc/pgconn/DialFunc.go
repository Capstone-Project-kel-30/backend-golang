// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	net "net"

	mock "github.com/stretchr/testify/mock"
)

// DialFunc is an autogenerated mock type for the DialFunc type
type DialFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, network, addr
func (_m *DialFunc) Execute(ctx context.Context, network string, addr string) (net.Conn, error) {
	ret := _m.Called(ctx, network, addr)

	var r0 net.Conn
	if rf, ok := ret.Get(0).(func(context.Context, string, string) net.Conn); ok {
		r0 = rf(ctx, network, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.Conn)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, network, addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDialFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewDialFunc creates a new instance of DialFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDialFunc(t mockConstructorTestingTNewDialFunc) *DialFunc {
	mock := &DialFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
