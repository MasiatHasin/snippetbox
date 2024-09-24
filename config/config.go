package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DB_URL string
}

func LoadConfig() *Config {
	// Load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	// Create and populate Config struct
	cfg := &Config{
		Port:   os.Getenv("SNIPPETBOX_PORT"),
		DB_URL: os.Getenv("DATABASE_URL"),
	}

	return cfg
}
