// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	usecase "github.com/vivaldy22/go-unit-test/6-mock/usecase"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: id
func (_m *UserService) GetByID(id int64) (usecase.User, error) {
	ret := _m.Called(id)

	var r0 usecase.User
	if rf, ok := ret.Get(0).(func(int64) usecase.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(usecase.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
