package domain

import (
	"context"
)

type Service struct {
	Id   int32  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ServiceDetail struct {
	Id       int32  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	PlanName string `json:"plan_name,omitempty"`
	Cost     int32  `json:"cost,omitempty"`
	MaxCount int32  `json:"max_count,omitempty"`
}

type ServiceRepository interface {
	FetchAll(ctx context.Context) ([]Service, error)
	GetDetailByID(ctx context.Context, id string) ([]ServiceDetail, error)
}

type ServiceUsecase interface {
	FetchAll(ctx context.Context) ([]Service, error)
	GetDetailByID(ctx context.Context, id string) ([]ServiceDetail, error)
}
