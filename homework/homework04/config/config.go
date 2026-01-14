package config

import (
	"os"
)

var (
	JWTSecret string
)

func LoadConfig() {
	JWTSecret = getEnv("JWT_SECRET", "your_secret_key")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}