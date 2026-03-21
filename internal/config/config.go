package config

import (
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
	sensorType := os.Getenv("sensor_type")
	pullerType := os.Getenv("puller_type")

	config := &Config{
		Url:        senderUrl,
		ServerType: pullerType,
		SensorType: sensorType,
	}

	return config, nil
}
