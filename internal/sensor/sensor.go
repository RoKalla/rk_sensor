package sensor

import (
	"math/rand"
	d "sensor/internal/domain"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

type Sensor struct{}

func New() *Sensor {
	return &Sensor{}
}

// func (s *Sensor) Start() error {
// }
// func (s *Sensor) Stop() error {
// }
func (s *Sensor) Read() *d.SensorData {
	return &d.SensorData{
		Data1: randomizeNumber(),
		Data2: randomizeNumber(),
	}
}

func randomizeNumber() int {
	return rng.Intn(100)
}
