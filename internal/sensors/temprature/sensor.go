package temprature

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

type TempratureSensor struct {
	id         string
	sensorType string
	timestamp  int64
	unit       string
	value      float32
	stop       chan struct{}
	lock       sync.RWMutex
}

// Creates a new TempratureSensor
func New() *TempratureSensor {
	return &TempratureSensor{
		id:         ksuid.New().String(),
		sensorType: "Temprature",
		timestamp:  time.Now().UTC().Unix(),
		value:      randomizeNumber(),
		unit:       "C",
		stop:       make(chan struct{}),
		lock:       sync.RWMutex{},
	}
}

// Id return the id of the Temprature sensor
func (s *TempratureSensor) Id() string {
	return s.id
}

// Type return the Type of the Temprature sensor
func (s *TempratureSensor) Type() string {
	return s.sensorType
}

// Timestamp return the Timestamp of the Temprature sensor
func (s *TempratureSensor) Timestamp() int64 {
	return s.timestamp
}

// Value return the Value of the Temprature sensor
func (s *TempratureSensor) Value() float32 {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.value
}

// Unit return the Unit of the Temprature sensor
func (s *TempratureSensor) Unit() string {
	return s.unit
}

// Start update the value of the Temprature sensor at intervals
func (s *TempratureSensor) Start() error {
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
func (s *TempratureSensor) Stop() error {
	s.stop <- struct{}{}
	return nil
}
