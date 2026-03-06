package main

import (
	"fmt"
	"os"
	"os/signal"
	h "rk_sensor/internal/network/http"
	"rk_sensor/internal/sensors"
	"syscall"
	"time"
)

func main() {
	// temprature

	httpConfig := h.NewConfig("http://127.0.0.1")
	sender := h.NewClient(httpConfig)

	sensor, err := sensors.GetSensor("temprature", sender)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sensor.Start()

	time.Sleep(5 * time.Second)

	sensor.Stop()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals
}
