package network

import (
	"fmt"
	"rk_sensor/internal/domain"
	"rk_sensor/internal/network/http"
)

func GetSender(senderType, target string) (domain.Sender, error) {
	switch senderType {
	case "http", "https":
		return http.NewClient(target), nil
	default:
		return nil, fmt.Errorf("Sender type not supproted")
	}
}
