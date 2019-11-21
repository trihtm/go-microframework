package events

import "github.com/ThreeDotsLabs/watermill/components/cqrs"

type EventBus struct {
	cqrs.EventBus
}
