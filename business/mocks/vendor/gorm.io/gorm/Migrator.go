// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	gorm "gorm.io/gorm"
	clause "gorm.io/gorm/clause"

	mock "github.com/stretchr/testify/mock"

	schema "gorm.io/gorm/schema"
)

// Migrator is an autogenerated mock type for the Migrator type
type Migrator struct {
	mock.Mock
}

// AddColumn provides a mock function with given fields: dst, field
func (_m *Migrator) AddColumn(dst interface{}, field string) error {
	ret := _m.Called(dst, field)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string) error); ok {
		r0 = rf(dst, field)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AlterColumn provides a mock function with given fields: dst, field
func (_m *Migrator) AlterColumn(dst interface{}, field string) error {
	ret := _m.Called(dst, field)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string) error); ok {
		r0 = rf(dst, field)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AutoMigrate provides a mock function with given fields: dst
func (_m *Migrator) AutoMigrate(dst ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, dst...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(dst...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ColumnTypes provides a mock function with given fields: dst
func (_m *Migrator) ColumnTypes(dst interface{}) ([]gorm.ColumnType, error) {
	ret := _m.Called(dst)

	var r0 []gorm.ColumnType
	if rf, ok := ret.Get(0).(func(interface{}) []gorm.ColumnType); ok {
		r0 = rf(dst)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gorm.ColumnType)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(dst)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateConstraint provides a mock function with given fields: dst, name
func (_m *Migrator) CreateConstraint(dst interface{}, name string) error {
	ret := _m.Called(dst, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string) error); ok {
		r0 = rf(dst, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateIndex provides a mock function with given fields: dst, name
func (_m *Migrator) CreateIndex(dst interface{}, name string) error {
	ret := _m.Called(dst, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string) error); ok {
		r0 = rf(dst, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTable provides a mock function with given fields: dst
func (_m *Migrator) CreateTable(dst ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, dst...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(dst...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateView provides a mock function with given fields: name, option
func (_m *Migrator) CreateView(name string, option gorm.ViewOption) error {
	ret := _m.Called(name, option)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, gorm.ViewOption) error); ok {
		r0 = rf(name, option)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CurrentDatabase provides a mock function with given fields:
func (_m *Migrator) CurrentDatabase() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// DropColumn provides a mock function with given fields: dst, field
func (_m *Migrator) DropColumn(dst interface{}, field string) error {
	ret := _m.Called(dst, field)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string) error); ok {
		r0 = rf(dst, field)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DropConstraint provides a mock function with given fields: dst, name
func (_m *Migrator) DropConstraint(dst interface{}, name string) error {
	ret := _m.Called(dst, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string) error); ok {
		r0 = rf(dst, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DropIndex provides a mock function with given fields: dst, name
func (_m *Migrator) DropIndex(dst interface{}, name string) error {
	ret := _m.Called(dst, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string) error); ok {
		r0 = rf(dst, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DropTable provides a mock function with given fields: dst
func (_m *Migrator) DropTable(dst ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, dst...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(dst...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DropView provides a mock function with given fields: name
func (_m *Migrator) DropView(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FullDataTypeOf provides a mock function with given fields: _a0
func (_m *Migrator) FullDataTypeOf(_a0 *schema.Field) clause.Expr {
	ret := _m.Called(_a0)

	var r0 clause.Expr
	if rf, ok := ret.Get(0).(func(*schema.Field) clause.Expr); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(clause.Expr)
	}

	return r0
}

// GetIndexes provides a mock function with given fields: dst
func (_m *Migrator) GetIndexes(dst interface{}) ([]gorm.Index, error) {
	ret := _m.Called(dst)

	var r0 []gorm.Index
	if rf, ok := ret.Get(0).(func(interface{}) []gorm.Index); ok {
		r0 = rf(dst)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gorm.Index)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(dst)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTables provides a mock function with given fields:
func (_m *Migrator) GetTables() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasColumn provides a mock function with given fields: dst, field
func (_m *Migrator) HasColumn(dst interface{}, field string) bool {
	ret := _m.Called(dst, field)

	var r0 bool
	if rf, ok := ret.Get(0).(func(interface{}, string) bool); ok {
		r0 = rf(dst, field)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// HasConstraint provides a mock function with given fields: dst, name
func (_m *Migrator) HasConstraint(dst interface{}, name string) bool {
	ret := _m.Called(dst, name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(interface{}, string) bool); ok {
		r0 = rf(dst, name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// HasIndex provides a mock function with given fields: dst, name
func (_m *Migrator) HasIndex(dst interface{}, name string) bool {
	ret := _m.Called(dst, name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(interface{}, string) bool); ok {
		r0 = rf(dst, name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// HasTable provides a mock function with given fields: dst
func (_m *Migrator) HasTable(dst interface{}) bool {
	ret := _m.Called(dst)

	var r0 bool
	if rf, ok := ret.Get(0).(func(interface{}) bool); ok {
		r0 = rf(dst)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MigrateColumn provides a mock function with given fields: dst, field, columnType
func (_m *Migrator) MigrateColumn(dst interface{}, field *schema.Field, columnType gorm.ColumnType) error {
	ret := _m.Called(dst, field, columnType)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, *schema.Field, gorm.ColumnType) error); ok {
		r0 = rf(dst, field, columnType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RenameColumn provides a mock function with given fields: dst, oldName, field
func (_m *Migrator) RenameColumn(dst interface{}, oldName string, field string) error {
	ret := _m.Called(dst, oldName, field)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, string) error); ok {
		r0 = rf(dst, oldName, field)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RenameIndex provides a mock function with given fields: dst, oldName, newName
func (_m *Migrator) RenameIndex(dst interface{}, oldName string, newName string) error {
	ret := _m.Called(dst, oldName, newName)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, string) error); ok {
		r0 = rf(dst, oldName, newName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RenameTable provides a mock function with given fields: oldName, newName
func (_m *Migrator) RenameTable(oldName interface{}, newName interface{}) error {
	ret := _m.Called(oldName, newName)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, interface{}) error); ok {
		r0 = rf(oldName, newName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMigrator interface {
	mock.TestingT
	Cleanup(func())
}

// NewMigrator creates a new instance of Migrator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMigrator(t mockConstructorTestingTNewMigrator) *Migrator {
	mock := &Migrator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}