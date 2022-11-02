package validator

import (
	"context"
	"errors"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/message/usecase/receive"
)

type ArgsValidator interface {
	ValidateArgs(ctx context.Context, args *receive.Args) error
}

var (
	ErrNotFoundSession = errors.New("session is not found")
)

func New(
	gatewayFindSession GatewayFindSession,
) ArgsValidator {
	return argsValidator{
		gatewayFindSession: gatewayFindSession,
	}
}

type argsValidator struct {
	gatewayFindSession GatewayFindSession
}

func (v argsValidator) ValidateArgs(
	ctx context.Context, args *receive.Args,
) error {
	sessionEntity, err := v.
		gatewayFindSession.
		Call(
			ctx, args.SessionID,
		)
	if err != nil {
		return fmt.Errorf("find session: %w", err)
	}

	if sessionEntity == nil {
		return ErrNotFoundSession
	}

	return nil
}
