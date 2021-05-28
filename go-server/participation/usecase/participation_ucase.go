package usecase

import (
	"context"
	"time"

	"go-server/domain"
	"go-server/participation"
)

type participationUsecase struct {
	participationRepo domain.ParticipationRepository
	contextTimeout    time.Duration
}

func NewParticipationUsecase(participationRepo domain.ParticipationRepository, timeout time.Duration) domain.ParticipationUsecase {
	return &participationUsecase{
		participationRepo: participationRepo,
		contextTimeout:    timeout,
	}
}

func (p *participationUsecase) IsAdmin(c context.Context, roomId int32) (res bool, err error) {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	user := c.Value(domain.CtxUserKey).(*domain.User)

	isAdmin, err := p.participationRepo.IsAdmin(ctx, roomId, user.Id)
	if !isAdmin || err != nil {
		return false, participation.ErrNotHost
	}
	return true, nil
}

func (p *participationUsecase) IsMember(c context.Context, roomId int32) (res bool, err error) {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	user := c.Value(domain.CtxUserKey).(*domain.User)

	res, err = p.participationRepo.IsMember(ctx, roomId, user.Id)
	return res, err
}
