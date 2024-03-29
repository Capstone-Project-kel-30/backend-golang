// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FilterFunc is an autogenerated mock type for the FilterFunc type
type FilterFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ns
func (_m *FilterFunc) Execute(ns []byte) bool {
	ret := _m.Called(ns)

	var r0 bool
	if rf, ok := ret.Get(0).(func([]byte) bool); ok {
		r0 = rf(ns)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewFilterFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewFilterFunc creates a new instance of FilterFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFilterFunc(t mockConstructorTestingTNewFilterFunc) *FilterFunc {
	mock := &FilterFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
