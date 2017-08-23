// +build linux,arm

package controller

import (

	// "fmt"

	"log"

	"github.com/yanndr/rpi/pwm"
)

//DRV833MotorController single motor implementation (DRV8833)
type DRV833MotorController struct {
	name             string
	pinNum1, pinNum2 uint8
	speed            float64
	pwmWriter        pwm.PwmWriter
}

//NewDRV833MotorController retrun a Motor instance
func NewDRV833MotorController(name string, pin1, pin2 uint8, pwmWriter pwm.PwmWriter) *DRV833MotorController {
	motor := new(DRV833MotorController)
	motor.name = name
	motor.pinNum1 = pin1
	motor.pinNum2 = pin2
	motor.pwmWriter = pwmWriter
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
	err := motor.pwmWriter.PwmWrite(motor.pinNum1, 0)
	if err != nil {
		log.Println(err)
		return
	}
	err = motor.pwmWriter.PwmWrite(motor.pinNum2, speed)
	if err != nil {
		log.Println(err)
	}
}

func (motor *DRV833MotorController) runBackward(speed float64) {
	err := motor.pwmWriter.PwmWrite(motor.pinNum2, 0)
	if err != nil {
		log.Println(err)
		return
	}
	err = motor.pwmWriter.PwmWrite(motor.pinNum1, speed)
	if err != nil {
		log.Println(err)
	}
}

//Stop stop the motor
func (motor *DRV833MotorController) Stop() {
	err := motor.pwmWriter.PwmWrite(motor.pinNum1, 0)
	if err != nil {
		log.Println(err)
		return
	}
	err = motor.pwmWriter.PwmWrite(motor.pinNum2, 0)
	if err != nil {
		log.Println(err)
	}
}

//Start start the motor
func (motor *DRV833MotorController) Start() {
	//pwm.PwmWrite(motor.speedPin, motor.speed)
}
