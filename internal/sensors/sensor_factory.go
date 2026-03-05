package sensors

import (
	"fmt"
	"rk_sensor/internal/sensors/temprature"
	"strings"
)

func GetSensor(name string) (Sensor, error) {

	value := strings.ToLower(name)
	switch value {
	case "temprature":
		return temprature.New(), nil
	default:
		return nil, fmt.Errorf("invalid sensor type '%s'", value)
	}

}
