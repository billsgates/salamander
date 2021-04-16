package domain

import (
	"context"
	"time"
)

type User struct {
	Id        int32     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Rating    int32     `json:"rating,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	FetchAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (*User, error)
}

type UserUsecase interface {
	Create(ctx context.Context, user *User) (*User, error)
	FetchAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (*User, error)
}
