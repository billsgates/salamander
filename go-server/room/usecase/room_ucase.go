package usecase

import (
	"context"
	"time"

	"go-server/domain"
)

type roomUsecase struct {
	roomRepo          domain.RoomRepository
	participationRepo domain.ParticipationRepository
	contextTimeout    time.Duration
}

func NewRoomUsecase(roomRepo domain.RoomRepository, participationRepo domain.ParticipationRepository, timeout time.Duration) domain.RoomUsecase {
	return &roomUsecase{
		roomRepo:          roomRepo,
		participationRepo: participationRepo,
		contextTimeout:    timeout,
	}
}

func (r *roomUsecase) Create(c context.Context, room *domain.Room) (res *domain.Room, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	res, err = r.roomRepo.Create(ctx, room)
	if err != nil {
		return
	}

	_, err = r.participationRepo.Create(ctx, &domain.Participation{
		UserId: room.AdminId,
		RoomId: room.Id,
		IsHost: true,
	})
	if err != nil {
		return
	}

	return res, nil
}
