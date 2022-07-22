// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// Extractor is an autogenerated mock type for the Extractor type
type Extractor struct {
	mock.Mock
}

// Execute provides a mock function with given fields: context
func (_m *Extractor) Execute(context echo.Context) (string, error) {
	ret := _m.Called(context)

	var r0 string
	if rf, ok := ret.Get(0).(func(echo.Context) string); ok {
		r0 = rf(context)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(echo.Context) error); ok {
		r1 = rf(context)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewExtractor interface {
	mock.TestingT
	Cleanup(func())
}

// NewExtractor creates a new instance of Extractor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExtractor(t mockConstructorTestingTNewExtractor) *Extractor {
	mock := &Extractor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
