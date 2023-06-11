package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbName      string
	Port        string
	LogPath     string
	MakeMigrate bool
}

func getEnvString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	} else {
		return value
	}
}

func GetConfig() *Config {
	godotenv.Load(".env")

	config := Config{}
	config.DbName = getEnvString("DBNAME", "database")
	config.Port = getEnvString("PORT", "8080")
	config.LogPath = getEnvString("LOG_PATH", "./log")

	makeMigrate := os.Getenv("MIGRATE")
	if makeMigrate == "" || makeMigrate != "true" {
		config.MakeMigrate = false
	} else {
		config.MakeMigrate = true
	}

	return &config
}
