// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	pgtype "github.com/jackc/pgtype"
	mock "github.com/stretchr/testify/mock"
)

// BinaryDecoder is an autogenerated mock type for the BinaryDecoder type
type BinaryDecoder struct {
	mock.Mock
}

// DecodeBinary provides a mock function with given fields: ci, src
func (_m *BinaryDecoder) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {
	ret := _m.Called(ci, src)

	var r0 error
	if rf, ok := ret.Get(0).(func(*pgtype.ConnInfo, []byte) error); ok {
		r0 = rf(ci, src)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewBinaryDecoder interface {
	mock.TestingT
	Cleanup(func())
}

// NewBinaryDecoder creates a new instance of BinaryDecoder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBinaryDecoder(t mockConstructorTestingTNewBinaryDecoder) *BinaryDecoder {
	mock := &BinaryDecoder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}