package config

type Configuration struct {
	Server ServerConfiguration
	Database DatabaseConfiguration
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