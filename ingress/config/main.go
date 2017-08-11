package config

import "os"

// Configuration represents the app's configuration values
type Configuration struct {
	Mode string
	Port string
}

const (
	// DebugMode represents a dev build
	DebugMode string = "debug"
	// ReleaseMode represents a production build
	ReleaseMode string = "production"
)

var current Configuration

func init() {
	env := os.Getenv("ENV")
	if env != "production" {
		env = DebugMode
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	current = Configuration{Mode: env, Port: port}
}

// Get returns the current configuration values
func Get() Configuration {
	return current
}
