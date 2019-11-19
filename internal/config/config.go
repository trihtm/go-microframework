package config

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
	Queue    QueueConfiguration
}

type ServerConfiguration struct {
	Host string
	Port int
}

type DatabaseConfiguration struct {
	/*
	   |--------------------------------------------------------------------------
	   | Default Database Connection Name
	   |--------------------------------------------------------------------------
	   |
	   | Here you may specify which of the database connections below you wish
	   | to use as your default connection for all database work. Of course
	   | you may use many connections at once using the Database library.
	   |
	*/
	Default string

	/*
	   |--------------------------------------------------------------------------
	   | Database Connections
	   |--------------------------------------------------------------------------
	   |
	   | Here are each of the database connections setup for your application.
	   |
	*/
	Connections map[string]DatabaseConnectionConfiguration
}

type DatabaseConnectionConfiguration map[string]interface {
}

type QueueConfiguration struct {
	/*
	   |--------------------------------------------------------------------------
	   | Default Queue Connection Name
	   |--------------------------------------------------------------------------
	   |
	   | Queue API supports an assortment of back-ends via a single
	   | API, giving you convenient access to each back-end using the same
	   | syntax for every one. Here you may define a default connection.
	   |
	*/
	Default string

	/*
	   |--------------------------------------------------------------------------
	   | Queue Connections
	   |--------------------------------------------------------------------------
	   |
	   | Here you may configure the connection information for each server that
	   | is used by your application. You are free to add more.
	   |
	*/
	Connections map[string]QueueConnectionConfiguration
}

type QueueConnectionConfiguration map[string]interface {
}
