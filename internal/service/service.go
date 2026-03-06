package service

import (
	sens "rk_sensor/internal/sensors"
)

type Sender interface {
	Send(sensor sens.Sensor) error
}

type SensorController struct {
	sensor sens.Sensor
	sender Sender
}

func New(sensor sens.Sensor, sender Sender) *SensorController {
	return &SensorController{
		sensor: sensor,
		sender: sender,
	}
}

func (s *SensorController) Start() {
	s.sensor.Start()
}

func (s *SensorController) Stop() {
	s.sensor.Stop()
}
