package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	AppEnv        string
	ServerAddress string
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	ShutdownGrace time.Duration
}

var AppConfig Config

func Load() {
	// Load from .env file (if present)
	_ = godotenv.Load()

	AppConfig = Config{
		AppEnv:        getEnv("APP_ENV", "development"),
		ServerAddress: getEnv("SERVER_ADDRESS", "localhost:8080"),
		ReadTimeout:   getEnvAsDuration("READ_TIMEOUT", 5*time.Second),
		WriteTimeout:  getEnvAsDuration("WRITE_TIMEOUT", 10*time.Second),
		ShutdownGrace: getEnvAsDuration("SHUTDOWN_GRACE", 5*time.Second),
	}
}

func getEnv(key, fallback string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return fallback
}

func getEnvAsDuration(key string, fallback time.Duration) time.Duration {
	valStr := getEnv(key, "")
	if valStr == "" {
		return fallback
	}
	valInt, err := strconv.Atoi(valStr)
	if err != nil {
		log.Info().Msg("Invalid duration for " + key + " : " + valStr + ", using fallback")
		return fallback
	}
	return time.Duration(valInt) * time.Second
}
