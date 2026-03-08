package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Url string
}

func ReadEnv() (*Config, error) {

	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	sender_url := os.Getenv("sender_url")
	if sender_url == "" {
		return nil, fmt.Errorf("Unabled to find Environment variable url")
	}

	config := &Config{
		Url: sender_url,
	}

	return config, nil
}
