package usecase_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go-server/domain"
	"go-server/domain/mocks"
	"go-server/room"
	ucase "go-server/room/usecase"
)

func BoolAddr(b bool) *bool {
	boolVar := b
	return &boolVar
}

func TestCreate(t *testing.T) {
	t.Run("Success, Given valid input, When creating a room, Then will return succeed.", func(t *testing.T) {
		mockRoomRepo := new(mocks.RoomRepository)
		mockInvitationRepo := new(mocks.InvitationRepository)
		mockServiceRepo := new(mocks.ServiceRepository)
		mockParticipationRepo := new(mocks.ParticipationRepository)

		u := ucase.NewRoomUsecase(mockRoomRepo, mockParticipationRepo, mockServiceRepo, mockInvitationRepo, 2*time.Second)

		mockRoomRequest := domain.RoomRequest{
			MaxCount:      4,
			ServiceId:     1,
			PlanName:      "Premium",
			PaymentPeriod: 12,
			IsPublic:      BoolAddr(true),
		}

		mockPlan := domain.Plan{
			PlanName: "Premium",
			Cost:     249,
			MaxCount: 4,
		}

		mockUser := domain.User{
			Id:    1,
			Name:  "Kevin Yu",
			Email: "kevin@ntu.im",
		}

		mockServiceRepo.
			On("GetPlanByKey", mock.Anything, mockRoomRequest.PlanName, fmt.Sprintf("%d", mockRoomRequest.ServiceId)).
			Return(&mockPlan, nil)

		mockRoomRepo.
			On("Create", mock.Anything, mock.MatchedBy(func(room *domain.Room) bool {
				return room.AdminId == mockUser.Id
			})).
			Return(int32(1), nil)

		mockParticipationRepo.
			On("Create", mock.Anything, mock.MatchedBy(func(participation *domain.Participation) bool {
				return participation.IsHost == true && participation.PaymentStatus == "confirmed" &&
					participation.RoomId == int32(1) && participation.UserId == mockUser.Id
			})).
			Return(nil)

		mockCtx := context.WithValue(context.Background(), domain.CtxUserKey, &mockUser)
		err := u.Create(mockCtx, &mockRoomRequest)
		assert.NoError(t, err)
	})

	t.Run("Fail, Given max count exceed, When creating a room, Then will return error.", func(t *testing.T) {
		mockRoomRepo := new(mocks.RoomRepository)
		mockInvitationRepo := new(mocks.InvitationRepository)
		mockServiceRepo := new(mocks.ServiceRepository)
		mockParticipationRepo := new(mocks.ParticipationRepository)

		u := ucase.NewRoomUsecase(mockRoomRepo, mockParticipationRepo, mockServiceRepo, mockInvitationRepo, 2*time.Second)

		mockRoomRequest := domain.RoomRequest{
			MaxCount:      5,
			ServiceId:     1,
			PlanName:      "Premium",
			PaymentPeriod: 12,
			IsPublic:      BoolAddr(true),
		}

		mockPlan := domain.Plan{
			PlanName: "Premium",
			Cost:     249,
			MaxCount: 4,
		}

		mockUser := domain.User{
			Id:    1,
			Name:  "Kevin Yu",
			Email: "kevin@ntu.im",
		}

		mockServiceRepo.
			On("GetPlanByKey", mock.Anything, mockRoomRequest.PlanName, fmt.Sprintf("%d", mockRoomRequest.ServiceId)).
			Return(&mockPlan, nil)

		mockRoomRepo.
			On("Create", mock.Anything, mock.MatchedBy(func(room *domain.Room) bool {
				return room.AdminId == mockUser.Id
			})).
			Return(int32(1), nil)

		mockParticipationRepo.
			On("Create", mock.Anything, mock.MatchedBy(func(participation *domain.Participation) bool {
				return participation.IsHost == true && participation.PaymentStatus == "confirmed" &&
					participation.RoomId == int32(1) && participation.UserId == mockUser.Id
			})).
			Return(nil)

		mockCtx := context.WithValue(context.Background(), domain.CtxUserKey, &mockUser)
		err := u.Create(mockCtx, &mockRoomRequest)
		assert.Error(t, err)
	})
}

func TestGenerateInvitationCode(t *testing.T) {
	t.Run("Success, Given is host, When generating an invitation code, Then will return code.", func(t *testing.T) {
		mockRoomRepo := new(mocks.RoomRepository)
		mockInvitationRepo := new(mocks.InvitationRepository)
		mockServiceRepo := new(mocks.ServiceRepository)
		mockParticipationRepo := new(mocks.ParticipationRepository)

		u := ucase.NewRoomUsecase(mockRoomRepo, mockParticipationRepo, mockServiceRepo, mockInvitationRepo, 2*time.Second)

		var (
			mockRoomId = int32(1)
		)

		mockUser := domain.User{
			Id:    1,
			Name:  "Kevin Yu",
			Email: "kevin@ntu.im",
		}

		mockParticipationRepo.
			On("IsAdmin", mock.Anything, mockRoomId, mockUser.Id).
			Return(true, nil)

		mockInvitationRepo.
			On("GenerateInvitationCode", mock.Anything, mockRoomId, mock.AnythingOfType("string")).
			Return(nil)

		mockCtx := context.WithValue(context.Background(), domain.CtxUserKey, &mockUser)
		_, err := u.GenerateInvitationCode(mockCtx, mockUser.Id)
		assert.NoError(t, err)
	})

	t.Run("Fail, Given is not host, When generating an invitation code, Then will return error.", func(t *testing.T) {
		mockRoomRepo := new(mocks.RoomRepository)
		mockInvitationRepo := new(mocks.InvitationRepository)
		mockServiceRepo := new(mocks.ServiceRepository)
		mockParticipationRepo := new(mocks.ParticipationRepository)

		u := ucase.NewRoomUsecase(mockRoomRepo, mockParticipationRepo, mockServiceRepo, mockInvitationRepo, 2*time.Second)

		var (
			mockRoomId = int32(1)
		)

		mockUser := domain.User{
			Id:    1,
			Name:  "Kevin Yu",
			Email: "kevin@ntu.im",
		}

		mockParticipationRepo.
			On("IsAdmin", mock.Anything, mockRoomId, mockUser.Id).
			Return(false, nil)

		mockInvitationRepo.
			On("GenerateInvitationCode", mock.Anything, mockRoomId, mock.AnythingOfType("string")).
			Return(nil)

		mockCtx := context.WithValue(context.Background(), domain.CtxUserKey, &mockUser)
		_, err := u.GenerateInvitationCode(mockCtx, mockUser.Id)
		assert.Error(t, err)
	})
}

