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
	ServiceName   string         `json:"service_name"`
	PlanName      string         `json:"plan_name"`
	IsHost        bool           `json:"is_host"`
	PaymentStatus *PaymentStatus `json:"payment_status"`
	RoomStatus    *RoomStatus    `json:"room_status"`
	Cost          int32          `json:"cost"`
}

type RoomInfoResponse struct {
	RoomId        int32           `json:"room_id,omitempty"`
	IsPublic      bool            `json:"is_public"`
	Announcement  string          `json:"announcement"`
	MaxCount      int32           `json:"max_count,omitempty"`
	PaymentPeriod int32           `json:"payment_period,omitempty"`
	RoomStatus    *RoomStatus     `json:"room_status,omitempty"`
	StartingTime  *time.Time      `json:"starting_time"`
	EndingTime    *time.Time      `json:"ending_time"`
	ServiceName   string          `json:"service_name,omitempty"`
	PlanName      string          `json:"plan_name,omitempty"`
	Role          string          `json:"role,omitempty"`
	PaymentFee    int32           `json:"payment_fee,omitempty"`
	Admin         *User           `json:"admin,omitempty" gorm:"-"`
	Members       []Participation `json:"members,omitempty" gorm:"-"`
}

type RoomRepository interface {
	Create(ctx context.Context, room *Room) (roomId int32, err error)
	Update(ctx context.Context, roomId int32, room *Room) error
	Delete(ctx context.Context, roomId int32) (err error)
}

type RoomUsecase interface {
	Create(ctx context.Context, room *RoomRequest) error
	Delete(ctx context.Context, roomId int32) error
	GetRoomInfo(ctx context.Context, roomId int32) (res *RoomInfoResponse, err error)
	GetRoomAdmin(ctx context.Context, roomId int32) (res *User, err error)
	GetRoomMembers(ctx context.Context, roomId int32) (res []Participation, err error)
	GetJoinedRooms(ctx context.Context) ([]RoomItem, error)
	GenerateInvitationCode(ctx context.Context, roomId int32) (string, error)
	JoinRoom(ctx context.Context, code string) error
	LeaveRoom(ctx context.Context, roomId int32, userId int32) error
	UpdateRoom(ctx context.Context, roomId int32, room *RoomRequest) error
}
