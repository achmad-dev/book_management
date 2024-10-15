package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// App
	Port string `env:"PORT" envDefault:"3000" json:"PORT,omitempty"`
	// Database
	DbHost     string `env:"DB_HOST" envDefault:"localhost" json:"DB_HOST,omitempty"`
	DbPort     string `env:"DB_PORT" envDefault:"5432" json:"DB_PORT,omitempty"`
	DbUser     string `env:"DB_USER" envDefault:"postgres" json:"DB_USER,omitempty"`
	DbPassword string `env:"DB_PASSWORD" envDefault:"postgres" json:"DB_PASSWORD,omitempty"`
	DbName     string `env:"DB_NAME" envDefault:"postgres" json:"DB_NAME,omitempty"`
}

func NewConfig(path string) (*Config, error) {
	// Load from .env file if it exists
	_ = godotenv.Load(path)

	// Create the config from environment variables
	config := Config{
		Port:       getEnv("PORT", "3000"),
		DbHost:     getEnv("DB_HOST", "localhost"),
		DbPort:     getEnv("DB_PORT", "5432"),
		DbUser:     getEnv("DB_USER", "postgres"),
		DbPassword: getEnv("DB_PASSWORD", "postgres"),
		DbName:     getEnv("DB_NAME", "postgres"),
	}

	return &config, nil
}

// Helper function to fetch environment variables with a default fallback
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
