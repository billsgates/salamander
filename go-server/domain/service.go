package domain

import (
	"context"
	"time"
)

type Service struct {
	Id        int32     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type ServiceRepository interface {
	FetchAll(ctx context.Context) ([]Service, error)
}

type ServiceUsecase interface {
	FetchAll(ctx context.Context) ([]Service, error)
}
