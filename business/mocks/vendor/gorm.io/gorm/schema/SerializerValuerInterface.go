// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	reflect "reflect"

	mock "github.com/stretchr/testify/mock"

	schema "gorm.io/gorm/schema"
)

// SerializerValuerInterface is an autogenerated mock type for the SerializerValuerInterface type
type SerializerValuerInterface struct {
	mock.Mock
}

// Value provides a mock function with given fields: ctx, field, dst, fieldValue
func (_m *SerializerValuerInterface) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	ret := _m.Called(ctx, field, dst, fieldValue)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, *schema.Field, reflect.Value, interface{}) interface{}); ok {
		r0 = rf(ctx, field, dst, fieldValue)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *schema.Field, reflect.Value, interface{}) error); ok {
		r1 = rf(ctx, field, dst, fieldValue)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewSerializerValuerInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewSerializerValuerInterface creates a new instance of SerializerValuerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSerializerValuerInterface(t mockConstructorTestingTNewSerializerValuerInterface) *SerializerValuerInterface {
	mock := &SerializerValuerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}