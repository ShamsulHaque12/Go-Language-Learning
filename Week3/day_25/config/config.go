package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port              string
	Env               string
	DatabaseURL       string
	JWTSecret         string
	JWTExpiresIn      time.Duration
	DBMaxOpenConns    int
	DBMaxIdleConns    int
	DBConnMaxLifetime time.Duration
}

func LoadConfig() *Config {
	// Attempt to load .env file; if it fails, continue with system envs
	if err := godotenv.Load(); err != nil {
		log.Println("[INFO] No .env file found or error loading, reading from environment variables")
	}

	return &Config{
		Port:              getEnv("PORT", "8080"),
		Env:               getEnv("APP_ENV", "development"),
		DatabaseURL:       getEnv("DATABASE_URL", "postgresql://neondb_owner:npg_a1JUNHun4kwl@ep-patient-wind-bei81258-pooler.c-9.us-east-2.aws.neon.tech/neondb?sslmode=require&channel_binding=require"),
		JWTSecret:         getEnv("JWT_SECRET", "super-secret-jwt-key-change-in-production-12345"),
		JWTExpiresIn:      getEnvAsDuration("JWT_EXPIRES_IN", 24*time.Hour),
		DBMaxOpenConns:    getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
		DBMaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS", 10),
		DBConnMaxLifetime: getEnvAsDuration("DB_CONN_MAX_LIFETIME", 30*time.Minute),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	valStr := os.Getenv(key)
	if valStr == "" {
		return fallback
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return fallback
	}
	return val
}

func getEnvAsDuration(key string, fallback time.Duration) time.Duration {
	valStr := os.Getenv(key)
	if valStr == "" {
		return fallback
	}
	val, err := time.ParseDuration(valStr)
	if err != nil {
		return fallback
	}
	return val
}
