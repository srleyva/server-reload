package config

// Configuration is the root struct for config type
type Configuration struct {
	Server  ServerConfiguration
	Logging LoggingConfiguration
}

// ServerConfiguration configuration of server
type ServerConfiguration struct {
	Port     int
	Hostname string
	List     bool
	Version  string
}

// LoggingConfiguration set the log config
type LoggingConfiguration struct {
	Debug  bool
	Format string
}
