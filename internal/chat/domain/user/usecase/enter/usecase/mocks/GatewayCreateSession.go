// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	session "github.com/vdrpkv/goexamples/internal/chat/domain/session"

	user "github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

// GatewayCreateSession is an autogenerated mock type for the GatewayCreateSession type
type GatewayCreateSession struct {
	mock.Mock
}

// Call provides a mock function with given fields: ctx, userID
func (_m *GatewayCreateSession) Call(ctx context.Context, userID user.ID) (*session.Entity, error) {
	ret := _m.Called(ctx, userID)

	var r0 *session.Entity
	if rf, ok := ret.Get(0).(func(context.Context, user.ID) *session.Entity); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*session.Entity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, user.ID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGatewayCreateSession interface {
	mock.TestingT
	Cleanup(func())
}

// NewGatewayCreateSession creates a new instance of GatewayCreateSession. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGatewayCreateSession(t mockConstructorTestingTNewGatewayCreateSession) *GatewayCreateSession {
	mock := &GatewayCreateSession{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
