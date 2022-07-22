// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// GormDataTypeInterface is an autogenerated mock type for the GormDataTypeInterface type
type GormDataTypeInterface struct {
	mock.Mock
}

// GormDataType provides a mock function with given fields:
func (_m *GormDataTypeInterface) GormDataType() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewGormDataTypeInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewGormDataTypeInterface creates a new instance of GormDataTypeInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGormDataTypeInterface(t mockConstructorTestingTNewGormDataTypeInterface) *GormDataTypeInterface {
	mock := &GormDataTypeInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
