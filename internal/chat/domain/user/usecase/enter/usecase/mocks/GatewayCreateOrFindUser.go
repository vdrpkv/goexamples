// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	user "github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

// GatewayCreateOrFindUser is an autogenerated mock type for the GatewayCreateOrFindUser type
type GatewayCreateOrFindUser struct {
	mock.Mock
}

// Call provides a mock function with given fields: ctx, userName
func (_m *GatewayCreateOrFindUser) Call(ctx context.Context, userName user.Name) (*user.Entity, error) {
	ret := _m.Called(ctx, userName)

	var r0 *user.Entity
	if rf, ok := ret.Get(0).(func(context.Context, user.Name) *user.Entity); ok {
		r0 = rf(ctx, userName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.Entity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, user.Name) error); ok {
		r1 = rf(ctx, userName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGatewayCreateOrFindUser interface {
	mock.TestingT
	Cleanup(func())
}

// NewGatewayCreateOrFindUser creates a new instance of GatewayCreateOrFindUser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGatewayCreateOrFindUser(t mockConstructorTestingTNewGatewayCreateOrFindUser) *GatewayCreateOrFindUser {
	mock := &GatewayCreateOrFindUser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
