package controller

import (
	d "sensor/internal/domain"
	"time"
)

type Sender interface {
	Send(data d.SensorData) error
}

type Sensor interface {
	Read() *d.SensorData
}

type SensorController struct {
	sensor  Sensor
	sender  Sender
	running bool
}

func New(sensor Sensor, sender Sender) *SensorController {
	return &SensorController{
		sensor:  sensor,
		sender:  sender,
		running: false,
	}
}

func (s *SensorController) Read() *d.SensorData {
	return s.sensor.Read()
}

func (s *SensorController) Start() {
	s.running = true
	go func() {
		for s.running {
			data := s.Read()
			s.sender.Send(*data)
			time.Sleep(5 * time.Second)
		}
	}()
}

func (s *SensorController) Stop() {
	s.running = false
}
