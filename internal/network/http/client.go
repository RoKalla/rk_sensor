package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	s "rk_sensor/internal/sensors"
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
	httpClient http.Client
	httpConfig Config
}

func NewConfig(url string) *Config {
	return &Config{
		url: url,
	}
}

func NewClient(config *Config) *Client {
	return &Client{
		httpClient: http.Client{},
		httpConfig: *config,
	}
}

func (c *Client) Send(data s.Sensor) error {
	contentType := "application/json"
	payload := Payload{
		id:         data.Id(),
		sensorType: data.Type(),
		timeStamp:  data.Timestamp(),
		unit:       data.Unit(),
		value:      data.Value(),
	}
	json, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("Unable to parse payload")
	}
	reader := bytes.NewReader(json)
	resp, err := c.httpClient.Post(c.httpConfig.url, contentType, reader)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	return nil
}
