// +build linux,arm

package gpio

import (
	"errors"
	"time"

	"github.com/stianeikeland/go-rpio"
)

const (
	maxIncremenation = 200000
)

//PulseDuration Duration of a pulse on one pin
func PulseDuration(pin rpio.Pin, state rpio.State) (time.Duration, error) {

	aroundState := rpio.Low
	if state == rpio.Low {
		aroundState = rpio.High
	}

	// Wait for any previous pulse to end
	i := 0
	for i = 0; i < maxIncremenation; i++ {
		v := pin.Read()
		if v == aroundState {
			break
		}
	}

	if i == maxIncremenation {
		return 0, errors.New("Error: Previous pulse never end")
	}

	// Wait until ECHO goes high
	for i = 0; i < maxIncremenation; i++ {
		v := pin.Read()

		if v == state {
			break
		}
	}

	if i == maxIncremenation {
		return 0, errors.New("Error: Echo never went high")
	}
	startTime := time.Now() // Record time when ECHO goes high

	// Wait until ECHO goes low
	for i = 0; i < maxIncremenation; i++ {
		v := pin.Read()

		if v == aroundState {
			break
		}
	}

	if i == maxIncremenation {
		return 0, errors.New("Error: Echo never went low")
	}

	return time.Since(startTime), nil // Calculate time lapsed for ECHO to transition from high to low
}

func Open() error {
	return rpio.Open()
}

func Close() {
	rpio.Close()
}
