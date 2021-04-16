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

type RoomRepository interface {
	Create(ctx context.Context, room *Room) (*Room, error)
}

type RoomUsecase interface {
	Create(ctx context.Context, room *Room) (*Room, error)
}
