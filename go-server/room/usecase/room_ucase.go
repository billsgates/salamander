package usecase

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"go-server/domain"
)

type roomUsecase struct {
	roomRepo          domain.RoomRepository
	participationRepo domain.ParticipationRepository
	serviceRepo       domain.ServiceRepository
	invitationRepo    domain.InvitationRepository
	contextTimeout    time.Duration
}

func NewRoomUsecase(roomRepo domain.RoomRepository, participationRepo domain.ParticipationRepository, serviceRepo domain.ServiceRepository, invitationRepo domain.InvitationRepository, timeout time.Duration) domain.RoomUsecase {
	return &roomUsecase{
		roomRepo:          roomRepo,
		participationRepo: participationRepo,
		serviceRepo:       serviceRepo,
		invitationRepo:    invitationRepo,
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

func (r *roomUsecase) GenerateInvitationCode(c context.Context, roomId int32) (res string, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	user := c.Value(domain.CtxUserKey).(*domain.User)
	_, err = r.participationRepo.IsAdmin(ctx, roomId, user.Id)
	if err != nil {
		return "", err
	}

	code := sha1.New()
	code.Write([]byte(time.Now().String()))
	code.Write([]byte(fmt.Sprint(roomId)))
	invitationCode := fmt.Sprintf("%x", code.Sum(nil))[0:7]

	err = r.invitationRepo.GenerateInvitationCode(ctx, roomId, invitationCode)
	if err != nil {
		return "", err
	}

	return invitationCode, nil
}

func (r *roomUsecase) JoinRoom(c context.Context, code string) (err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	roomId, err := r.invitationRepo.ConsumeInvitationCode(ctx, code)
	if err != nil {
		return err
	}

	user := c.Value(domain.CtxUserKey).(*domain.User)

	err = r.participationRepo.Create(ctx, &domain.Participation{
		UserId: user.Id,
		RoomId: roomId,
		IsHost: false,
	})
	if err != nil {
		r.invitationRepo.ResumeInvitationCode(ctx, code)
		return err
	}

	return nil
}
