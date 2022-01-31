package env

import (
	"os"
)

func GetEnvVar(key string, fallbackValue string) string {
	value := os.Getenv(key)

	if value != "" {
		return value
	}

	return fallbackValue
}
