package sensorhandler

import (
	"fmt"
	"rk_sensor/internal/domain"
	"time"
)

type SensorHandler struct {
	puller domain.Puller
	sender domain.Sender
	sensor domain.Sensor
	stop   chan struct{}
}

func New(sender domain.Sender, sensor domain.Sensor, puller domain.Puller) *SensorHandler {
	return &SensorHandler{
		sender: sender,
		sensor: sensor,
		puller: puller,
		stop:   make(chan struct{}),
	}
}

func (sh *SensorHandler) Start() error {

	if err := sh.sensor.Start(); err != nil {
		return err
	}
	if sh.sender != nil {
		fmt.Println("Starting sender")
		sh.startSender()
	}
	if sh.puller != nil {
		sh.puller.StartServer(sh.sensor)
	}
	return nil
}

func (sh *SensorHandler) Stop() error {

	if err := sh.sensor.Stop(); err != nil {
		return err
	}

	sh.stop <- struct{}{}

	return nil
}

func (sh *SensorHandler) startSender() {
	go func() {
	forloop:
		for {
			select {
			case <-time.After(10 * time.Second):
				if err := sh.sender.Send(sh.sensor); err != nil {
					fmt.Println(err)
				}
			case <-sh.stop:
				break forloop
			}
		}
		// time.Sleep(10 * time.Second)
	}()
}
