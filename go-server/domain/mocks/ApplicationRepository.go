// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ApplicationRepository is an autogenerated mock type for the ApplicationRepository type
type ApplicationRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, roomId, userId
func (_m *ApplicationRepository) Create(ctx context.Context, roomId int32, userId int32) error {
	ret := _m.Called(ctx, roomId, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32) error); ok {
		r0 = rf(ctx, roomId, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
