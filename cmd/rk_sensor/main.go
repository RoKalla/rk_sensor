package main

import (
	"fmt"
	"os"
	"os/signal"
	h "rk_sensor/internal/network/http"
	"rk_sensor/internal/sensors"
	"rk_sensor/internal/service"
	"syscall"
	"time"
)

func main() {
	fmt.Println("slacko_taow")
	// temprature
	sensor, err := sensors.GetSensor("temprature")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	httpConfig := h.NewConfig("http://127.0.0.1")
	sender := h.NewClient(httpConfig)
	service := service.New(sensor, sender)
	service.Start()

	time.Sleep(5 * time.Second)

	service.Stop()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals
}
