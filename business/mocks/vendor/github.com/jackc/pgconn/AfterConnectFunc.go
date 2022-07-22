// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	pgconn "github.com/jackc/pgconn"
	mock "github.com/stretchr/testify/mock"
)

// AfterConnectFunc is an autogenerated mock type for the AfterConnectFunc type
type AfterConnectFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, _a1
func (_m *AfterConnectFunc) Execute(ctx context.Context, _a1 *pgconn.PgConn) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *pgconn.PgConn) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAfterConnectFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewAfterConnectFunc creates a new instance of AfterConnectFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAfterConnectFunc(t mockConstructorTestingTNewAfterConnectFunc) *AfterConnectFunc {
	mock := &AfterConnectFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
