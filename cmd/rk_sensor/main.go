package main

import (
	"fmt"
	"os"
	"os/signal"
	c "rk_sensor/internal/config"
	"rk_sensor/internal/network"
	sensorhandler "rk_sensor/internal/sensorHandler"
	"rk_sensor/internal/sensors"
	"syscall"
)

func main() {
	config, configErr := c.ReadEnv()
	if configErr != nil {
		fmt.Println(configErr)
		// os.Exit(1)
	}
	sender, SendErr := network.GetSender("")
	if SendErr != nil {
		fmt.Println(SendErr)
		os.Exit(1)
	}

	puller, PullErr := network.GetPuller(config.ServerType)
	if PullErr != nil {
		fmt.Println(PullErr)
		os.Exit(1)
	}

	sensor, SensErr := sensors.GetSensor(config.SensorType)
	if SensErr != nil {
		fmt.Println(SensErr)
		os.Exit(1)
	}

	handler := sensorhandler.New(sender, sensor, puller)

	handler.Start()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals
}
