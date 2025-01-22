package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string
	DatabaseURL   string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		DatabaseURL:   os.Getenv("DATABASE_URL"),
	}, nil
}
