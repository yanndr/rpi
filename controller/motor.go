package controller

//SingleMotorController Controller for one motor
type SingleMotorController interface {
	SetSpeed(speed float64)
	Start()
	Stop()
}
