// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "sagara/domain/entity"

	exceptions "sagara/pkg/exceptions"

	mock "github.com/stretchr/testify/mock"

	request "sagara/internal/delivery/request"

	uuid "github.com/google/uuid"
)

// OfficeUseCase is an autogenerated mock type for the OfficeUseCase type
type OfficeUseCase struct {
	mock.Mock
}

// DetailOffice provides a mock function with given fields: ctx, id
func (_m *OfficeUseCase) DetailOffice(ctx context.Context, id uuid.UUID) (*entity.Office, *exceptions.CustomError) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Office
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *entity.Office); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Office)
		}
	}

	var r1 *exceptions.CustomError
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) *exceptions.CustomError); ok {
		r1 = rf(ctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*exceptions.CustomError)
		}
	}

	return r0, r1
}

// ListOffice provides a mock function with given fields: ctx, options
func (_m *OfficeUseCase) ListOffice(ctx context.Context, options *request.Option) ([]*entity.Office, int32, *exceptions.CustomError) {
	ret := _m.Called(ctx, options)

	var r0 []*entity.Office
	if rf, ok := ret.Get(0).(func(context.Context, *request.Option) []*entity.Office); ok {
		r0 = rf(ctx, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Office)
		}
	}

	var r1 int32
	if rf, ok := ret.Get(1).(func(context.Context, *request.Option) int32); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Get(1).(int32)
	}

	var r2 *exceptions.CustomError
	if rf, ok := ret.Get(2).(func(context.Context, *request.Option) *exceptions.CustomError); ok {
		r2 = rf(ctx, options)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*exceptions.CustomError)
		}
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewOfficeUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewOfficeUseCase creates a new instance of OfficeUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOfficeUseCase(t mockConstructorTestingTNewOfficeUseCase) *OfficeUseCase {
	mock := &OfficeUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
