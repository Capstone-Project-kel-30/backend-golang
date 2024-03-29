// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// TxBeginner is an autogenerated mock type for the TxBeginner type
type TxBeginner struct {
	mock.Mock
}

// BeginTx provides a mock function with given fields: ctx, opts
func (_m *TxBeginner) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	ret := _m.Called(ctx, opts)

	var r0 *sql.Tx
	if rf, ok := ret.Get(0).(func(context.Context, *sql.TxOptions) *sql.Tx); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Tx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sql.TxOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTxBeginner interface {
	mock.TestingT
	Cleanup(func())
}

// NewTxBeginner creates a new instance of TxBeginner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTxBeginner(t mockConstructorTestingTNewTxBeginner) *TxBeginner {
	mock := &TxBeginner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
