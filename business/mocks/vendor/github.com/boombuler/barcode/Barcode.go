// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	color "image/color"

	barcode "github.com/boombuler/barcode"

	image "image"

	mock "github.com/stretchr/testify/mock"
)

// Barcode is an autogenerated mock type for the Barcode type
type Barcode struct {
	mock.Mock
}

// At provides a mock function with given fields: x, y
func (_m *Barcode) At(x int, y int) color.Color {
	ret := _m.Called(x, y)

	var r0 color.Color
	if rf, ok := ret.Get(0).(func(int, int) color.Color); ok {
		r0 = rf(x, y)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(color.Color)
		}
	}

	return r0
}

// Bounds provides a mock function with given fields:
func (_m *Barcode) Bounds() image.Rectangle {
	ret := _m.Called()

	var r0 image.Rectangle
	if rf, ok := ret.Get(0).(func() image.Rectangle); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(image.Rectangle)
	}

	return r0
}

// ColorModel provides a mock function with given fields:
func (_m *Barcode) ColorModel() color.Model {
	ret := _m.Called()

	var r0 color.Model
	if rf, ok := ret.Get(0).(func() color.Model); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(color.Model)
		}
	}

	return r0
}

// Content provides a mock function with given fields:
func (_m *Barcode) Content() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Metadata provides a mock function with given fields:
func (_m *Barcode) Metadata() barcode.Metadata {
	ret := _m.Called()

	var r0 barcode.Metadata
	if rf, ok := ret.Get(0).(func() barcode.Metadata); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(barcode.Metadata)
	}

	return r0
}

type mockConstructorTestingTNewBarcode interface {
	mock.TestingT
	Cleanup(func())
}

// NewBarcode creates a new instance of Barcode. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBarcode(t mockConstructorTestingTNewBarcode) *Barcode {
	mock := &Barcode{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}