package events

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/pkg/errors"
	"go-microframework/pkg/support"
)

type EventProcessor struct {
	router *message.Router
	logger watermill.LoggerAdapter
	//marshaler
}

func (processor EventProcessor) validateEvent(event interface{}) error {
	// EventHandler's NewEvent must return a pointer, because it is used to unmarshal
	if err := support.IsPointer(event); err != nil {
		return errors.Wrap(err, "command must be a non-nil pointer")
	}

	return nil
}
