package messagebus

import (
	usecaseMessageSendGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/gateway/message"
	usecaseMessageSubscribeGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/message/subscribe/gateway/message"
	usecaseMessageUnsubscribeGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/message/unsubscribe/gateway/message"
)

type Gateways struct {
	UseCaseMessageSendGateways        UseCaseMessageSendGateways
	UseCaseMessageSubscribeGateways   UseCaseMessageSubscribeGateways
	UseCaseMessageUnsubscribeGateways UseCaseMessageUnsubscribeGateways
}

type UseCaseMessageSendGateways struct {
	Broadcaster usecaseMessageSendGateway.Broadcaster
}

type UseCaseMessageSubscribeGateways struct {
	Subscriber usecaseMessageSubscribeGateway.Subscriber
}

type UseCaseMessageUnsubscribeGateways struct {
	Unsubscriber usecaseMessageUnsubscribeGateway.Unsubscriber
}
