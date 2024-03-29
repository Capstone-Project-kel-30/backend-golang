// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// pipeBuffer is an autogenerated mock type for the pipeBuffer type
type pipeBuffer struct {
	mock.Mock
}

// Len provides a mock function with given fields:
func (_m *pipeBuffer) Len() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Read provides a mock function with given fields: p
func (_m *pipeBuffer) Read(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Write provides a mock function with given fields: p
func (_m *pipeBuffer) Write(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewPipeBuffer interface {
	mock.TestingT
	Cleanup(func())
}

// newPipeBuffer creates a new instance of pipeBuffer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newPipeBuffer(t mockConstructorTestingTnewPipeBuffer) *pipeBuffer {
	mock := &pipeBuffer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
