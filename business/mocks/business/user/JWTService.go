// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	jwt "github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// JWTService is an autogenerated mock type for the JWTService type
type JWTService struct {
	mock.Mock
}

// GenerateToken provides a mock function with given fields: userID
func (_m *JWTService) GenerateToken(userID string) string {
	ret := _m.Called(userID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ValidateToken provides a mock function with given fields: token, ctx
func (_m *JWTService) ValidateToken(token string, ctx echo.Context) *jwt.Token {
	ret := _m.Called(token, ctx)

	var r0 *jwt.Token
	if rf, ok := ret.Get(0).(func(string, echo.Context) *jwt.Token); ok {
		r0 = rf(token, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	return r0
}

type mockConstructorTestingTNewJWTService interface {
	mock.TestingT
	Cleanup(func())
}

// NewJWTService creates a new instance of JWTService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewJWTService(t mockConstructorTestingTNewJWTService) *JWTService {
	mock := &JWTService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}