package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DSN  string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		Port: getEnv("PORT", "8080"),
		DSN:  getEnv("DSN", "postgres://postgres:@localhost:5432/todo"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
