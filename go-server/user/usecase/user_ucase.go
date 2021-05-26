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

	err = u.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
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

func (u *userUsecase) GetByID(c context.Context, id string) (res *domain.UserInfo, err error) {
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

func (u *userUsecase) GetUserRating(c context.Context, id string) (res *domain.RatingResponse, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	userRes, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	res = &domain.RatingResponse{
		Rating:      userRes.Rating,
		RatingCount: userRes.RatingCount,
	}
	return res, nil
}

func (u *userUsecase) Update(c context.Context, user *domain.UserRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	err = u.userRepo.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) UpdateRating(c context.Context, id string, rating int32) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	// get original rating to calculate new avg rating
	userRes, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	newRating := (userRes.Rating*float32(userRes.RatingCount) + float32(rating)) / float32(userRes.RatingCount+1)

	err = u.userRepo.UpdateRating(ctx, id, newRating)
	if err != nil {
		return err
	}
	return nil
}
