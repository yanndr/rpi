// +build windows darwin linux,!arm

package controller

import (
	"fmt"

	"github.com/yanndr/rpi/pwm"
)

type BaseMotorController struct {
	name string
}

type L298NMotorController struct {
	BaseMotorController
}

type DRV833MotorController struct {
	BaseMotorController
}

func (mc *BaseMotorController) SetSpeed(speed float64) {
	fmt.Println(mc.name, " setSpeed at ", speed)
}

func (mc *BaseMotorController) Start() {
	fmt.Println(mc.name, " started")
}

func (mc *BaseMotorController) Stop() {
	fmt.Println(mc.name, " stopped")
}

func NewL298NMotorController(name string, pin1, pin2, speedPin uint8) *L298NMotorController {
	motor := new(L298NMotorController)
	motor.name = name
	return motor
}

func NewDRV833MotorController(name string, pin1, pin2 uint8, pwmWriter pwm.PwmWriter) *DRV833MotorController {
	motor := new(DRV833MotorController)
	motor.name = name
	return motor
}
