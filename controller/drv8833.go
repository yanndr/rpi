// +build linux,arm

package controller

import (
	// "fmt"

	"github.com/stianeikeland/go-rpio"
	"github.com/yanndr/rpi/pwm"
)

//DRV833MotorController single motor implementation (DRV8833)
type DRV833MotorController struct {
	name             string
	pinNum1, pinNum2 uint8
	pin1             rpio.Pin
	pin2             rpio.Pin
	speed            float64
}

//NewDRV833MotorController retrun a Motor instance
func NewDRV833MotorController(name string, pin1, pin2 uint8) *DRV833MotorController {
	motor := new(DRV833MotorController)
	motor.name = name
	motor.pinNum1 = pin1
	motor.pinNum2 = pin2
	motor.pin1 = rpio.Pin(pin1)
	motor.pin2 = rpio.Pin(pin2)
	motor.pin1.Output()
	motor.pin2.Output()
	motor.pin1.Low()
	motor.pin2.Low()
	return motor
}

//SetSpeed Set the speed of the motor
func (motor *DRV833MotorController) SetSpeed(speed float64) {
	// fmt.Println("Set Speed")
	if speed < -1 {
		speed = -1
	} else if speed > 1 {
		speed = 1
	}

	if speed > 0 {
		motor.runForward(speed)
		// fmt.Println(motor.name, " Forward ", motor.pinNum1, " ", motor.pinNum2)
	} else if speed < 0 {
		speed = speed * -1
		motor.runBackward(speed)
		// fmt.Println("Backward")
	}

	// fmt.Printf("%v speed:%v \n", motor.name, speed)

	motor.speed = speed

}

func (motor *DRV833MotorController) runForward(speed float64) {
	pwm.PwmWrite(motor.pinNum1, 0)
	pwm.PwmWrite(motor.pinNum2, speed)
}

func (motor *DRV833MotorController) runBackward(speed float64) {
	pwm.PwmWrite(motor.pinNum2, 0)
	pwm.PwmWrite(motor.pinNum1, speed)
}

//Stop stop the motor
func (motor *DRV833MotorController) Stop() {
	pwm.PwmWrite(motor.pinNum1, 0)
	pwm.PwmWrite(motor.pinNum2, 0)
}

//Start start the motor
func (motor *DRV833MotorController) Start() {
	//pwm.PwmWrite(motor.speedPin, motor.speed)
}
