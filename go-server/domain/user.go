package domain

import (
	"context"
)

type User struct {
	Id             int32  `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Email          string `json:"email,omitempty"`
	Phone          string `json:"phone"`
	PasswordDigest string `json:"password_digest,omitempty"`
	Rating         int32  `json:"rating"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FetchAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	GetByEmailPassword(ctx context.Context, email string, password string) (*User, error)
}

type UserUsecase interface {
	Create(ctx context.Context, user *User) error
	FetchAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	GetByEmailPassword(ctx context.Context, email string, password string) (*User, error)
}
