package led

import (
	"time"

	"github.com/stianeikeland/go-rpio"
	"github.com/yanndr/rpi/pwm"
)

type LedController struct {
	pinNums []uint8
	pins    []rpio.Pin
	ticker  *time.Ticker
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

func (c *LedController) BlinkAll(d time.Duration, min,max,inc float64) {

	c.ticker = time.NewTicker(d)
	go func() {
		i := min
	increase:
		true
		for range c.ticker.C {
			if i == max && increase{
				increase = false
			}else i==min && !increase{
				increase = true
			}

			if i < max && increase {
				i += inc
			} else if i > min && !increase {
				i -= inc
			}
			c.SetAllValue(i)
		}
	}()

}
