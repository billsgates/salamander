package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"go-server/domain"
	"go-server/room"
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

func (r *roomUsecase) Create(c context.Context, roomRequest *domain.RoomRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	user := c.Value(domain.CtxUserKey).(*domain.User)

	roomRequest.AdminId = user.Id

	plan, err := r.serviceRepo.GetPlanByKey(ctx, roomRequest.PlanName, fmt.Sprintf("%d", roomRequest.ServiceId))
	if err != nil {
		return err
	}

	if plan.MaxCount < roomRequest.MaxCount {
		return room.ErrMaxCountExceed
	}

	roomId, err := r.roomRepo.Create(ctx, &domain.Room{
		ServiceId:     roomRequest.ServiceId,
		PlanName:      roomRequest.PlanName,
		MaxCount:      roomRequest.MaxCount,
		PaymentPeriod: roomRequest.PaymentPeriod,
		AdminId:       roomRequest.AdminId,
		IsPublic:      *roomRequest.IsPublic,
	})
	if err != nil {
		return err
	}

	err = r.participationRepo.Create(ctx, &domain.Participation{
		UserId:        user.Id,
		RoomId:        roomId,
		PaymentStatus: "confirmed",
		IsHost:        true,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *roomUsecase) GetJoinedRooms(c context.Context) (res []domain.RoomItem, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	user := c.Value(domain.CtxUserKey).(*domain.User)

	res, err = r.participationRepo.GetJoinedRooms(ctx, user.Id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *roomUsecase) GenerateInvitationCode(c context.Context, roomId int32, userId int32) (res string, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	isAdmin, err := r.participationRepo.IsAdmin(ctx, roomId, userId)
	if !isAdmin || err != nil {
		return "", room.ErrNotHost
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
		return room.ErrInvalidInvitationCode
	}

	user := c.Value(domain.CtxUserKey).(*domain.User)

	err = r.participationRepo.Create(ctx, &domain.Participation{
		UserId:        user.Id,
		RoomId:        roomId,
		PaymentStatus: "unpaid",
		IsHost:        false,
	})
	if err != nil {
		r.invitationRepo.ResumeInvitationCode(ctx, code)
		return room.ErrAlreadyJoined
	}

	return nil
}
