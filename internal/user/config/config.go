package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// App
	Port      string `env:"PORT" envDefault:"3000" json:"PORT,omitempty"`
	JwtSecret string `env:"JWT_SECRET" json:"JWT_SECRET,omitempty"`
	// Database
	DbHost     string `env:"DB_HOST" envDefault:"localhost" json:"DB_HOST,omitempty"`
	DbPort     string `env:"DB_PORT" envDefault:"5432" json:"DB_PORT,omitempty"`
	DbUser     string `env:"DB_USER" envDefault:"postgres" json:"DB_USER,omitempty"`
	DbPassword string `env:"DB_PASSWORD" envDefault:"postgres" json:"DB_PASSWORD,omitempty"`
	DbName     string `env:"DB_NAME" envDefault:"postgres" json:"DB_NAME,omitempty"`

	// Redis
	RedisHost     string `env:"REDIS_HOST" envDefault:"localhost" json:"REDIS_HOST,omitempty"`
	RedisPort     string `env:"REDIS_PORT" envDefault:"6379" json:"REDIS_PORT,omitempty"`
	RedisPassword string `env:"REDIS_PASSWORD" envDefault:"" json:"REDIS_PASSWORD,omitempty"`
	RedisDB       string `env:"REDIS_DB" envDefault:"0" json:"REDIS_DB,omitempty"`

	// Auth
	Cost string `env:"COST" envDefault:"10" json:"COST,omitempty"`

	// Microservice
	AuthorServiceHost   string `env:"AUTHOR_SERVICE_HOST" envDefault:"localhost:4040" json:"AUTHOR_SERVICE_HOST,omitempty"`
	CategoryServiceHost string `env:"CATEGORY_SERVICE_HOST" envDefault:"localhost:5050" json:"CATEGORY_SERVICE_HOST,omitempty"`
	BookServiceHost     string `env:"BOOK_SERVICE_HOST" envDefault:"localhost:6060" json:"BOOK_SERVICE_HOST,omitempty"`
}

func NewConfig(path string) (*Config, error) {
	// Load from .env file if it exists
	_ = godotenv.Load(path)

	// Create the config from environment variables
	config := Config{
		Port:                getEnv("PORT", "3000"),
		JwtSecret:           getEnv("JWT_SECRET", ""),
		DbHost:              getEnv("DB_HOST", "localhost"),
		DbPort:              getEnv("DB_PORT", "5432"),
		DbUser:              getEnv("DB_USER", "postgres"),
		DbPassword:          getEnv("DB_PASSWORD", "postgres"),
		DbName:              getEnv("DB_NAME", "postgres"),
		RedisHost:           getEnv("REDIS_HOST", "localhost"),
		RedisPort:           getEnv("REDIS_PORT", "6379"),
		RedisPassword:       getEnv("REDIS_PASSWORD", ""),
		RedisDB:             getEnv("REDIS_DB", "0"),
		Cost:                getEnv("COST", "10"),
		AuthorServiceHost:   getEnv("AUTHOR_SERVICE_HOST", "localhost:4040"),
		CategoryServiceHost: getEnv("CATEGORY_SERVICE_HOST", "localhost:5050"),
		BookServiceHost:     getEnv("BOOK_SERVICE_HOST", "localhost:6060"),
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
