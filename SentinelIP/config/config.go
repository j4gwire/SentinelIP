package config

import (
	"fmt"
	"os"
)

// LoadConfig reads configuration from environment variables
func LoadConfig() error {
	// Example for loading an environment variable
	port := os.Getenv("PORT")
	if port == "" {
		return fmt.Errorf("PORT is not set")
	}

	// Load other necessary config values here
	return nil
}
