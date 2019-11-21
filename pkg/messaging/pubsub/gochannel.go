package pubsub

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"go-microframework/internal/config"
)

type GoChannelFactory struct {
	connections map[string]*gochannel.GoChannel
	logger watermill.LoggerAdapter
}

func (ps *GoChannelFactory) Connect(connectionConfig config.QueueConnectionConfiguration) (connection *gochannel.GoChannel) {
	queueName := connectionConfig["queue"].(string)

	connection, found := ps.connections[queueName]

	if found {
		return connection
	}

	connection = ps.newConnection(connectionConfig)

	if ps.connections == nil {
		ps.connections = make(map[string]*gochannel.GoChannel)
	}

	ps.connections[queueName] = connection

	return connection
}

func (ps *GoChannelFactory) newConnection(connectionConfig config.QueueConnectionConfiguration) *gochannel.GoChannel {
	//bufferSize := connectionConfig["buffer_size"].(int64)
	persistent := connectionConfig["persistent"].(bool)
	blockPublish := connectionConfig["block_publish"].(bool)

	return gochannel.NewGoChannel(
		gochannel.Config{
			//OutputChannelBuffer: bufferSize,
			Persistent: persistent,
			BlockPublishUntilSubscriberAck: blockPublish,
		},
		ps.logger,
	)
}