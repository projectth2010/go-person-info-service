// utils/env.go
package utils

import "os"

// GetEnv gets an environment variable or returns a default value
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// GetEnvRequired gets an environment variable or panics if not found
func GetEnvRequired(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	panic("Environment variable not set: " + key)
}
