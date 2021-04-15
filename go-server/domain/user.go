package domain

import (
	"context"
	"time"
)

type User struct {
	ID        int32     `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Rating    int32     `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	Fetch(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (User, error)
}

type UserUsecase interface {
	Fetch(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (User, error)
}
