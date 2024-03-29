// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// BackendMessage is an autogenerated mock type for the BackendMessage type
type BackendMessage struct {
	mock.Mock
}

// Backend provides a mock function with given fields:
func (_m *BackendMessage) Backend() {
	_m.Called()
}

// Decode provides a mock function with given fields: data
func (_m *BackendMessage) Decode(data []byte) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Encode provides a mock function with given fields: dst
func (_m *BackendMessage) Encode(dst []byte) []byte {
	ret := _m.Called(dst)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(dst)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

type mockConstructorTestingTNewBackendMessage interface {
	mock.TestingT
	Cleanup(func())
}

// NewBackendMessage creates a new instance of BackendMessage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBackendMessage(t mockConstructorTestingTNewBackendMessage) *BackendMessage {
	mock := &BackendMessage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
