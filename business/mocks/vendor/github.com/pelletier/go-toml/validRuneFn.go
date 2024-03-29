// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// validRuneFn is an autogenerated mock type for the validRuneFn type
type validRuneFn struct {
	mock.Mock
}

// Execute provides a mock function with given fields: r
func (_m *validRuneFn) Execute(r rune) bool {
	ret := _m.Called(r)

	var r0 bool
	if rf, ok := ret.Get(0).(func(rune) bool); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTnewValidRuneFn interface {
	mock.TestingT
	Cleanup(func())
}

// newValidRuneFn creates a new instance of validRuneFn. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newValidRuneFn(t mockConstructorTestingTnewValidRuneFn) *validRuneFn {
	mock := &validRuneFn{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
