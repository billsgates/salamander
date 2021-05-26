package domain

import (
	"context"
)

type User struct {
	Id             int32   `json:"id,omitempty"`
	Name           string  `json:"name,omitempty"`
	Email          string  `json:"email,omitempty"`
	Phone          string  `json:"phone"`
	PasswordDigest string  `json:"password_digest,omitempty"`
	Rating         float32 `json:"rating"`
}

type UserInfo struct {
	Id          int32   `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Email       string  `json:"email,omitempty"`
	Phone       string  `json:"phone"`
	Rating      float32 `json:"rating"`
	RatingCount int32   `json:"rating_count"`
}

type UserRequest struct {
	Id       int32  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
}

type RatingRequest struct {
	Rating int32 `json:"rating,omitempty"`
}

type RatingResponse struct {
	Rating      float32 `json:"rating,omitempty"`
	RatingCount int32   `json:"rating_count"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FetchAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (*UserInfo, error)
	GetByEmailPassword(ctx context.Context, email string, password string) (*User, error)
	Update(ctx context.Context, user *UserRequest) error
	UpdateRating(ctx context.Context, id string, rating float32) error
}

type UserUsecase interface {
	Create(ctx context.Context, user *User) error
	FetchAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (*UserInfo, error)
	GetByEmailPassword(ctx context.Context, email string, password string) (*User, error)
	GetUserRating(ctx context.Context, id string) (*RatingResponse, error)
	Update(ctx context.Context, user *UserRequest) error
	UpdateRating(ctx context.Context, id string, rating int32) error
}
