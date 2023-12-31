// Code generated by mockery v2.30.1. DO NOT EDIT.

package write_test

import (
	command "github.com/rcgc/go-hexagonal.git/internal/class/application/command"
	dto "github.com/rcgc/go-hexagonal.git/internal/class/infrastructure/repository/json/dto"

	mock "github.com/stretchr/testify/mock"
)

// MapperMock is an autogenerated mock type for the Mapper type
type MapperMock struct {
	mock.Mock
}

// CommandToDTOClass provides a mock function with given fields: cmd
func (_m *MapperMock) CommandToDTOClass(cmd command.Update) dto.ClassStudent {
	ret := _m.Called(cmd)

	var r0 dto.ClassStudent
	if rf, ok := ret.Get(0).(func(command.Update) dto.ClassStudent); ok {
		r0 = rf(cmd)
	} else {
		r0 = ret.Get(0).(dto.ClassStudent)
	}

	return r0
}

// NewMapperMock creates a new instance of MapperMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMapperMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *MapperMock {
	mock := &MapperMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}