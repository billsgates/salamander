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

func (u *userUsecase) Create(c context.Context, user *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user.CreatedAt = time.Now()
	return u.userRepo.Create(ctx, user)
}

func (u *userUsecase) GetByID(c context.Context, id string) (res domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.userRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	return res, nil
}

func (u *userUsecase) Update(c context.Context, user *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user.UpdatedAt = time.Now()
	return u.userRepo.Update(ctx, user)
}
