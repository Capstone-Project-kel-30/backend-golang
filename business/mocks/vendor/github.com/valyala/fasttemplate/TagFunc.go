// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// TagFunc is an autogenerated mock type for the TagFunc type
type TagFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: w, tag
func (_m *TagFunc) Execute(w io.Writer, tag string) (int, error) {
	ret := _m.Called(w, tag)

	var r0 int
	if rf, ok := ret.Get(0).(func(io.Writer, string) int); ok {
		r0 = rf(w, tag)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(io.Writer, string) error); ok {
		r1 = rf(w, tag)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTagFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewTagFunc creates a new instance of TagFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTagFunc(t mockConstructorTestingTNewTagFunc) *TagFunc {
	mock := &TagFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
