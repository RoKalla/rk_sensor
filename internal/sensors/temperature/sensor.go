package temperature

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/segmentio/ksuid"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// randomizeNumber return a (psudo) random number between 0 and 100
func randomizeNumber() float32 {
	return rng.Float32() * 100
}

type TemperatureSensor struct {
	id         string
	sensorType string
	timestamp  int64
	unit       string
	value      float32
	stop       chan struct{}
	lock       sync.RWMutex
}

// Creates a new TemperatureSensor
func New() *TemperatureSensor {
	return &TemperatureSensor{
		id:         ksuid.New().String(),
		sensorType: "Temperature",
		timestamp:  time.Now().UTC().Unix(),
		value:      randomizeNumber(),
		unit:       "C",
		stop:       make(chan struct{}),
		lock:       sync.RWMutex{},
	}
}

// Id return the id of the Temperature sensor
func (s *TemperatureSensor) Id() string {
	return s.id
}

// Type return the Type of the Temperature sensor
func (s *TemperatureSensor) Type() string {
	return s.sensorType
}

// Timestamp return the Timestamp of the Temperature sensor
func (s *TemperatureSensor) Timestamp() int64 {
	return s.timestamp
}

// Value return the Value of the Temperature sensor
func (s *TemperatureSensor) Value() float32 {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.value
}

// Unit return the Unit of the Temperature sensor
func (s *TemperatureSensor) Unit() string {
	return s.unit
}

// Start update the value of the Temperature sensor at intervals
func (s *TemperatureSensor) Start() error {
	go func() {
		fmt.Printf("Sensor %s started\n", s.id)
	forloop:
		for {
			select {
			case <-s.stop:
				fmt.Printf("Sensor %s stopped\n", s.id)
				break forloop
			case <-time.After(5 * time.Second):
				s.lock.Lock()
				s.value = randomizeNumber()
				s.lock.Unlock()
			}
		}
	}()
	return nil
}

// Stop stops the update interval started by calling Start
func (s *TemperatureSensor) Stop() error {
	s.stop <- struct{}{}
	return nil
}
