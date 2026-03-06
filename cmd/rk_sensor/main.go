package main

import (
	"fmt"
	"os"
	"os/signal"
	"rk_sensor/internal/network"
	"rk_sensor/internal/network/server"
	"rk_sensor/internal/sensors"
	"syscall"
	"time"
)

func main() {

	sender, SendErr := network.GetSender("http://localhost:8080/hello")
	if SendErr != nil {
		fmt.Println(SendErr)
		os.Exit(1)
	}

	sensor, SensErr := sensors.GetSensor("temprature", sender)
	if SensErr != nil {
		fmt.Println(SensErr)
		os.Exit(1)
	}

	go server.Start()

	time.Sleep(time.Second)

	fmt.Println("Server started!")
	sensor.Start()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals
}
