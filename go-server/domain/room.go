package domain

import (
	"context"
	"time"
)

type Room struct {
	Id              int32     `json:"id,omitempty"`
	AccountName     string    `json:"account_name,omitempty"`
	AccountPassword string    `json:"account_password,omitempty"`
	StartingTime    time.Time `json:"starting_time,omitempty"`
	EndingTime      time.Time `json:"ending_time,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	MaxCount        int32     `json:"max_count,omitempty"`
	AdminId         int32     `json:"admin_id,omitempty"`
	ServiceId       int32     `json:"service_id,omitempty"`
	PlanName        string    `json:"plan_name,omitempty"`
}

type RoomCreateRequest struct {
	MaxCount  int32  `json:"max_count" binding:"required"`
	AdminId   int32  `json:"admin_id,omitempty"`
	ServiceId int32  `json:"service_id" binding:"required"`
	PlanName  string `json:"plan_name" binding:"required"`
}

type RoomJoinRequest struct {
	InvitationCode string `json:"invitation_code" binding:"required"`
}

type RoomInfo struct {
	Name     string `json:"name"`
	PlanName string `json:"plan_name"`
	IsHost   bool   `json:"is_host"`
}

type RoomRepository interface {
	Create(ctx context.Context, room *Room) error
}

type RoomUsecase interface {
	Create(ctx context.Context, room *Room) error
	GetJoinedRooms(ctx context.Context, id int32) ([]RoomInfo, error)
	GenerateInvitationCode(ctx context.Context, roomId int32, userId int32) (string, error)
	JoinRoom(ctx context.Context, code string) error
}
