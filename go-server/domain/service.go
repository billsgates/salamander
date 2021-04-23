package domain

import (
	"context"
)

type Service struct {
	Id       int32  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	PlanName string `json:"plan_name,omitempty"`
}

type ServiceRepository interface {
	FetchAll(ctx context.Context) ([]Service, error)
	GetDetailByID(ctx context.Context, id string) ([]Service, error)
}

type ServiceUsecase interface {
	FetchAll(ctx context.Context) ([]Service, error)
	GetDetailByID(ctx context.Context, id string) ([]Service, error)
}
