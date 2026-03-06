package main

import (
	"fmt"
	"os"
	"os/signal"
	h "rk_sensor/internal/network/http"
	"rk_sensor/internal/sensors"
	"syscall"
)

func main() {
	httpConfig := h.NewConfig("http://localhost:8080/hello")

	sender := h.NewClient(httpConfig)

	sensor, err := sensors.GetSensor("temprature", sender)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Server started!")
	sensor.Start()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals
}
