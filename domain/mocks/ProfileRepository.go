// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/Capstone-Kel-23/BE-Rest-API/domain"
	mock "github.com/stretchr/testify/mock"
)

// ProfileRepository is an autogenerated mock type for the ProfileRepository type
type ProfileRepository struct {
	mock.Mock
}

// FindByUserID provides a mock function with given fields: userid
func (_m *ProfileRepository) FindByUserID(userid string) (*domain.Profile, error) {
	ret := _m.Called(userid)

	var r0 *domain.Profile
	if rf, ok := ret.Get(0).(func(string) *domain.Profile); ok {
		r0 = rf(userid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: profile
func (_m *ProfileRepository) Save(profile *domain.Profile) (*domain.Profile, error) {
	ret := _m.Called(profile)

	var r0 *domain.Profile
	if rf, ok := ret.Get(0).(func(*domain.Profile) *domain.Profile); ok {
		r0 = rf(profile)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Profile) error); ok {
		r1 = rf(profile)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateByUserID provides a mock function with given fields: userid, profile
func (_m *ProfileRepository) UpdateByUserID(userid string, profile *domain.Profile) (*domain.Profile, error) {
	ret := _m.Called(userid, profile)

	var r0 *domain.Profile
	if rf, ok := ret.Get(0).(func(string, *domain.Profile) *domain.Profile); ok {
		r0 = rf(userid, profile)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *domain.Profile) error); ok {
		r1 = rf(userid, profile)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProfileRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewProfileRepository creates a new instance of ProfileRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProfileRepository(t mockConstructorTestingTNewProfileRepository) *ProfileRepository {
	mock := &ProfileRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
