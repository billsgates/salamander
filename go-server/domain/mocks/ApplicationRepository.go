// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "go-server/domain"

	mock "github.com/stretchr/testify/mock"
)

// ApplicationRepository is an autogenerated mock type for the ApplicationRepository type
type ApplicationRepository struct {
	mock.Mock
}

// AcceptApplication provides a mock function with given fields: ctx, roomId, userId
func (_m *ApplicationRepository) AcceptApplication(ctx context.Context, roomId int32, userId int32) error {
	ret := _m.Called(ctx, roomId, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32) error); ok {
		r0 = rf(ctx, roomId, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: ctx, applicationRequest
func (_m *ApplicationRepository) Create(ctx context.Context, applicationRequest *domain.ApplicationRequest) error {
	ret := _m.Called(ctx, applicationRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ApplicationRequest) error); ok {
		r0 = rf(ctx, applicationRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteApplication provides a mock function with given fields: ctx, roomId, userId
func (_m *ApplicationRepository) DeleteApplication(ctx context.Context, roomId int32, userId int32) error {
	ret := _m.Called(ctx, roomId, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32) error); ok {
		r0 = rf(ctx, roomId, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchAll provides a mock function with given fields: ctx, roomId
func (_m *ApplicationRepository) FetchAll(ctx context.Context, roomId int32) ([]domain.Application, error) {
	ret := _m.Called(ctx, roomId)

	var r0 []domain.Application
	if rf, ok := ret.Get(0).(func(context.Context, int32) []domain.Application); ok {
		r0 = rf(ctx, roomId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Application)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, roomId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
