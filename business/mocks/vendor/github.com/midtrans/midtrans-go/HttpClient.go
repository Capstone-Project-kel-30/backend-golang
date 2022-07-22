// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	io "io"

	midtrans "github.com/midtrans/midtrans-go"
	mock "github.com/stretchr/testify/mock"
)

// HttpClient is an autogenerated mock type for the HttpClient type
type HttpClient struct {
	mock.Mock
}

// Call provides a mock function with given fields: method, url, apiKey, options, body, result
func (_m *HttpClient) Call(method string, url string, apiKey *string, options *midtrans.ConfigOptions, body io.Reader, result interface{}) *midtrans.Error {
	ret := _m.Called(method, url, apiKey, options, body, result)

	var r0 *midtrans.Error
	if rf, ok := ret.Get(0).(func(string, string, *string, *midtrans.ConfigOptions, io.Reader, interface{}) *midtrans.Error); ok {
		r0 = rf(method, url, apiKey, options, body, result)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*midtrans.Error)
		}
	}

	return r0
}

type mockConstructorTestingTNewHttpClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewHttpClient creates a new instance of HttpClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHttpClient(t mockConstructorTestingTNewHttpClient) *HttpClient {
	mock := &HttpClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
