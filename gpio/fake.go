// +build windows darwin linux,!arm

package gpio

import (
	"time"
)

const (
	maxIncremenation = 200000
)

//PulseDuration Duration of a pulse on one pin
func PulseDuration(pin uint8, state uint8) (time.Duration, error) {

	startTime := time.Now() // Record time when ECHO goes high

	return time.Since(startTime), nil // Calculate time lapsed for ECHO to transition from high to low
}

func Open() error {
	return nil
}

func Close() {

}
