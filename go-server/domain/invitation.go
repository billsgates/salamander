package domain

import (
	"context"
)

type Invitation struct {
	RoomId         int32  `json:"room_id,omitempty"`
	InvitationCode string `json:"invitation_code,omitempty"`
	IsValid        bool   `json:"is_valid,omitempty"`
}

type InvitationRepository interface {
	GenerateInvitationCode(ctx context.Context, roomId int32, code string) error
	ConsumeInvitationCode(ctx context.Context, code string) (roomId int32, err error)
}
