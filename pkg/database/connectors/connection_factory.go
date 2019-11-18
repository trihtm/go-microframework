package connectors

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-microframework/internal/config"
)

type ConnectionFactory struct {

}

/**
 * Establish a GORM connection based on the configuration
 */
func (factory *ConnectionFactory) Make(config config.DatabaseConfiguration, name string) *gorm.DB {
	connectionConfig := config.Connections[name]

	return factory.createConnection(connectionConfig)
}

/**
 * Create a new connection instance.
 */
func (factory *ConnectionFactory) createConnection(connectionConfig config.DatabaseConnectionConfiguration) (connection *gorm.DB) {
	driver := connectionConfig["driver"].(string)

	switch driver {
	case "mysql":
		connection = MySQLConnector{}.Connect(connectionConfig)
	default:
		panic(fmt.Errorf("Unsupported driver [%s]\n", driver))
	}

	return connection
}