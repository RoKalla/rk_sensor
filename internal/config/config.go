package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Url        string
	SensorType string
}

func ReadEnv() (*Config, error) {

	godotenv.Load(".env")

	senderUrl := os.Getenv("sender_url")
	if senderUrl == "" {
		return nil, fmt.Errorf("Unabled to find environment variable 'sender_url'")
	}

	sensorType := os.Getenv("sensor_type")
	if sensorType == "" {
		return nil, fmt.Errorf("Unabled to find environment variable 'sensor_type'")
	}

	config := &Config{
		Url:        senderUrl,
		SensorType: sensorType,
	}

	return config, nil
}
