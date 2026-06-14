package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBReadURL  string
	DBWriteURL string
	GRPCPort   string
	RedisHost  string
	RedisAddr  string
	JWTSecret  string
	Env        string
	LogLevel   string
}

func LoadConfig() (*Config, error) {
	// Only load .env file if it exists (for local development)
	// In Docker, environment variables are set via docker-compose.yml
	if _, err := os.Stat(".env"); err == nil {
		godotenv.Load()
	}

	cfg := &Config{
		DBReadURL:  os.Getenv("DB_READ_URL"),
		DBWriteURL: os.Getenv("DB_WRITE_URL"),
		GRPCPort:   os.Getenv("GRPC_PORT"),
		RedisHost:  os.Getenv("REDIS_HOST"),
		RedisAddr:  os.Getenv("REDIS_ADDR"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		Env:        os.Getenv("ENV"),
		LogLevel:   os.Getenv("LOG_LEVEL"),
	}

	// Validate required fields
	if cfg.DBReadURL == "" || cfg.DBWriteURL == "" {
		log.Println("failed to load env: missing database URLs")
		return nil, errors.New("missing required environment variables")
	}

	if cfg.GRPCPort == "" {
		log.Println("failed to load env: missing GRPC_PORT")
		return nil, errors.New("missing GRPC_PORT")
	}

	if cfg.JWTSecret == "" {
		log.Println("failed to load env: missing JWT_SECRET")
		return nil, errors.New("missing JWT_SECRET")
	}

	return cfg, nil
}
