package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Url        string
	ServerType string
	SensorType string
}

func ReadEnv() (*Config, error) {

	godotenv.Load(".env")

	senderUrl := os.Getenv("sender_url")
	if senderUrl == "" {
		// return nil, fmt.Errorf("Unabled to find environment variable 'sender_url'")
	}

	sensorType := os.Getenv("sensor_type")
	if sensorType == "" {
		return nil, fmt.Errorf("Unabled to find environment variable 'sensor_type'")
	}

	pullerType := os.Getenv("puller_type")
	if pullerType == "" {
		return nil, fmt.Errorf("Unabled to find environment variable 'pullerType'")
	}

	config := &Config{
		Url:        senderUrl,
		ServerType: pullerType,
		SensorType: sensorType,
	}

	return config, nil
}
