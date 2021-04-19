package domain

import (
	"context"
	"time"
)

type Participation struct {
	UserId    int32     `json:"user_id,omitempty"`
	RoomId    int32     `json:"room_id,omitempty"`
	JoinedAt  time.Time `json:"starting_time,omitempty"`
	LeftAt    time.Time `json:"ending_time,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	IsHost    bool      `json:"is_host,omitempty"`
}

type ParticipationRepository interface {
	Create(ctx context.Context, participation *Participation) (*Participation, error)
}
