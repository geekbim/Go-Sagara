// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "sagara/domain/entity"

	mock "github.com/stretchr/testify/mock"

	request "sagara/internal/delivery/request"
)

// RegionRepository is an autogenerated mock type for the RegionRepository type
type RegionRepository struct {
	mock.Mock
}

// CountRegion provides a mock function with given fields: ctx, options
func (_m *RegionRepository) CountRegion(ctx context.Context, options *request.Option) (int32, error) {
	ret := _m.Called(ctx, options)

	var r0 int32
	if rf, ok := ret.Get(0).(func(context.Context, *request.Option) int32); ok {
		r0 = rf(ctx, options)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *request.Option) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRegion provides a mock function with given fields: ctx, options
func (_m *RegionRepository) GetRegion(ctx context.Context, options *request.Option) ([]*entity.Region, error) {
	ret := _m.Called(ctx, options)

	var r0 []*entity.Region
	if rf, ok := ret.Get(0).(func(context.Context, *request.Option) []*entity.Region); ok {
		r0 = rf(ctx, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Region)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *request.Option) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRegionRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRegionRepository creates a new instance of RegionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRegionRepository(t mockConstructorTestingTNewRegionRepository) *RegionRepository {
	mock := &RegionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
