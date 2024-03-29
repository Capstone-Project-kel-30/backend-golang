// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// TxCommitter is an autogenerated mock type for the TxCommitter type
type TxCommitter struct {
	mock.Mock
}

// Commit provides a mock function with given fields:
func (_m *TxCommitter) Commit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Rollback provides a mock function with given fields:
func (_m *TxCommitter) Rollback() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTxCommitter interface {
	mock.TestingT
	Cleanup(func())
}

// NewTxCommitter creates a new instance of TxCommitter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTxCommitter(t mockConstructorTestingTNewTxCommitter) *TxCommitter {
	mock := &TxCommitter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
