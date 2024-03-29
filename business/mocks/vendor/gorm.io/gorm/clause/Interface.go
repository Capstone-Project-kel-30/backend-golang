// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	clause "gorm.io/gorm/clause"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// Build provides a mock function with given fields: _a0
func (_m *Interface) Build(_a0 clause.Builder) {
	_m.Called(_a0)
}

// MergeClause provides a mock function with given fields: _a0
func (_m *Interface) MergeClause(_a0 *clause.Clause) {
	_m.Called(_a0)
}

// Name provides a mock function with given fields:
func (_m *Interface) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewInterface creates a new instance of Interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInterface(t mockConstructorTestingTNewInterface) *Interface {
	mock := &Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
