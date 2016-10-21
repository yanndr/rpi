package controller

import (
	"testing"
)

func TestDualMotor(t *testing.T) {
	dmc, _ := NewDualMotorController(NewL298NMotorController("left motor", 1, 2, 3),
		NewL298NMotorController("right motor", 4, 5, 6),
		0,
		0,
		1,
		0.1)

	dmc.Start()
	dmc.SetSpeed(1)
	dmc.Stop()
}
