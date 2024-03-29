// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	sql "database/sql"

	mock "github.com/stretchr/testify/mock"
)

// GetDBConnector is an autogenerated mock type for the GetDBConnector type
type GetDBConnector struct {
	mock.Mock
}

// GetDBConn provides a mock function with given fields:
func (_m *GetDBConnector) GetDBConn() (*sql.DB, error) {
	ret := _m.Called()

	var r0 *sql.DB
	if rf, ok := ret.Get(0).(func() *sql.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.DB)
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

type mockConstructorTestingTNewGetDBConnector interface {
	mock.TestingT
	Cleanup(func())
}

// NewGetDBConnector creates a new instance of GetDBConnector. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGetDBConnector(t mockConstructorTestingTNewGetDBConnector) *GetDBConnector {
	mock := &GetDBConnector{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
