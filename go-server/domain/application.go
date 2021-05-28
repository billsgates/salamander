package domain

import (
	"context"
)

type Application struct {
	UserId          int32  `json:"user_id,omitempty"`
	UserName        string `json:"user_name,omitempty"`
	UserRating      int32  `json:"user_rating,omitempty"`
	ApplicationDate string `json:"application_date,omitempty"`
	RoomId          int32  `json:"room_id,omitempty"`
}

type ApplicationRequest struct {
	UserId int32 `json:"user_id,omitempty"`
	RoomId int32 `json:"room_id,omitempty"`
}

type ApplicationUsecase interface {
	Create(ctx context.Context, roomId int32) error
	FetchAll(ctx context.Context, roomId int32) (res []Application, err error)
	AcceptApplication(ctx context.Context, roomId int32, userId int32) (err error)
}

type ApplicationRepository interface {
	Create(ctx context.Context, roomId int32, userId int32) error
	FetchAll(ctx context.Context, roomId int32) (res []Application, err error)
	AcceptApplication(ctx context.Context, roomId int32, userId int32) (err error)
}
