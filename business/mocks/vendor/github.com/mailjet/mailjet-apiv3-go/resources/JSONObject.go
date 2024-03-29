// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// JSONObject is an autogenerated mock type for the JSONObject type
type JSONObject struct {
	mock.Mock
}

type mockConstructorTestingTNewJSONObject interface {
	mock.TestingT
	Cleanup(func())
}

// NewJSONObject creates a new instance of JSONObject. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewJSONObject(t mockConstructorTestingTNewJSONObject) *JSONObject {
	mock := &JSONObject{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
