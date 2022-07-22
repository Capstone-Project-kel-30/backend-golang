// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	logger "gorm.io/gorm/logger"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// Error provides a mock function with given fields: _a0, _a1, _a2
func (_m *Interface) Error(_a0 context.Context, _a1 string, _a2 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _a2...)
	_m.Called(_ca...)
}

// Info provides a mock function with given fields: _a0, _a1, _a2
func (_m *Interface) Info(_a0 context.Context, _a1 string, _a2 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _a2...)
	_m.Called(_ca...)
}

// LogMode provides a mock function with given fields: _a0
func (_m *Interface) LogMode(_a0 logger.LogLevel) logger.Interface {
	ret := _m.Called(_a0)

	var r0 logger.Interface
	if rf, ok := ret.Get(0).(func(logger.LogLevel) logger.Interface); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logger.Interface)
		}
	}

	return r0
}

// Trace provides a mock function with given fields: ctx, begin, fc, err
func (_m *Interface) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	_m.Called(ctx, begin, fc, err)
}

// Warn provides a mock function with given fields: _a0, _a1, _a2
func (_m *Interface) Warn(_a0 context.Context, _a1 string, _a2 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _a2...)
	_m.Called(_ca...)
}

type mockConstructorTestingTNewInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewInterface creates a new instance of Interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInterface(t mockConstructorTestingTNewInterface) *Interface {
	mock := &Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}