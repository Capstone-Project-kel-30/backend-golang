// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/mashbens/cps/business/admin/entity"
	mock "github.com/stretchr/testify/mock"
)

// AdminRepo is an autogenerated mock type for the AdminRepo type
type AdminRepo struct {
	mock.Mock
}

// DeleteAdmin provides a mock function with given fields: adminID
func (_m *AdminRepo) DeleteAdmin(adminID string) error {
	ret := _m.Called(adminID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(adminID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAdminByEmail provides a mock function with given fields: email
func (_m *AdminRepo) FindAdminByEmail(email string) (entity.Admin, error) {
	ret := _m.Called(email)

	var r0 entity.Admin
	if rf, ok := ret.Get(0).(func(string) entity.Admin); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(entity.Admin)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAdminByID provides a mock function with given fields: id
func (_m *AdminRepo) FindAdminByID(id string) (entity.Admin, error) {
	ret := _m.Called(id)

	var r0 entity.Admin
	if rf, ok := ret.Get(0).(func(string) entity.Admin); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.Admin)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllAdmins provides a mock function with given fields: search
func (_m *AdminRepo) FindAllAdmins(search string) []entity.Admin {
	ret := _m.Called(search)

	var r0 []entity.Admin
	if rf, ok := ret.Get(0).(func(string) []entity.Admin); ok {
		r0 = rf(search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Admin)
		}
	}

	return r0
}

// InsertAdmin provides a mock function with given fields: _a0
func (_m *AdminRepo) InsertAdmin(_a0 entity.Admin) (entity.Admin, error) {
	ret := _m.Called(_a0)

	var r0 entity.Admin
	if rf, ok := ret.Get(0).(func(entity.Admin) entity.Admin); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(entity.Admin)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Admin) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAdmin provides a mock function with given fields: _a0
func (_m *AdminRepo) UpdateAdmin(_a0 entity.Admin) (entity.Admin, error) {
	ret := _m.Called(_a0)

	var r0 entity.Admin
	if rf, ok := ret.Get(0).(func(entity.Admin) entity.Admin); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(entity.Admin)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Admin) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAdminRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewAdminRepo creates a new instance of AdminRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAdminRepo(t mockConstructorTestingTNewAdminRepo) *AdminRepo {
	mock := &AdminRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
