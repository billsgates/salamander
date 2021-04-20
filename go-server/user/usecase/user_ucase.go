package usecase

import (
	"context"
	"time"

	"go-server/domain"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepo domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       userRepo,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) Create(c context.Context, user *domain.User) (res *domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *userUsecase) FetchAll(c context.Context) (res []domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.userRepo.FetchAll(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *userUsecase) GetByID(c context.Context, id string) (res *domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *userUsecase) GetByEmailPassword(c context.Context, email string, password string) (res *domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.userRepo.GetByEmailPassword(ctx, email, password)
	if err != nil {
		return nil, err
	}
	return res, nil
}
