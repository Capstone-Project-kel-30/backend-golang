// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	reflect "reflect"

	mock "github.com/stretchr/testify/mock"

	validator "github.com/go-playground/validator"
)

// StructLevel is an autogenerated mock type for the StructLevel type
type StructLevel struct {
	mock.Mock
}

// Current provides a mock function with given fields:
func (_m *StructLevel) Current() reflect.Value {
	ret := _m.Called()

	var r0 reflect.Value
	if rf, ok := ret.Get(0).(func() reflect.Value); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(reflect.Value)
	}

	return r0
}

// ExtractType provides a mock function with given fields: field
func (_m *StructLevel) ExtractType(field reflect.Value) (reflect.Value, reflect.Kind, bool) {
	ret := _m.Called(field)

	var r0 reflect.Value
	if rf, ok := ret.Get(0).(func(reflect.Value) reflect.Value); ok {
		r0 = rf(field)
	} else {
		r0 = ret.Get(0).(reflect.Value)
	}

	var r1 reflect.Kind
	if rf, ok := ret.Get(1).(func(reflect.Value) reflect.Kind); ok {
		r1 = rf(field)
	} else {
		r1 = ret.Get(1).(reflect.Kind)
	}

	var r2 bool
	if rf, ok := ret.Get(2).(func(reflect.Value) bool); ok {
		r2 = rf(field)
	} else {
		r2 = ret.Get(2).(bool)
	}

	return r0, r1, r2
}

// Parent provides a mock function with given fields:
func (_m *StructLevel) Parent() reflect.Value {
	ret := _m.Called()

	var r0 reflect.Value
	if rf, ok := ret.Get(0).(func() reflect.Value); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(reflect.Value)
	}

	return r0
}

// ReportError provides a mock function with given fields: field, fieldName, structFieldName, tag, param
func (_m *StructLevel) ReportError(field interface{}, fieldName string, structFieldName string, tag string, param string) {
	_m.Called(field, fieldName, structFieldName, tag, param)
}

// ReportValidationErrors provides a mock function with given fields: relativeNamespace, relativeActualNamespace, errs
func (_m *StructLevel) ReportValidationErrors(relativeNamespace string, relativeActualNamespace string, errs validator.ValidationErrors) {
	_m.Called(relativeNamespace, relativeActualNamespace, errs)
}

// Top provides a mock function with given fields:
func (_m *StructLevel) Top() reflect.Value {
	ret := _m.Called()

	var r0 reflect.Value
	if rf, ok := ret.Get(0).(func() reflect.Value); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(reflect.Value)
	}

	return r0
}

// Validator provides a mock function with given fields:
func (_m *StructLevel) Validator() *validator.Validate {
	ret := _m.Called()

	var r0 *validator.Validate
	if rf, ok := ret.Get(0).(func() *validator.Validate); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*validator.Validate)
		}
	}

	return r0
}

type mockConstructorTestingTNewStructLevel interface {
	mock.TestingT
	Cleanup(func())
}

// NewStructLevel creates a new instance of StructLevel. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStructLevel(t mockConstructorTestingTNewStructLevel) *StructLevel {
	mock := &StructLevel{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
