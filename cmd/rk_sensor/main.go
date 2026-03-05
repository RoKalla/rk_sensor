package main

import (
	"os"
	"os/signal"
	h "sensor/internal/clients/http"
	c "sensor/internal/controller"
	s "sensor/internal/sensor"
	"syscall"
)

func main() {
	sensor := s.New()
	httpConfig := h.NewConfig("http://127.0.0.1")
	sender := h.NewClient(httpConfig)
	controller := c.New(sensor, sender)
	controller.Start()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals
	controller.Stop()
}
