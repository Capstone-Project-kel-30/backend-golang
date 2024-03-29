// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// BeforeCreateInterface is an autogenerated mock type for the BeforeCreateInterface type
type BeforeCreateInterface struct {
	mock.Mock
}

// BeforeCreate provides a mock function with given fields: _a0
func (_m *BeforeCreateInterface) BeforeCreate(_a0 *gorm.DB) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewBeforeCreateInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewBeforeCreateInterface creates a new instance of BeforeCreateInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBeforeCreateInterface(t mockConstructorTestingTNewBeforeCreateInterface) *BeforeCreateInterface {
	mock := &BeforeCreateInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
