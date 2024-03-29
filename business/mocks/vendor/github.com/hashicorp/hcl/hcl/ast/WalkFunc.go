// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	ast "github.com/hashicorp/hcl/hcl/ast"
	mock "github.com/stretchr/testify/mock"
)

// WalkFunc is an autogenerated mock type for the WalkFunc type
type WalkFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *WalkFunc) Execute(_a0 ast.Node) (ast.Node, bool) {
	ret := _m.Called(_a0)

	var r0 ast.Node
	if rf, ok := ret.Get(0).(func(ast.Node) ast.Node); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ast.Node)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(ast.Node) bool); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

type mockConstructorTestingTNewWalkFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewWalkFunc creates a new instance of WalkFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWalkFunc(t mockConstructorTestingTNewWalkFunc) *WalkFunc {
	mock := &WalkFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
