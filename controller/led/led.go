package led

import (
	// "fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
	"github.com/yanndr/rpi/pwm"
)

type LedController struct {
	pinNums   []uint8
	pins      []rpio.Pin
	ticker    *time.Ticker
	pwmWriter pwm.PwmWriter
}

func NewLedController(pwmWriter pwm.PwmWriter, pins ...uint8) *LedController {

	c := &LedController{}
	c.pwmWriter = pwmWriter
	c.pinNums = make([]uint8, len(pins))
	c.pins = make([]rpio.Pin, len(pins))
	for i, pin := range pins {
		c.pinNums[i] = pin
		c.pins[i] = rpio.Pin(pin)
		c.pins[i].Output()
		c.pins[i].Low()
	}

	return c
}

func (c *LedController) SetAllValue(value float64) {
	for _, pin := range c.pinNums {
		c.pwmWriter.PwmWrite(pin, value)
		// fmt.Println("led set value ", value, " to ", pin)
	}
}

func (c *LedController) SetAllOn() {
	c.SetAllValue(1)
}

func (c *LedController) SetAllOff() {
	c.SetAllValue(0)
}

func (c *LedController) BlinkAll(d time.Duration, min, max, inc float64) {

	c.ticker = time.NewTicker(d)
	go func() {
		i := min
		increase := true
		for range c.ticker.C {
			if i >= max && increase {
				i = max
				increase = false
			} else if i <= min && !increase {
				i = min
				increase = true
			}

			if i < max && increase {
				i += inc
			} else if i > min && !increase {
				i -= inc
			}
			if i < min {
				i = min
			} else if i > max {
				i = max
			}

			c.SetAllValue(i)
		}
	}()

}
