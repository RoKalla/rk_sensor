package temprature

import (
	"math/rand"
	"time"

	"github.com/segmentio/ksuid"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomizeNumber() float32 {
	return rng.Float32() * 100
}

type TempratureSensor struct {
	id         string
	sensorType string
	timeStamp  int64
	unit       string
	value      float32
}

func New() *TempratureSensor {
	return &TempratureSensor{
		id:         ksuid.New().String(),
		sensorType: "Temprature",
		timeStamp:  time.Now().UTC().Unix(),
		value:      randomizeNumber(),
		unit:       "C",
	}
}

func (s *TempratureSensor) Id() string {
	return s.id
}

func (s *TempratureSensor) Type() string {
	return s.sensorType
}

func (s *TempratureSensor) Timestamp() int64 {
	return s.timeStamp
}

func (s *TempratureSensor) Value() float32 {
	return s.value
}

func (s *TempratureSensor) Unit() string {
	return s.unit
}
