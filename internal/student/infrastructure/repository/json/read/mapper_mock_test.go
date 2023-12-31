// Code generated by mockery v2.30.1. DO NOT EDIT.

package read_test

import (
	domain "github.com/rcgc/go-hexagonal.git/internal/student/domain"
	dto "github.com/rcgc/go-hexagonal.git/internal/student/infrastructure/repository/json/dto"
	mock "github.com/stretchr/testify/mock"
)

// MapperMock is an autogenerated mock type for the Mapper type
type MapperMock struct {
	mock.Mock
}

// DTOClassesToDomain provides a mock function with given fields: classes
func (_m *MapperMock) DTOClassesToDomain(classes []dto.Class) []domain.Class {
	ret := _m.Called(classes)

	var r0 []domain.Class
	if rf, ok := ret.Get(0).(func([]dto.Class) []domain.Class); ok {
		r0 = rf(classes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Class)
		}
	}

	return r0
}

// DTOProfileToDomain provides a mock function with given fields: email, profile
func (_m *MapperMock) DTOProfileToDomain(email string, profile dto.Profile) domain.Profile {
	ret := _m.Called(email, profile)

	var r0 domain.Profile
	if rf, ok := ret.Get(0).(func(string, dto.Profile) domain.Profile); ok {
		r0 = rf(email, profile)
	} else {
		r0 = ret.Get(0).(domain.Profile)
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