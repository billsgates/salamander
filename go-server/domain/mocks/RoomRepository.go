// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "go-server/domain"

	mock "github.com/stretchr/testify/mock"
)

// RoomRepository is an autogenerated mock type for the RoomRepository type
type RoomRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, room
func (_m *RoomRepository) Create(ctx context.Context, room *domain.Room) (int32, error) {
	ret := _m.Called(ctx, room)

	var r0 int32
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Room) int32); ok {
		r0 = rf(ctx, room)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Room) error); ok {
		r1 = rf(ctx, room)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, roomId
func (_m *RoomRepository) Delete(ctx context.Context, roomId int32) error {
	ret := _m.Called(ctx, roomId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(ctx, roomId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPublicRooms provides a mock function with given fields: ctx
func (_m *RoomRepository) GetPublicRooms(ctx context.Context) ([]domain.RoomPublic, error) {
	ret := _m.Called(ctx)

	var r0 []domain.RoomPublic
	if rf, ok := ret.Get(0).(func(context.Context) []domain.RoomPublic); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.RoomPublic)
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

// Update provides a mock function with given fields: ctx, roomId, room
func (_m *RoomRepository) Update(ctx context.Context, roomId int32, room *domain.Room) error {
	ret := _m.Called(ctx, roomId, room)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, *domain.Room) error); ok {
		r0 = rf(ctx, roomId, room)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRoundId provides a mock function with given fields: ctx, roomId, roundId
func (_m *RoomRepository) UpdateRoundId(ctx context.Context, roomId int32, roundId int32) error {
	ret := _m.Called(ctx, roomId, roundId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32) error); ok {
		r0 = rf(ctx, roomId, roundId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
