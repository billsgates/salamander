package usecase

import (
	"context"
	"time"

	"go-server/domain"
)

type serviceUsecase struct {
	serviceRepo    domain.ServiceRepository
	contextTimeout time.Duration
}

func NewServiceUsecase(serviceRepo domain.ServiceRepository, timeout time.Duration) domain.ServiceUsecase {
	return &serviceUsecase{
		serviceRepo:    serviceRepo,
		contextTimeout: timeout,
	}
}

func (s *serviceUsecase) FetchAll(c context.Context) (res []domain.Service, err error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	res, err = s.serviceRepo.FetchAll(ctx)
	if err != nil {
		return nil, err
	}
	return
}
