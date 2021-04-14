package domain

import (
	"context"
	"time"
)

type User struct {
	ID        int64     `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Rating    int64     `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id int64) (User, error)
	Update(ctx context.Context, user *User) error
}

type UserUsecase interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id int64) (User, error)
	Update(ctx context.Context, user *User) error
}
