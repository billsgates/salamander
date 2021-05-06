// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "go-server/domain"

	mock "github.com/stretchr/testify/mock"
)

// ServiceRepository is an autogenerated mock type for the ServiceRepository type
type ServiceRepository struct {
	mock.Mock
}

// FetchAll provides a mock function with given fields: ctx
func (_m *ServiceRepository) FetchAll(ctx context.Context) ([]domain.Service, error) {
	ret := _m.Called(ctx)

	var r0 []domain.Service
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Service); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDetailByID provides a mock function with given fields: ctx, id
func (_m *ServiceRepository) GetDetailByID(ctx context.Context, id string) ([]domain.ServiceDetail, error) {
	ret := _m.Called(ctx, id)

	var r0 []domain.ServiceDetail
	if rf, ok := ret.Get(0).(func(context.Context, string) []domain.ServiceDetail); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ServiceDetail)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlanByKey provides a mock function with given fields: ctx, planName, serviceId
func (_m *ServiceRepository) GetPlanByKey(ctx context.Context, planName string, serviceId string) (*domain.Plan, error) {
	ret := _m.Called(ctx, planName, serviceId)

	var r0 *domain.Plan
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *domain.Plan); ok {
		r0 = rf(ctx, planName, serviceId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Plan)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, planName, serviceId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}