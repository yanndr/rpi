// +build windows darwin linux,!arm

package sensor

import (
	"time"
)

const (
	pulseDelay = 10 * time.Microsecond
)

//DistanceSensor interface for an ultra sound sensor
type DistanceSensor interface {
	Distance() (float64, error)
}

type HCSRO4Sensor struct {
	echo, trigger uint8
}

func NewHCSRO4Sensor(trigger, echo uint8) *HCSRO4Sensor {
	sensor := new(HCSRO4Sensor)

	return sensor
}

func (sensor *HCSRO4Sensor) Distance() (float64, error) {
	return 8, nil
}
