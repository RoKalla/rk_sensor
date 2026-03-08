package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"rk_sensor/internal/domain"
)

type Payload struct {
	id         string  `json:"id"`
	sensorType string  `json:"sensorType"`
	timeStamp  int64   `json:"timestamp"`
	unit       string  `json:"unit"`
	value      float32 `json:"value"`
}

type Config struct {
	url string
}

type Client struct {
	httpClient *http.Client
	target     string
	// httpConfig *Config
}

// func NewConfig(url string) *Config {
// 	return &Config{
// 		url: url,
// 	}
// }

func NewClient(target string) *Client {
	return &Client{
		httpClient: &http.Client{},
		target:     target,
	}
}

func (c *Client) Send(sensor domain.Sensor) error {
	contentType := "application/json"
	payload := &Payload{
		id:         sensor.Id(),
		sensorType: sensor.Type(),
		timeStamp:  sensor.Timestamp(),
		unit:       sensor.Unit(),
		value:      sensor.Value(),
	}
	json, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("Unable to parse payload")
	}
	reader := bytes.NewReader(json)
	resp, err := c.httpClient.Post(c.target, contentType, reader)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	return nil
}
