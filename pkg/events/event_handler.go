package events

import "context"

type EventHandler interface {
	Handle(ctx context.Context, event interface{}) error
}
