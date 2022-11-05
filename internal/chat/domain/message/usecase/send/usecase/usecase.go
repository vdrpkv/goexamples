package usecase

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/message/usecase/send"
)

type UseCase interface {
	Do(ctx context.Context, args *send.Args) (*send.Result, error)
}

func New(
	gatewayCreateMessage GatewayCreateMessage,
	gatewayNotifySessionsAboutNewMessage GatewayNotifySessionsAboutNewMessage,
) UseCase {
	return useCase{
		gatewayCreateMessage:                 gatewayCreateMessage,
		gatewayNotifySessionsAboutNewMessage: gatewayNotifySessionsAboutNewMessage,
	}
}

type useCase struct {
	gatewayCreateMessage                 GatewayCreateMessage
	gatewayNotifySessionsAboutNewMessage GatewayNotifySessionsAboutNewMessage
}

func (uc useCase) Do(
	ctx context.Context,
	args *send.Args,
) (
	*send.Result,
	error,
) {
	messageEntity, err := uc.
		gatewayCreateMessage.
		Call(
			ctx,
			args.AuthorUserID,
			args.MessageText,
		)
	if err != nil {
		return nil, fmt.Errorf("create message: %w", err)
	}

	err = uc.
		gatewayNotifySessionsAboutNewMessage.
		Call(
			ctx, messageEntity,
		)
	if err != nil {
		return nil, fmt.Errorf("notify sessions about new message: %w", err)
	}

	return &send.Result{}, nil
}
