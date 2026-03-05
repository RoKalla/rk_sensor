package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	d "sensor/internal/domain"
)

type Payload struct {
	Data1 int `json:"data1"`
	Data2 int `json:"data2"`
}

type Config struct {
	url string // Target for post requests
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

func (c *Client) Send(data d.SensorData) error {
	contentType := "application/json"
	payload := Payload(data)
	json, _ := json.Marshal(payload)
	reader := bytes.NewReader(json)
	resp, err := c.httpClient.Post(c.httpConfig.url, contentType, reader)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	return nil
}
