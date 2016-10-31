package led

import (
	"github.com/stianeikeland/go-rpio"
	"github.com/yanndr/rpi/pwm"
)

type LedController struct {
	pinNums []uint8
	pins    []rpio.Pin
}

func NewLedController(pins ...uint8) *LedController {

	c := &LedController{}

	i := 0
	for _, pin := range pins {
		c.pinNums[i]
		c.pins[i] = rpio.Pin(pin)
		c.pins[i].Output()
		c.pins[i].Low()
	}

	return c
}

func (c *LedController) SetAllValue(value float64) {
	for _, pin := range c.pinNums {
		pwm.PwmWrite(pin, value)
	}
}

func (c *LedController) SetAllOn() {
	c.SetAllValue(1)
}

func (c *LedController) SetAllOff() {
	c.SetAllValue(0)
}

func (c *LedController) Blink() {

}
