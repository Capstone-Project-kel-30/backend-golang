// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// SavePointerDialectorInterface is an autogenerated mock type for the SavePointerDialectorInterface type
type SavePointerDialectorInterface struct {
	mock.Mock
}

// RollbackTo provides a mock function with given fields: tx, name
func (_m *SavePointerDialectorInterface) RollbackTo(tx *gorm.DB, name string) error {
	ret := _m.Called(tx, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) error); ok {
		r0 = rf(tx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SavePoint provides a mock function with given fields: tx, name
func (_m *SavePointerDialectorInterface) SavePoint(tx *gorm.DB, name string) error {
	ret := _m.Called(tx, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) error); ok {
		r0 = rf(tx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewSavePointerDialectorInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewSavePointerDialectorInterface creates a new instance of SavePointerDialectorInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSavePointerDialectorInterface(t mockConstructorTestingTNewSavePointerDialectorInterface) *SavePointerDialectorInterface {
	mock := &SavePointerDialectorInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
