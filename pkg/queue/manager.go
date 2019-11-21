package queue

import (
	"go-microframework/internal/config"
)

type QueueManager struct {
	/**
	The config instance
	 */
	config *config.Configuration
}