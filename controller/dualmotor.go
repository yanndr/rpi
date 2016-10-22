package controller

import (
	"errors"
	"math"
)

//MotorsController interface to control a motor
type MotorsController interface {
	SetSpeed(speed float64)
	SetBalance(left, right float64)
	Start()
	Stop()
	RotateRight()
	RotateLeft()
	IsMoving() bool
}

//DualMotorController MotorController for 2 motors
type DualMotorController struct {
	leftMotor       SingleMotorController
	rightMotor      SingleMotorController
	speed           float64
	leftBalance     float64
	rightBalance    float64
	leftCorrection  float64
	rightCorrection float64
	maxSpeed        float64
	minSpeed        float64
}

//NewDualMotorController Return an instance of a dualMotor controller
func NewDualMotorController(leftMotor, rightMotor SingleMotorController, leftCorrection, rightCorrection, maxSpeed, minSpeed float64) (*DualMotorController, error) {
	dualMotor := new(DualMotorController)
	dualMotor.leftMotor = leftMotor
	dualMotor.rightMotor = rightMotor
	dualMotor.speed = 0

	if maxSpeed > 1 || maxSpeed <= 0 {
		return nil, errors.New("Error max speed should be between 0.1 and 1")
	}

	if minSpeed >= 1 || minSpeed <= 0 {
		return nil, errors.New("Error min speed should be between 0.1 and 0.99")
	}

	if minSpeed > maxSpeed {
		return nil, errors.New("Error min speed can't be greater than max speed")
	}

	if leftCorrection > 0.99 || leftCorrection < 0 {
		return nil, errors.New("Error left correction should be between 0 and 0.99")
	}

	if rightCorrection > 0.99 || rightCorrection < 0 {
		return nil, errors.New("Error right correction should be between 0 and 0.99")
	}

	dualMotor.maxSpeed = maxSpeed
	dualMotor.minSpeed = minSpeed
	dualMotor.leftCorrection = 1 - leftCorrection
	dualMotor.rightCorrection = 1 - rightCorrection

	return dualMotor, nil
}

//SetSpeed Set the speed of the motors
func (dualMotor *DualMotorController) SetSpeed(speed float64) {

	if speed < -1 {
		speed = -1
	} else if speed > 1 {
		speed = 1
	} else if speed == 0 {
		dualMotor.rightMotor.SetSpeed(0)
		dualMotor.leftMotor.SetSpeed(0)
		return
	}

	dualMotor.speed = speed
	dualMotor.applySpeed()

}

func (dualMotor *DualMotorController) applySpeed() {

	maxSpeed := dualMotor.maxSpeed

	if math.Abs(dualMotor.speed) < dualMotor.maxSpeed {
		maxSpeed = 1
	} else {
		maxSpeed = maxSpeed / math.Abs(dualMotor.speed)
	}

	rightSpeed := dualMotor.speed * dualMotor.rightBalance * dualMotor.rightCorrection * maxSpeed
	leftSpeed := dualMotor.speed * dualMotor.leftBalance * dualMotor.leftCorrection * maxSpeed

	if rightSpeed > 0 {
		if math.Abs(rightSpeed) < dualMotor.minSpeed {
			rightSpeed = dualMotor.minSpeed
		}

		if math.Abs(leftSpeed) < dualMotor.minSpeed {
			leftSpeed = dualMotor.minSpeed
		}
	} else {
		if math.Abs(rightSpeed) > -dualMotor.minSpeed {
			rightSpeed = -dualMotor.minSpeed
		}

		if math.Abs(leftSpeed) > -dualMotor.minSpeed {
			leftSpeed = -dualMotor.minSpeed
		}
	}

	dualMotor.rightMotor.SetSpeed(rightSpeed)
	dualMotor.leftMotor.SetSpeed(leftSpeed)
}

//SetBalance set the balance of the power of the motor
func (dualMotor *DualMotorController) SetBalance(left, right float64) {

	if left < 0 {
		left = 0
	} else if left > 1 {
		left = 1
	}

	if right < 0 {
		right = 0
	} else if left > 1 {
		right = 1
	}

	dualMotor.rightBalance = right
	dualMotor.leftBalance = left

	dualMotor.applySpeed()
}

//Stop stop the motors
func (dualMotor *DualMotorController) Stop() {
	dualMotor.leftMotor.Stop()
	dualMotor.rightMotor.Stop()
	dualMotor.SetSpeed(0)
}

//Start stop the motors
func (dualMotor *DualMotorController) Start() {
	dualMotor.SetSpeed(dualMotor.speed)
}

//RotateRight Run both motors to make a quick rotation to the right
func (dualMotor *DualMotorController) RotateRight() {
	dualMotor.leftMotor.SetSpeed(dualMotor.speed)
	dualMotor.rightMotor.SetSpeed(-dualMotor.speed)
}

//RotateLeft Run both motors to make a quick rotation to the left
func (dualMotor *DualMotorController) RotateLeft() {
	dualMotor.leftMotor.SetSpeed(-dualMotor.speed)
	dualMotor.rightMotor.SetSpeed(dualMotor.speed)
}

//IsMoving return true if the bot is moving
func (dualMotor *DualMotorController) IsMoving() bool {
	return dualMotor.speed != 0
}
