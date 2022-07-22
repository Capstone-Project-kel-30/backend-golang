// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	http2 "golang.org/x/net/http2"
)

// frameParser is an autogenerated mock type for the frameParser type
type frameParser struct {
	mock.Mock
}

// Execute provides a mock function with given fields: fc, fh, countError, payload
func (_m *frameParser) Execute(fc *http2.frameCache, fh http2.FrameHeader, countError func(string), payload []byte) (http2.Frame, error) {
	ret := _m.Called(fc, fh, countError, payload)

	var r0 http2.Frame
	if rf, ok := ret.Get(0).(func(*http2.frameCache, http2.FrameHeader, func(string), []byte) http2.Frame); ok {
		r0 = rf(fc, fh, countError, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http2.Frame)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*http2.frameCache, http2.FrameHeader, func(string), []byte) error); ok {
		r1 = rf(fc, fh, countError, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewFrameParser interface {
	mock.TestingT
	Cleanup(func())
}

// newFrameParser creates a new instance of frameParser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newFrameParser(t mockConstructorTestingTnewFrameParser) *frameParser {
	mock := &frameParser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}