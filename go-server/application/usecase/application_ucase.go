package usecase

import (
	"context"
	"time"

	"go-server/domain"
)

type applicationUsecase struct {
	applicationRepo domain.ApplicationRepository
	contextTimeout  time.Duration
}

func NewApplicationUsecase(applicationRepo domain.ApplicationRepository, timeout time.Duration) domain.ApplicationUsecase {
	return &applicationUsecase{
		applicationRepo: applicationRepo,
		contextTimeout:  timeout,
	}
}

func (a *applicationUsecase) Create(c context.Context, roomId int32) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	user := c.Value(domain.CtxUserKey).(*domain.User)

	err = a.applicationRepo.Create(ctx, roomId, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (a *applicationUsecase) FetchAll(c context.Context, roomId int32) (res []domain.Application, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.applicationRepo.FetchAll(ctx, roomId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
