// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// mapFunc is an autogenerated mock type for the mapFunc type
type mapFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *mapFunc) Execute(_a0 *cases.context) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*cases.context) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTnewMapFunc interface {
	mock.TestingT
	Cleanup(func())
}

// newMapFunc creates a new instance of mapFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMapFunc(t mockConstructorTestingTnewMapFunc) *mapFunc {
	mock := &mapFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}