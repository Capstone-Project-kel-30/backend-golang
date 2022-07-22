// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// AfterDeleteInterface is an autogenerated mock type for the AfterDeleteInterface type
type AfterDeleteInterface struct {
	mock.Mock
}

// AfterDelete provides a mock function with given fields: _a0
func (_m *AfterDeleteInterface) AfterDelete(_a0 *gorm.DB) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAfterDeleteInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewAfterDeleteInterface creates a new instance of AfterDeleteInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAfterDeleteInterface(t mockConstructorTestingTNewAfterDeleteInterface) *AfterDeleteInterface {
	mock := &AfterDeleteInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
