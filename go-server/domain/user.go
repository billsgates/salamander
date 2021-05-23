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

type UserInfo struct {
	Id     int32  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Phone  string `json:"phone"`
	Rating int32  `json:"rating"`
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
	GetByID(ctx context.Context, id string) (*UserInfo, error)
	GetByEmailPassword(ctx context.Context, email string, password string) (*User, error)
	Update(ctx context.Context, user *UserRequest) error
}

type UserUsecase interface {
	Create(ctx context.Context, user *User) error
	FetchAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (*UserInfo, error)
	GetByEmailPassword(ctx context.Context, email string, password string) (*User, error)
	Update(ctx context.Context, user *UserRequest) error
}
