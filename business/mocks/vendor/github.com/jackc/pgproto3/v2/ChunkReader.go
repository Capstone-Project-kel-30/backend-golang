// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ChunkReader is an autogenerated mock type for the ChunkReader type
type ChunkReader struct {
	mock.Mock
}

// Next provides a mock function with given fields: n
func (_m *ChunkReader) Next(n int) ([]byte, error) {
	ret := _m.Called(n)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(int) []byte); ok {
		r0 = rf(n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewChunkReader interface {
	mock.TestingT
	Cleanup(func())
}

// NewChunkReader creates a new instance of ChunkReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewChunkReader(t mockConstructorTestingTNewChunkReader) *ChunkReader {
	mock := &ChunkReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}