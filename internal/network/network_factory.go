package network

import (
	"fmt"
	"rk_sensor/internal/domain"
	"rk_sensor/internal/network/http"
	"strings"
)

func GetSender(url string) (domain.Sender, error) {

	schema := strings.Split(url, "://")
	if len(schema) < 2 {
		return nil, fmt.Errorf("Schema missing from URL: %s", url)
	}

	switch schema[0] {
	case "http", "https":
		return http.NewClient(url), nil
	default:
		return nil, fmt.Errorf("Schema type not supported: %s", url)
	}
}
