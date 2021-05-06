package domain

import (
	"context"
	"time"
)

type RoomStatus string

const (
	CREATED RoomStatus = "created"
	START   RoomStatus = "start"
	END     RoomStatus = "end"
)

type Room struct {
	Id            int32       `json:"id,omitempty"`
	Announcement  string      `json:"announcement,omitempty"`
	IsPublic      bool        `json:"is_public,omitempty"`
	RoomStatus    *RoomStatus `json:"room_status,omitempty"`
	StartingTime  time.Time   `json:"starting_time,omitempty"`
	EndingTime    time.Time   `json:"ending_time,omitempty"`
	PaymentPeriod int32       `json:"payment_period"`
	CreatedAt     time.Time   `json:"created_at,omitempty"`
	UpdatedAt     time.Time   `json:"updated_at,omitempty"`
	MaxCount      int32       `json:"max_count,omitempty"`
	AdminId       int32       `json:"admin_id,omitempty"`
	ServiceId     int32       `json:"service_id,omitempty"`
	PlanName      string      `json:"plan_name,omitempty"`
}

type RoomRequest struct {
	MaxCount      int32  `json:"max_count" binding:"required"`
	AdminId       int32  `json:"admin_id,omitempty"`
	ServiceId     int32  `json:"service_id" binding:"required"`
	PlanName      string `json:"plan_name" binding:"required"`
	PaymentPeriod int32  `json:"payment_period" binding:"required"`
	IsPublic      *bool  `json:"is_public" binding:"required"`
	Announcement  string `json:"announcement,omitempty"`
}

type RoomJoinRequest struct {
	InvitationCode string `json:"invitation_code" binding:"required"`
}

type RoomItem struct {
	RoomId        int32          `json:"room_id"`
	Name          string         `json:"name"`
	PlanName      string         `json:"plan_name"`
	IsHost        bool           `json:"is_host"`
	PaymentStatus *PaymentStatus `json:"payment_status"`
	RoomStatus    *RoomStatus    `json:"room_status"`
}

type RoomRepository interface {
	Create(ctx context.Context, room *Room) (roomId int32, err error)
}

type RoomUsecase interface {
	Create(ctx context.Context, room *RoomRequest) error
	GetJoinedRooms(ctx context.Context) ([]RoomItem, error)
	GenerateInvitationCode(ctx context.Context, roomId int32) (string, error)
	JoinRoom(ctx context.Context, code string) error
	LeaveRoom(ctx context.Context, roomId int32, userId int32) error
}
