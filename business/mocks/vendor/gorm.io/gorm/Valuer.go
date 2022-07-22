// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	clause "gorm.io/gorm/clause"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"
)

// Valuer is an autogenerated mock type for the Valuer type
type Valuer struct {
	mock.Mock
}

// GormValue provides a mock function with given fields: _a0, _a1
func (_m *Valuer) GormValue(_a0 context.Context, _a1 *gorm.DB) clause.Expr {
	ret := _m.Called(_a0, _a1)

	var r0 clause.Expr
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB) clause.Expr); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(clause.Expr)
	}

	return r0
}

type mockConstructorTestingTNewValuer interface {
	mock.TestingT
	Cleanup(func())
}

// NewValuer creates a new instance of Valuer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewValuer(t mockConstructorTestingTNewValuer) *Valuer {
	mock := &Valuer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
