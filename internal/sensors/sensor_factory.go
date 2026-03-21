package sensors

import (
	"fmt"
	"rk_sensor/internal/domain"
	"rk_sensor/internal/sensors/temperature"
	"strings"
)

// GetSensor return specified sensortype.
// Return error if specified does not exists
func GetSensor(sensorType string) (domain.Sensor, error) {

	value := strings.ToLower(sensorType)
	switch value {
	case "temperature":
		return temperature.New(), nil
	default:
		return nil, fmt.Errorf("invalid sensor type '%s'", value)
	}

}
