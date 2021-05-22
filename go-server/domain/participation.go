package domain

import (
	"context"
	"time"
)

type PaymentStatus string

const (
	UNPAID    PaymentStatus = "unpaid"
	PENDING   PaymentStatus = "pending"
	CONFIRMED PaymentStatus = "confirmed"
)

type Participation struct {
	UserId        int32  `json:"user_id,omitempty"`
	UserName      string `json:"user_name,omitempty"`
	RoomId        int32  `json:"room_id,omitempty"`
	PaymentStatus string `json:"payment_status,omitempty"`
	IsHost        bool   `json:"is_host,omitempty"`
}

type ParticipationRequest struct {
	UserId int32 `json:"user_id,omitempty" binding:"required"`
	RoomId int32 `json:"room_id,omitempty" binding:"required"`
}

type ParticipationUsecase interface {
	IsMember(ctx context.Context, roomId int32) (bool, error)
	IsAdmin(ctx context.Context, roomId int32) (bool, error)
}

type ParticipationRepository interface {
	Create(ctx context.Context, participation *Participation) error
	GetRoomInfo(ctx context.Context, roomId int32) (res *RoomInfoResponse, err error)
	GetRoomAdmin(ctx context.Context, roomId int32) (res *User, err error)
	GetRoomMembers(ctx context.Context, roomId int32) (res []Participation, err error)
	GetJoinedRooms(ctx context.Context, userId int32) ([]RoomItem, error)
	GetRoomMemberByStartingTime(ctx context.Context, starting_time time.Time) (res []Participation, err error)
	IsAdmin(ctx context.Context, roomId int32, userId int32) (bool, error)
	IsMember(ctx context.Context, roomId int32, userId int32) (bool, error)
	LeaveRoom(ctx context.Context, roomId int32, userId int32) error
}
