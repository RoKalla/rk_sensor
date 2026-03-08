package sensors

import (
	"fmt"
	"rk_sensor/internal/domain"
	"rk_sensor/internal/sensors/temprature"
	"strings"
)

// GetSensor return specified sensortype.
// Return error if specified does not exists
func GetSensor(sensorType string) (domain.Sensor, error) {

	value := strings.ToLower(sensorType)
	switch value {
	case "temprature":
		return temprature.New(), nil
	default:
		return nil, fmt.Errorf("invalid sensor type '%s'", value)
	}

}
