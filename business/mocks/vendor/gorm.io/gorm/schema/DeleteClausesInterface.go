// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	clause "gorm.io/gorm/clause"

	schema "gorm.io/gorm/schema"
)

// DeleteClausesInterface is an autogenerated mock type for the DeleteClausesInterface type
type DeleteClausesInterface struct {
	mock.Mock
}

// DeleteClauses provides a mock function with given fields: _a0
func (_m *DeleteClausesInterface) DeleteClauses(_a0 *schema.Field) []clause.Interface {
	ret := _m.Called(_a0)

	var r0 []clause.Interface
	if rf, ok := ret.Get(0).(func(*schema.Field) []clause.Interface); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]clause.Interface)
		}
	}

	return r0
}

type mockConstructorTestingTNewDeleteClausesInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewDeleteClausesInterface creates a new instance of DeleteClausesInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDeleteClausesInterface(t mockConstructorTestingTNewDeleteClausesInterface) *DeleteClausesInterface {
	mock := &DeleteClausesInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
