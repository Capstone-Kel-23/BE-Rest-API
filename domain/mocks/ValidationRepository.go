// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/Capstone-Kel-23/BE-Rest-API/domain"
	mock "github.com/stretchr/testify/mock"
)

// ValidationRepository is an autogenerated mock type for the ValidationRepository type
type ValidationRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: validation
func (_m *ValidationRepository) Delete(validation *domain.Validation) error {
	ret := _m.Called(validation)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Validation) error); ok {
		r0 = rf(validation)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByCode provides a mock function with given fields: code
func (_m *ValidationRepository) FindByCode(code string) (*domain.Validation, error) {
	ret := _m.Called(code)

	var r0 *domain.Validation
	if rf, ok := ret.Get(0).(func(string) *domain.Validation); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Validation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByEmailAndType provides a mock function with given fields: email, t
func (_m *ValidationRepository) FindByEmailAndType(email string, t string) (*domain.Validation, error) {
	ret := _m.Called(email, t)

	var r0 *domain.Validation
	if rf, ok := ret.Get(0).(func(string, string) *domain.Validation); ok {
		r0 = rf(email, t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Validation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: validation
func (_m *ValidationRepository) Save(validation *domain.Validation) (*domain.Validation, error) {
	ret := _m.Called(validation)

	var r0 *domain.Validation
	if rf, ok := ret.Get(0).(func(*domain.Validation) *domain.Validation); ok {
		r0 = rf(validation)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Validation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Validation) error); ok {
		r1 = rf(validation)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewValidationRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewValidationRepository creates a new instance of ValidationRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewValidationRepository(t mockConstructorTestingTNewValidationRepository) *ValidationRepository {
	mock := &ValidationRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
