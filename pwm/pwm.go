package pwm

type PwmWriter interface {
	PwmWrite(pin uint8, val float64) (err error)
}
