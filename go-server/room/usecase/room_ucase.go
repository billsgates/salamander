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
		ServiceId: roomRequest.ServiceId,
		PlanName:  roomRequest.PlanName,
		MaxCount:  roomRequest.MaxCount,
		AdminId:   roomRequest.AdminId,
		IsPublic:  *roomRequest.IsPublic,
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

func (r *roomUsecase) GenerateInvitationCode(c context.Context, roomId int32) (res string, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	user := c.Value(domain.CtxUserKey).(*domain.User)

	isAdmin, err := r.participationRepo.IsAdmin(ctx, roomId, user.Id)
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

	roomInfo, err := r.participationRepo.GetRoomInfo(ctx, roomId)
	members, err := r.participationRepo.GetRoomMembers(ctx, roomId)

	if len(members) >= int(roomInfo.MaxCount) {
		r.invitationRepo.ResumeInvitationCode(ctx, code)
		return room.ErrRoomFull
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

func (r *roomUsecase) LeaveRoom(c context.Context, roomId int32, userId int32) (err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	user := c.Value(domain.CtxUserKey).(*domain.User)

	isAdmin, err := r.participationRepo.IsAdmin(ctx, roomId, user.Id)
	if !isAdmin || err != nil {
		return room.ErrNotHost
	}

	err = r.participationRepo.LeaveRoom(ctx, roomId, userId)
	if err != nil {
		return err
	}

	return nil
}

func (r *roomUsecase) GetRoomInfo(c context.Context, roomId int32) (res *domain.RoomInfoResponse, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	user := c.Value(domain.CtxUserKey).(*domain.User)

	_, err = r.participationRepo.IsAdmin(ctx, roomId, user.Id)
	if err != nil {
		return nil, room.ErrNotMember
	}

	res, err = r.participationRepo.GetRoomInfo(ctx, roomId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *roomUsecase) GetRoomAdmin(c context.Context, roomId int32) (res *domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	user := c.Value(domain.CtxUserKey).(*domain.User)

	_, err = r.participationRepo.IsAdmin(ctx, roomId, user.Id)
	if err != nil {
		return nil, room.ErrNotMember
	}

	res, err = r.participationRepo.GetRoomAdmin(ctx, roomId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *roomUsecase) GetRoomMembers(c context.Context, roomId int32) (res []domain.Participation, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	user := c.Value(domain.CtxUserKey).(*domain.User)

	_, err = r.participationRepo.IsAdmin(ctx, roomId, user.Id)
	if err != nil {
		return nil, room.ErrNotMember
	}

	res, err = r.participationRepo.GetRoomMembers(ctx, roomId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *roomUsecase) UpdateRoom(c context.Context, roomId int32, roomRequest *domain.RoomRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	user := c.Value(domain.CtxUserKey).(*domain.User)

	_, err = r.participationRepo.IsAdmin(ctx, roomId, user.Id)
	if err != nil {
		return room.ErrNotHost
	}

	plan, err := r.serviceRepo.GetPlanByKey(ctx, roomRequest.PlanName, fmt.Sprintf("%d", roomRequest.ServiceId))
	if err != nil {
		return err
	}

	if plan.MaxCount < roomRequest.MaxCount {
		return room.ErrMaxCountExceed
	}

	err = r.roomRepo.Update(ctx, roomId, &domain.Room{
		ServiceId:     roomRequest.ServiceId,
		PlanName:      roomRequest.PlanName,
		MaxCount:      roomRequest.MaxCount,
		PaymentPeriod: roomRequest.PaymentPeriod,
		IsPublic:      *roomRequest.IsPublic,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *roomUsecase) Delete(c context.Context, roomId int32) (err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	user := c.Value(domain.CtxUserKey).(*domain.User)

	isAdmin, err := r.participationRepo.IsAdmin(ctx, roomId, user.Id)
	if !isAdmin || err != nil {
		return room.ErrNotHost
	}

	err = r.roomRepo.Delete(ctx, roomId)
	if err != nil {
		return err
	}

	return nil
}

func (r *roomUsecase) GetTodayStartingMember(c context.Context) (res []domain.Participation, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	// truncate timestamp to date only
	now := time.Now().Truncate(24 * time.Hour)
	res, err = r.participationRepo.GetRoomMemberByStartingTime(ctx, now)
	if err != nil {
		return nil, err
	}

	return res, nil
}
