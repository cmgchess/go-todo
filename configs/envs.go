package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser string
	DBPass string
	DBHost string
	DBName string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		DBUser: getEnv("DB_USER", "postgres"),
		DBPass: getEnv("DB_PASS", ""),
		DBHost: getEnv("DB_HOST", "localhost"),
		DBName: getEnv("DB_NAME", "todo"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func (cfg Config) FormatDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, 5432, cfg.DBName,
	)
}
