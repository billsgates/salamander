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
	UserId        int32     `json:"user_id,omitempty"`
	UserName      string    `json:"user_name,omitempty"`
	RoomId        int32     `json:"room_id,omitempty"`
	PaymentStatus string    `json:"payment_status,omitempty"`
	JoinedAt      time.Time `json:"starting_time,omitempty"`
	LeftAt        time.Time `json:"ending_time,omitempty"`
	IsHost        bool      `json:"is_host,omitempty"`
}

type ParticipationRequest struct {
	UserId int32 `json:"user_id,omitempty" binding:"required"`
	RoomId int32 `json:"room_id,omitempty" binding:"required"`
}

type ParticipationRepository interface {
	Create(ctx context.Context, participation *Participation) error
	GetRoomInfo(ctx context.Context, roomId int32) (res *RoomInfoResponse, err error)
	GetRoomMembers(ctx context.Context, roomId int32) (res []Participation, err error)
	GetJoinedRooms(ctx context.Context, userId int32) ([]RoomItem, error)
	IsAdmin(ctx context.Context, roomId int32, userId int32) (bool, error)
	LeaveRoom(ctx context.Context, roomId int32, userId int32) error
}
