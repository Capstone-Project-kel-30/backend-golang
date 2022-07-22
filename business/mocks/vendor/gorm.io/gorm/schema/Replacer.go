// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Replacer is an autogenerated mock type for the Replacer type
type Replacer struct {
	mock.Mock
}

// Replace provides a mock function with given fields: name
func (_m *Replacer) Replace(name string) string {
	ret := _m.Called(name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewReplacer interface {
	mock.TestingT
	Cleanup(func())
}

// NewReplacer creates a new instance of Replacer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReplacer(t mockConstructorTestingTNewReplacer) *Replacer {
	mock := &Replacer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}