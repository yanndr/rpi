// +build windows darwin linux,!arm

package controller

import (
	"fmt"
)

//SingleMotorController Controller for one motor
type SingleMotorController interface {
	SetSpeed(speed float64)
	Start()
	Stop()
}

type L298NMotorController struct {
	name string
}

func (mc *L298NMotorController) SetSpeed(speed float64) {
	fmt.Println(mc.name, " setSpeed at ", speed)
}

func (mc *L298NMotorController) Start() {
	fmt.Println(mc.name, " started")
}

func (mc *L298NMotorController) Stop() {
	fmt.Println(mc.name, " stopped")
}

func NewL298NMotorController(name string, pin1, pin2, speedPin uint8) *L298NMotorController {
	motor := new(L298NMotorController)
	motor.name = name
	return motor
}
