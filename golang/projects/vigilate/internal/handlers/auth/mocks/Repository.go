// Code generated by mockery 2.9.2. DO NOT EDIT.

package mocks

import (
	models "github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Authenticate provides a mock function with given fields: email, testPassword
func (_m *Repository) Authenticate(email string, testPassword string) (int, string, error) {
	ret := _m.Called(email, testPassword)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string) int); ok {
		r0 = rf(email, testPassword)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(email, testPassword)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(email, testPassword)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeleteToken provides a mock function with given fields: token
func (_m *Repository) DeleteToken(token string) error {
	ret := _m.Called(token)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserById provides a mock function with given fields: id
func (_m *Repository) GetUserById(id int) (models.User, error) {
	ret := _m.Called(id)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(int) models.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertRememberMeToken provides a mock function with given fields: id, token
func (_m *Repository) InsertRememberMeToken(id int, token string) error {
	ret := _m.Called(id, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(id, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
