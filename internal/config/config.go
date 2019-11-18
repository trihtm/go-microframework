package config

type Configuration struct {
	Server   ServerConfiguration
}

type ServerConfiguration struct {
	Host string
	Port int
}