package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBReadURL  string
	DBWriteURL string
	GRPCPort   string
}

var AppConfig Config

func LoadConfig() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env")
	}

	AppConfig = Config{
		DBReadURL:  os.Getenv("DB_READ_URL"),
		DBWriteURL: os.Getenv("DB_WRITE_URL"),
		GRPCPort:   os.Getenv("GRPC_PORT"),
	}
}
