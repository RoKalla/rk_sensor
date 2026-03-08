package sensorhandler

import (
	"fmt"
	"rk_sensor/internal/domain"
	"time"
)

type SensorHandler struct {
	sender domain.Sender
	sensor domain.Sensor
	stop   chan struct{}
}

func New(sender domain.Sender, sensor domain.Sensor) *SensorHandler {
	return &SensorHandler{
		sender: sender,
		sensor: sensor,
		stop:   make(chan struct{}),
	}
}

func (sh *SensorHandler) Start() error {

	if err := sh.sensor.Start(); err != nil {
		return err
	}
	sh.startLoop()
	return nil
}

func (sh *SensorHandler) Stop() error {

	if err := sh.sensor.Stop(); err != nil {
		return err
	}

	sh.stop <- struct{}{}

	return nil
}

func (sh *SensorHandler) startLoop() {
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
