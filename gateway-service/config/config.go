package config

import (
	"os"
)

type Config struct {
	HTTPPort             string
	AuthServiceAddr      string
	InterviewServiceAddr string
	LogLevel             string
	RequestTimeout       int
}

func LoadConfig() *Config {
	return &Config{
		HTTPPort:             getEnv("HTTP_PORT", "8080"),
		AuthServiceAddr:      getEnv("AUTH_SERVICE_ADDR", "localhost:50051"),
		InterviewServiceAddr: getEnv("INTERVIEW_SERVICE_ADDR", "localhost:50051"),
		LogLevel:             getEnv("LOG_LEVEL", "info"),
		RequestTimeout:       10,
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