func TestJoinRoom(t *testing.T) {
	t.Run("Success, Given invitation code is valid and user has not join room yet, When entering the code, Then will succcessfully join the room.", func(t *testing.T) {
		mockRoomRepo := new(mocks.RoomRepository)
		mockInvitationRepo := new(mocks.InvitationRepository)
		mockServiceRepo := new(mocks.ServiceRepository)
		mockParticipationRepo := new(mocks.ParticipationRepository)

		u := ucase.NewRoomUsecase(mockRoomRepo, mockParticipationRepo, mockServiceRepo, mockInvitationRepo, 2*time.Second)

		var (
			mockRoomId         = int32(1)
			mockInvitationCode = "asdg212"
		)

		mockUser := domain.User{
			Id:    1,
			Name:  "Kevin Yu",
			Email: "kevin@ntu.im",
		}

		mockInvitationRepo.
			On("ConsumeInvitationCode", mock.Anything, mock.AnythingOfType("string")).
			Return(mockRoomId, nil)

		mockParticipationRepo.
			On("Create", mock.Anything, mock.MatchedBy(func(participation *domain.Participation) bool {
				return participation.UserId == mockUser.Id && participation.RoomId == mockRoomId &&
					participation.PaymentStatus == "unpaid" && participation.IsHost == false
			})).
			Return(nil)

		mockCtx := context.WithValue(context.Background(), domain.CtxUserKey, &mockUser)
		err := u.JoinRoom(mockCtx, mockInvitationCode)
		assert.NoError(t, err)
	})

	t.Run("Fali, Given invitation code is not valid and user has not join room yet, When entering the code, Then will return error.", func(t *testing.T) {
		mockRoomRepo := new(mocks.RoomRepository)
		mockInvitationRepo := new(mocks.InvitationRepository)
		mockServiceRepo := new(mocks.ServiceRepository)
		mockParticipationRepo := new(mocks.ParticipationRepository)

		u := ucase.NewRoomUsecase(mockRoomRepo, mockParticipationRepo, mockServiceRepo, mockInvitationRepo, 2*time.Second)

		var (
			mockRoomId         = int32(1)
			mockInvitationCode = "asdg212"
		)

		mockUser := domain.User{
			Id:    1,
			Name:  "Kevin Yu",
			Email: "kevin@ntu.im",
		}

		mockInvitationRepo.
			On("ConsumeInvitationCode", mock.Anything, mock.AnythingOfType("string")).
			Return(mockRoomId, room.ErrInvalidInvitationCode)

		mockCtx := context.WithValue(context.Background(), domain.CtxUserKey, &mockUser)
		err := u.JoinRoom(mockCtx, mockInvitationCode)
		assert.Error(t, err)
	})

	t.Run("Fali, Given invitation code is valid and user has joined the room, When entering the code, Then will return error.", func(t *testing.T) {
		mockRoomRepo := new(mocks.RoomRepository)
		mockInvitationRepo := new(mocks.InvitationRepository)
		mockServiceRepo := new(mocks.ServiceRepository)
		mockParticipationRepo := new(mocks.ParticipationRepository)

		u := ucase.NewRoomUsecase(mockRoomRepo, mockParticipationRepo, mockServiceRepo, mockInvitationRepo, 2*time.Second)

		var (
			mockRoomId         = int32(1)
			mockInvitationCode = "asdg212"
		)

		mockUser := domain.User{
			Id:    1,
			Name:  "Kevin Yu",
			Email: "kevin@ntu.im",
		}

		mockInvitationRepo.
			On("ConsumeInvitationCode", mock.Anything, mock.AnythingOfType("string")).
			Return(mockRoomId, nil)

		mockInvitationRepo.
			On("ResumeInvitationCode", mock.Anything, mock.AnythingOfType("string")).
			Return(nil)

		mockParticipationRepo.
			On("Create", mock.Anything, mock.MatchedBy(func(participation *domain.Participation) bool {
				return participation.UserId == mockUser.Id && participation.RoomId == mockRoomId &&
					participation.PaymentStatus == "unpaid" && participation.IsHost == false
			})).
			Return(room.ErrAlreadyJoined)

		mockCtx := context.WithValue(context.Background(), domain.CtxUserKey, &mockUser)
		err := u.JoinRoom(mockCtx, mockInvitationCode)
		assert.Error(t, err)
	})
}
