// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SMTPClientInterface is an autogenerated mock type for the SMTPClientInterface type
type SMTPClientInterface struct {
	mock.Mock
}

// SendMail provides a mock function with given fields: from, to, msg
func (_m *SMTPClientInterface) SendMail(from string, to []string, msg []byte) error {
	ret := _m.Called(from, to, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string, []byte) error); ok {
		r0 = rf(from, to, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewSMTPClientInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewSMTPClientInterface creates a new instance of SMTPClientInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSMTPClientInterface(t mockConstructorTestingTNewSMTPClientInterface) *SMTPClientInterface {
	mock := &SMTPClientInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}