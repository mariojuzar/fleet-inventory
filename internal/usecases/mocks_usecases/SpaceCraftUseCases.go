// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks_usecases

import (
	context "context"

	request "github.com/mariojuzar/fleet-inventory/internal/usecases/request"
	mock "github.com/stretchr/testify/mock"

	response "github.com/mariojuzar/fleet-inventory/internal/usecases/response"
)

// SpaceCraftUseCases is an autogenerated mock type for the SpaceCraftUseCases type
type SpaceCraftUseCases struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, req
func (_m *SpaceCraftUseCases) Create(ctx context.Context, req *request.SpaceShipCreateRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *request.SpaceShipCreateRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *SpaceCraftUseCases) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Edit provides a mock function with given fields: ctx, id, req
func (_m *SpaceCraftUseCases) Edit(ctx context.Context, id int, req *request.SpaceShipEditRequest) error {
	ret := _m.Called(ctx, id, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, *request.SpaceShipEditRequest) error); ok {
		r0 = rf(ctx, id, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, req
func (_m *SpaceCraftUseCases) Fetch(ctx context.Context, req *request.SpaceShipFetchRequest) ([]response.SpaceCraftFetchResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 []response.SpaceCraftFetchResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *request.SpaceShipFetchRequest) ([]response.SpaceCraftFetchResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *request.SpaceShipFetchRequest) []response.SpaceCraftFetchResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]response.SpaceCraftFetchResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *request.SpaceShipFetchRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, id
func (_m *SpaceCraftUseCases) Get(ctx context.Context, id int) (*response.SpaceCraftResponse, error) {
	ret := _m.Called(ctx, id)

	var r0 *response.SpaceCraftResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*response.SpaceCraftResponse, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *response.SpaceCraftResponse); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*response.SpaceCraftResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSpaceCraftUseCases creates a new instance of SpaceCraftUseCases. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSpaceCraftUseCases(t interface {
	mock.TestingT
	Cleanup(func())
}) *SpaceCraftUseCases {
	mock := &SpaceCraftUseCases{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
