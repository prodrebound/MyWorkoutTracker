package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	DBName     string
	Env        string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	return &Config{
		ServerPort: getEnv("SERVER_PORT", ":8080"),
		DBName:     getEnv("DB_NAME", "fitness.db"),
		Env:        getEnv("ENV", "development"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
