package domain

import (
	"context"
)

type Application struct {
	ApplicationId int32 `json:"application_id,omitempty"`
	UserId        int32 `json:"user_id,omitempty"`
	RoomId        int32 `json:"room_id,omitempty"`
}

type ApplicationUsecase interface {
	Create(ctx context.Context, roomId int32) error
}

type ApplicationRepository interface {
	Create(ctx context.Context, roomId int32, userId int32) error
}
