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
	ucase "go-server/room/usecase"
)

func BoolAddr(b bool) *bool {
	boolVar := b
	return &boolVar
}

func TestCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRoomRepo := new(mocks.RoomRepository)
		mockInvitationRepo := new(mocks.InvitationRepository)
		mockServiceRepo := new(mocks.ServiceRepository)
		mockParticipationRepo := new(mocks.ParticipationRepository)

		u := ucase.NewRoomUsecase(mockRoomRepo, mockParticipationRepo, mockServiceRepo, mockInvitationRepo, 2*time.Second)

		mockRoomRequest := domain.RoomRequest{
			MaxCount:      4,
			ServiceId:     1,
			AdminId:       1,
			PlanName:      "Premium",
			PaymentPeriod: 12,
			IsPublic:      BoolAddr(true),
		}

		mockRoom := domain.Room{
			ServiceId:     mockRoomRequest.ServiceId,
			PlanName:      mockRoomRequest.PlanName,
			MaxCount:      mockRoomRequest.MaxCount,
			PaymentPeriod: mockRoomRequest.PaymentPeriod,
			AdminId:       mockRoomRequest.AdminId,
			IsPublic:      *mockRoomRequest.IsPublic,
		}

		mockPlan := domain.Plan{
			PlanName: "Premium",
			Cost:     249,
			MaxCount: 4,
		}

		mockParticipation := domain.Participation{
			UserId:        mockRoom.AdminId,
			RoomId:        mockRoom.Id,
			PaymentStatus: "confirmed",
			IsHost:        true,
		}

		mockServiceRepo.
			On("GetPlanByKey", mock.Anything, mockRoomRequest.PlanName, fmt.Sprintf("%d", mockRoomRequest.ServiceId)).
			Return(&mockPlan, nil)

		mockRoomRepo.
			On("Create", mock.Anything, &mockRoom).
			Return(int32(0), nil)

		mockParticipationRepo.
			On("Create", mock.Anything, &mockParticipation).
			Return(nil)

		err := u.Create(context.TODO(), &mockRoomRequest)
		assert.NoError(t, err)
	})
}
