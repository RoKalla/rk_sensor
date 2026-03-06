package temprature

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"rk_sensor/internal/domain"

	"github.com/segmentio/ksuid"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomizeNumber() float32 {
	return rng.Float32() * 100
}

type TempratureSensor struct {
	id         string
	sensorType string
	timestamp  int64
	unit       string
	value      float32
	stop       chan struct{}
	lock       sync.RWMutex
	sender     domain.Sender
}

func New(sender domain.Sender) *TempratureSensor {
	return &TempratureSensor{
		id:         ksuid.New().String(),
		sensorType: "Temprature",
		timestamp:  time.Now().UTC().Unix(),
		value:      randomizeNumber(),
		unit:       "C",
		stop:       make(chan struct{}),
	}
}

func (s *TempratureSensor) Id() string {
	return s.id
}

func (s *TempratureSensor) Type() string {
	return s.sensorType
}

func (s *TempratureSensor) Timestamp() int64 {
	return s.timestamp
}

func (s *TempratureSensor) Value() float32 {
	return s.value
}

func (s *TempratureSensor) Unit() string {
	return s.unit
}

func (s *TempratureSensor) Start() error {
	go func() {
		fmt.Printf("Sensor %s started\n", s.id)
	forloop:
		for {
			select {
			case <-s.stop:
				fmt.Printf("Sensor %s stopped\n", s.id)
				break forloop
			case <-time.After(1 * time.Second):
				s.lock.Lock()
				s.value = randomizeNumber()
				fmt.Println(s.value)
				s.lock.Unlock()
			}
		}
	}()
	return nil
}

func (s *TempratureSensor) Stop() error {
	s.stop <- struct{}{}
	return nil
}
