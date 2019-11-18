package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-microframework/internal/config"
	"go-microframework/pkg/database/connectors"
)

type Manager struct {
	Factory *connectors.ConnectionFactory
	Config  *config.Configuration

	connections map[string]*gorm.DB
}

/**
Get the default connection name.
*/
func (manager *Manager) GetDefaultConnection() string {
	return manager.Config.Database.Default
}

/**
Get a database connection instance.
*/
func (manager *Manager) Connection(name string) *gorm.DB {
	if name == "" {
		name = manager.GetDefaultConnection()
	}

	if connection, ok := manager.connections[name]; ok {
		return connection
	}

	connection := manager.makeConnection(name)

	// If we haven't created this connection, we'll create it based on the config
	// provided in the application.
	if manager.connections == nil {
		manager.connections = make(map[string]*gorm.DB)
	}

	manager.connections[name] = connection

	return connection
}

/**
Make the database connection instance.
*/
func (manager *Manager) makeConnection(name string) *gorm.DB {
	return manager.Factory.Make(manager.Config.Database, name)
}

/**
Get the configuration for a connection.
*/
func (manager *Manager) configuration(name string) interface{} {
	connectionsConfig := manager.Config.Database.Connections

	if connectionConfig, ok := connectionsConfig[name]; ok {
		return connectionConfig
	}

	panic(fmt.Errorf("Database [%s] not configured.\n", name))
}
