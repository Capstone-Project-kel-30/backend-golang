// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// DecodeHookFunc is an autogenerated mock type for the DecodeHookFunc type
type DecodeHookFunc struct {
	mock.Mock
}

type mockConstructorTestingTNewDecodeHookFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewDecodeHookFunc creates a new instance of DecodeHookFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDecodeHookFunc(t mockConstructorTestingTNewDecodeHookFunc) *DecodeHookFunc {
	mock := &DecodeHookFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
