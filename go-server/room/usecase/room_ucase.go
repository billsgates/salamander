package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go-server/domain"
)

type roomUsecase struct {
	roomRepo          domain.RoomRepository
	participationRepo domain.ParticipationRepository
	serviceRepo       domain.ServiceRepository
	contextTimeout    time.Duration
}

func NewRoomUsecase(roomRepo domain.RoomRepository, participationRepo domain.ParticipationRepository, serviceRepo domain.ServiceRepository, timeout time.Duration) domain.RoomUsecase {
	return &roomUsecase{
		roomRepo:          roomRepo,
		participationRepo: participationRepo,
		serviceRepo:       serviceRepo,
		contextTimeout:    timeout,
	}
}

func (r *roomUsecase) Create(c context.Context, room *domain.Room) (err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	plan, err := r.serviceRepo.GetPlanByKey(ctx, room.PlanName, fmt.Sprintf("%d", room.ServiceId))
	if err != nil {
		return err
	}

	if plan.MaxCount < room.MaxCount {
		return errors.New("max count exceed")
	}

	err = r.roomRepo.Create(ctx, room)
	if err != nil {
		return err
	}

	err = r.participationRepo.Create(ctx, &domain.Participation{
		UserId: room.AdminId,
		RoomId: room.Id,
		IsHost: true,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *roomUsecase) GetJoinedRooms(c context.Context, id int32) (res []domain.RoomInfo, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	res, err = r.participationRepo.GetJoinedRooms(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
