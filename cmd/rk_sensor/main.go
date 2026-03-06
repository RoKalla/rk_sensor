package main

import (
	"fmt"
	"os"
	"rk_sensor/internal/sensors"
)

func main() {
	fmt.Println("slacko_taow")
	// temprature
	sensor, err := sensors.GetSensor("temprature")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(sensor.Timestamp())

	// sensor := s.New()
	// httpConfig := h.NewConfig("http://127.0.0.1")
	// sender := h.NewClient(httpConfig)
	// controller := c.New(sensor, sender)
	// controller.Start()

	// signals := make(chan os.Signal, 1)
	// signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// <-signals
	// controller.Stop()
}
