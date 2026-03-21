package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rk_sensor/internal/domain"
)

type Server struct {
	sensor domain.Sensor
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) StartServer(sensor domain.Sensor) {
	s.sensor = sensor
	fmt.Println("Starting server...")
	http.HandleFunc("/sensor", s.sensorFunc)
	http.ListenAndServe(":8080", nil)
}

func (s *Server) sensorFunc(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close() //nolint:errcheck
	payload := &Payload{
		Id:         s.sensor.Id(),
		SensorType: s.sensor.Type(),
		TimeStamp:  s.sensor.Timestamp(),
		Unit:       s.sensor.Unit(),
		Value:      s.sensor.Value(),
	}
	json, err := json.Marshal(payload)
	if err != nil {
		fmt.Print("Unable to parse payload")
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)
}
