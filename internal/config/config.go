package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server ServerConfig
	Mongo  MongoConfig
}

type ServerConfig struct {
	Port    string
	GinMode string
}

type MongoConfig struct {
	URI      string
	Database string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	return &Config{
		Server: ServerConfig{
			Port:    getEnv("PORT", "8080"),
			GinMode: getEnv("GIN_MODE", "debug"),
		},
		Mongo: MongoConfig{
			URI:      getEnv("MONGO_DB_URL", getEnv("MONGO_URI", "mongodb://localhost:27017")),
			Database: getEnv("MONGO_DATABASE", "go_chat"),
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
