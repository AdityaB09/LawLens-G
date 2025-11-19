package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
	DBName   string
	DBSSL    string
}

func LoadConfig() Config {
	_ = godotenv.Load()

	cfg := Config{
		Port:   getEnv("PORT", "8080"),
		DBHost: getEnv("DB_HOST", "db"),
		DBPort: getEnv("DB_PORT", "5432"),
		DBUser: getEnv("DB_USER", "lawlens"),
		DBPass: getEnv("DB_PASSWORD", "lawlens"),
		DBName: getEnv("DB_NAME", "lawlens"),
		DBSSL:  getEnv("DB_SSLMODE", "disable"),
	}

	log.Printf("Loaded config: port=%s db=%s@%s:%s/%s", cfg.Port, cfg.DBUser, cfg.DBHost, cfg.DBPort, cfg.DBName)
	return cfg
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
