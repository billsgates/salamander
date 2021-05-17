package domain

import (
	"context"
)

type Round struct {
	RoomId          int32  `json:"room_id,omitempty"`
	StartingTime    string `json:"starting_time,omitempty" binding:"required"`
	RoundInterval   int32  `json:"round_interval,omitempty" binding:"required"`
	PaymentDeadline int32  `json:"payment_deadline,omitempty" binding:"required"`
	IsAddCalendar   *bool  `json:"is_add_calendar,omitempty" binding:"required"`
}

type RoundRepository interface {
	AddRound(ctx context.Context, roomId int32, round *Round) error
}
