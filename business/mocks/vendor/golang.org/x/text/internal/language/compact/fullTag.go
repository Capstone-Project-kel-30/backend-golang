// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	language "golang.org/x/text/internal/language"
)

// fullTag is an autogenerated mock type for the fullTag type
type fullTag struct {
	mock.Mock
}

// IsRoot provides a mock function with given fields:
func (_m *fullTag) IsRoot() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Parent provides a mock function with given fields:
func (_m *fullTag) Parent() language.Tag {
	ret := _m.Called()

	var r0 language.Tag
	if rf, ok := ret.Get(0).(func() language.Tag); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(language.Tag)
	}

	return r0
}

type mockConstructorTestingTnewFullTag interface {
	mock.TestingT
	Cleanup(func())
}

// newFullTag creates a new instance of fullTag. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newFullTag(t mockConstructorTestingTnewFullTag) *fullTag {
	mock := &fullTag{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}