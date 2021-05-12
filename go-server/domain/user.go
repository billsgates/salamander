package domain

import (
	"context"
	"time"
)

type User struct {
	Id             int32     `json:"id,omitempty"`
	Name           string    `json:"name,omitempty"`
	Email          string    `json:"email,omitempty"`
	PasswordDigest string    `json:"password_digest,omitempty"`
	Rating         int32     `json:"rating,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

type UserRequest struct {
	Id       int32  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FetchAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	GetByEmailPassword(ctx context.Context, email string, password string) (*User, error)
	Update(ctx context.Context, user *UserRequest) error
}

type UserUsecase interface {
	Create(ctx context.Context, user *User) error
	FetchAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	GetByEmailPassword(ctx context.Context, email string, password string) (*User, error)
	Update(ctx context.Context, user *UserRequest) error
}
