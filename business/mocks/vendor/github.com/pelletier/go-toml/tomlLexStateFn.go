// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// tomlLexStateFn is an autogenerated mock type for the tomlLexStateFn type
type tomlLexStateFn struct {
	mock.Mock
}

// Execute provides a mock function with given fields:
func (_m *tomlLexStateFn) Execute() toml.tomlLexStateFn {
	ret := _m.Called()

	var r0 toml.tomlLexStateFn
	if rf, ok := ret.Get(0).(func() toml.tomlLexStateFn); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(toml.tomlLexStateFn)
		}
	}

	return r0
}

type mockConstructorTestingTnewTomlLexStateFn interface {
	mock.TestingT
	Cleanup(func())
}

// newTomlLexStateFn creates a new instance of tomlLexStateFn. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newTomlLexStateFn(t mockConstructorTestingTnewTomlLexStateFn) *tomlLexStateFn {
	mock := &tomlLexStateFn{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
