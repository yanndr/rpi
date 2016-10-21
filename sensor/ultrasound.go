// +build linux,arm

package sensor

import (
	"time"

	"github.com/golang/glog"
	"github.com/stianeikeland/go-rpio"
	"github.com/yanndr/rpi/gpio"
)

const (
	pulseDelay = 10 * time.Microsecond
)

//DistanceSensor interface for an ultra sound sensor
type DistanceSensor interface {
	Distance() (float64, error)
}

type HCSRO4Sensor struct {
	echo, trigger rpio.Pin
}

func NewHCSRO4Sensor(trigger, echo uint8) *HCSRO4Sensor {
	sensor := new(HCSRO4Sensor)
	sensor.echo = rpio.Pin(echo)
	sensor.trigger = rpio.Pin(trigger)
	sensor.trigger.Output()
	sensor.echo.Input()

	return sensor
}

func (sensor *HCSRO4Sensor) Distance() (float64, error) {

	sensor.trigger.Low()
	time.Sleep(time.Microsecond * 3)

	// Generate a TRIGGER pulse
	sensor.trigger.High()
	time.Sleep(pulseDelay)
	sensor.trigger.Low()

	glog.V(2).Infof("HCSRO4: waiting for echo to go high")

	duration, err := gpio.PulseDuration(sensor.echo, rpio.High)

	if err != nil {
		return 0, err
	}

	// Calculate the distance based on the time computed
	distance := float64(duration.Nanoseconds()) / 10000000 * (340 / 2)

	return distance, nil
}
