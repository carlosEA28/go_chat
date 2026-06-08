package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server ServerConfig
	// Database DatabaseConfig
}

type ServerConfig struct {
	Port    string
	GinMode string
}

// type DatabaseConfig struct {
// 	Host     string
// 	Port     string
// 	User     string
// 	Password string
// 	Name     string
// 	SSLMode  string
// }

func Load() (*Config, error) {
	_ = godotenv.Load()

	return &Config{
		Server: ServerConfig{
			Port:    getEnv("PORT", "8080"),
			GinMode: getEnv("GIN_MODE", "debug"),
		},
		// Database: DatabaseConfig{
		// 	Host:     ,
		// 	Port:     ,
		// 	User:     ,
		// 	Password: ,
		// 	Name:     ,
		// 	SSLMode:  ,
		// },
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
