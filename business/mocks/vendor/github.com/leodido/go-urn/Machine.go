// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	urn "github.com/leodido/go-urn"
	mock "github.com/stretchr/testify/mock"
)

// Machine is an autogenerated mock type for the Machine type
type Machine struct {
	mock.Mock
}

// Error provides a mock function with given fields:
func (_m *Machine) Error() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Parse provides a mock function with given fields: input
func (_m *Machine) Parse(input []byte) (*urn.URN, error) {
	ret := _m.Called(input)

	var r0 *urn.URN
	if rf, ok := ret.Get(0).(func([]byte) *urn.URN); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*urn.URN)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMachine interface {
	mock.TestingT
	Cleanup(func())
}

// NewMachine creates a new instance of Machine. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMachine(t mockConstructorTestingTNewMachine) *Machine {
	mock := &Machine{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
