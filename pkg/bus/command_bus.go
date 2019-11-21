package bus

import "github.com/ThreeDotsLabs/watermill/components/cqrs"

type CommandBus struct {
	cqrs.CommandBus
}
