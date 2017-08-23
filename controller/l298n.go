// +build linux,arm

package controller

import (
	"fmt"

	"github.com/stianeikeland/go-rpio"
	"github.com/yanndr/rpi/pwm"
)

//L298NMotorController single motor implementation (L298n)
type L298NMotorController struct {
	name      string
	pin1      rpio.Pin
	pin2      rpio.Pin
	speedPin  uint8
	speed     float64
	pwmwriter pwm.PwmWriter
}

//NewL298NMotorController retrun a Motor instance
func NewL298NMotorController(name string, pin1, pin2, speedPin uint8) *L298NMotorController {
	motor := new(L298NMotorController)
	motor.name = name
	motor.pin1 = rpio.Pin(pin1)
	motor.pin2 = rpio.Pin(pin2)
	motor.pin1.Output()
	motor.pin2.Output()
	motor.pin1.Low()
	motor.pin2.Low()
	motor.speedPin = speedPin
	return motor
}

//SetSpeed Set the speed of the motor
func (motor *L298NMotorController) SetSpeed(speed float64) {
	if speed < -1 {
		speed = -1
	} else if speed > 1 {
		speed = 1
	}

	if speed > 0 {
		motor.runForward()
	} else if speed < 0 {
		motor.runBackward()
		speed = speed * -1
	}

	fmt.Printf("%v speed:%v \n", motor.name, speed)

	motor.speed = speed
	motor.pwmwriter.PwmWrite(motor.speedPin, speed)
}

func (motor *L298NMotorController) runForward() {
	motor.pin1.Low()
	motor.pin2.High()
}

func (motor *L298NMotorController) runBackward() {
	motor.pin1.High()
	motor.pin2.Low()
}

//Stop stop the motor
func (motor *L298NMotorController) Stop() {
	motor.pwmwriter.PwmWrite(motor.speedPin, 0)
}

//Start start the motor
func (motor *L298NMotorController) Start() {
	motor.pwmwriter.PwmWrite(motor.speedPin, motor.speed)
}
