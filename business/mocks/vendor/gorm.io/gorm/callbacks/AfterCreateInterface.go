// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// AfterCreateInterface is an autogenerated mock type for the AfterCreateInterface type
type AfterCreateInterface struct {
	mock.Mock
}

// AfterCreate provides a mock function with given fields: _a0
func (_m *AfterCreateInterface) AfterCreate(_a0 *gorm.DB) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAfterCreateInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewAfterCreateInterface creates a new instance of AfterCreateInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAfterCreateInterface(t mockConstructorTestingTNewAfterCreateInterface) *AfterCreateInterface {
	mock := &AfterCreateInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}