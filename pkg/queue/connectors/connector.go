package connectors

import (
	"go-microframework/pkg/contracts/queue"
)

type Connector interface {
	/**
	Establish a queue connection
	 */
	Connect(config interface{}) queue.Queue
}