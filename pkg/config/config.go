package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port         string
	Environment  string
	LogLevel     string
	ReadTimeout  int
	WriteTimeout int
}

func Load() *Config {
	return &Config{
		Port:         getEnv("PORT", "8080"),
		Environment:  getEnv("ENVIRONMENT", "development"),
		LogLevel:     getEnv("LOG_LEVEL", "info"),
		ReadTimeout:  getEnvAsInt("READ_TIMEOUT", 15),
		WriteTimeout: getEnvAsInt("WRITE_TIMEOUT", 15),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
