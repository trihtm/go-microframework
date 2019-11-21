package bus

import "context"

type CommandHandler interface {
	Handle(ctx context.Context, cmd interface{}) error
}
