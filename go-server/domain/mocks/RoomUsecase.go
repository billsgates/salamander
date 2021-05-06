// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "go-server/domain"

	mock "github.com/stretchr/testify/mock"
)

// RoomUsecase is an autogenerated mock type for the RoomUsecase type
type RoomUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, room
func (_m *RoomUsecase) Create(ctx context.Context, room *domain.RoomRequest) error {
	ret := _m.Called(ctx, room)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.RoomRequest) error); ok {
		r0 = rf(ctx, room)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateInvitationCode provides a mock function with given fields: ctx, roomId
func (_m *RoomUsecase) GenerateInvitationCode(ctx context.Context, roomId int32) (string, error) {
	ret := _m.Called(ctx, roomId)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, int32) string); ok {
		r0 = rf(ctx, roomId)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, roomId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJoinedRooms provides a mock function with given fields: ctx
func (_m *RoomUsecase) GetJoinedRooms(ctx context.Context) ([]domain.RoomItem, error) {
	ret := _m.Called(ctx)

	var r0 []domain.RoomItem
	if rf, ok := ret.Get(0).(func(context.Context) []domain.RoomItem); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.RoomItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoomInfo provides a mock function with given fields: ctx, roomId
func (_m *RoomUsecase) GetRoomInfo(ctx context.Context, roomId int32) (*domain.RoomInfoResponse, error) {
	ret := _m.Called(ctx, roomId)

	var r0 *domain.RoomInfoResponse
	if rf, ok := ret.Get(0).(func(context.Context, int32) *domain.RoomInfoResponse); ok {
		r0 = rf(ctx, roomId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.RoomInfoResponse)
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

// JoinRoom provides a mock function with given fields: ctx, code
func (_m *RoomUsecase) JoinRoom(ctx context.Context, code string) error {
	ret := _m.Called(ctx, code)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, code)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LeaveRoom provides a mock function with given fields: ctx, roomId, userId
func (_m *RoomUsecase) LeaveRoom(ctx context.Context, roomId int32, userId int32) error {
	ret := _m.Called(ctx, roomId, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32) error); ok {
		r0 = rf(ctx, roomId, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}