// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// redirectLogic is an autogenerated mock type for the redirectLogic type
type redirectLogic struct {
	mock.Mock
}

// Execute provides a mock function with given fields: scheme, host, uri
func (_m *redirectLogic) Execute(scheme string, host string, uri string) (bool, string) {
	ret := _m.Called(scheme, host, uri)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, string) bool); ok {
		r0 = rf(scheme, host, uri)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string, string, string) string); ok {
		r1 = rf(scheme, host, uri)
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

type mockConstructorTestingTnewRedirectLogic interface {
	mock.TestingT
	Cleanup(func())
}

// newRedirectLogic creates a new instance of redirectLogic. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newRedirectLogic(t mockConstructorTestingTnewRedirectLogic) *redirectLogic {
	mock := &redirectLogic{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
