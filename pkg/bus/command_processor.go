package bus

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/pkg/errors"
	"go-microframework/pkg/support"
)

type CommandProcessor struct {
	router message.Router
	logger watermill.LoggerAdapter
}

func (processor CommandProcessor) validateCommand(cmd interface{}) error {
	// CommandHandler's NewCommand must return a pointer, because it is used to unmarshal
	if err := support.IsPointer(cmd); err != nil {
		return errors.Wrap(err, "command must be a non-nil pointer")
	}

	return nil
}
