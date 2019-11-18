package database

import (
	"go-microframework/internal/config"
	"go-microframework/pkg/database/connectors"
	"go.uber.org/fx"
)

type result struct {
	fx.Out

	ConnectionFactory *connectors.ConnectionFactory
	DatabaseManager *Manager
}

func NewService() fx.Option {
	return fx.Options(
		fx.Provide(
			register,
		),
	)
}

func register(config *config.Configuration) result {
	// The connection factory is used to create the actual connection instances on
	// the database. We will inject the factory into the manager so that it may
	// make the connections while they are actually needed and not of before.
	connectionFactory := &connectors.ConnectionFactory{}

	// The database manager is used to resolve various connections, since multiple
	// connections might be managed. It also implements the connection resolver
	// interface which may be used by other components requiring connections.
	databaseManager := &Manager{
		Factory:     connectionFactory,
		Config:      config,
	}

	return result{
		ConnectionFactory: connectionFactory,
		DatabaseManager:   databaseManager,
	}
}