package connectors

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-microframework/internal/config"
)

type MySQLConnector struct {
}

/**
Establish a database connection.
 */
func (connector MySQLConnector) Connect(config config.DatabaseConnectionConfiguration) *gorm.DB {
	db, err := gorm.Open("mysql", connector.getDSN(config))

	if err != nil {
		panic(fmt.Errorf("Cannot connect to database || %s \n", err))
	}

	return db
}

/**
Create a DSN (Data Source Name) from a configuration.
 */
func (connector MySQLConnector) getDSN(config config.DatabaseConnectionConfiguration) string {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		config["username"], config["password"], config["host"], config["database"], config["charset"],
	)

	return connectionString
}