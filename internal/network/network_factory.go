package network

import (
	"fmt"
	"rk_sensor/internal/domain"
	"rk_sensor/internal/network/http"
	"strings"
)

func GetPuller(pullerType string) (domain.Puller, error) {
	switch pullerType {
	case "":
		return nil, nil
	case "http", "https":
		return http.NewServer(), nil
	default:
		return nil, fmt.Errorf("Puller Type not supported (server): %s", pullerType)
	}
}

func GetSender(url string) (domain.Sender, error) {

	if url == "" {
		return nil, nil
	}

	schema := strings.Split(url, "://")
	if len(schema) < 2 {
		return nil, fmt.Errorf("Schema missing from URL (client): %s", url)
	}
	switch schema[0] {
	case "http", "https":
		return http.NewClient(url), nil
	default:
		return nil, fmt.Errorf("Schema type not supported (client): %s", url)
	}
}
