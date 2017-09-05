package led

import (
	"time"

	"github.com/yanndr/rpi/pwm"
)

type LedController struct {
	pinNums   []uint8
	channels  map[uint8]chan bool
	ticker    *time.Ticker
	pwmWriter pwm.PwmWriter
	inProc    map[uint8]bool
}

func NewLedController(pwmWriter pwm.PwmWriter, pins ...uint8) *LedController {

	c := &LedController{}
	c.pwmWriter = pwmWriter
	c.pinNums = make([]uint8, len(pins))
	c.channels = make(map[uint8]chan bool)
	c.inProc = make(map[uint8]bool)
	for i, pin := range pins {
		c.pinNums[i] = pin
		c.channels[pin] = make(chan bool)
		// c.pins[i] = rpio.Pin(pin)
		// c.pins[i].Output()
		// c.pins[i].Low()
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

// func (c *LedController) TurnOn(led uint8, speed time.Duration, intensity func(float64) float64) {
// 	ticker := time.NewTicker(speed)

// 	i := 0.0
// 	for {
// 		select {
// 		case <-ticker.C:
// 			val := intensity(i)
// 			if val > 1 {
// 				c.pwmWriter.PwmWrite(led, 1.0)
// 			}
// 			if val < 0 {
// 				c.pwmWriter.PwmWrite(led, 0)
// 			}
// 			fmt.Println("i=", i, " val=", val)
// 			i = i + 0.1
// 			c.pwmWriter.PwmWrite(led, val)
// 			if i >= 1 {
// 				fmt.Println("temrination")
// 				ticker.Stop()
// 				return
// 			}
// 		}
// 	}

// }

func (c *LedController) TurnOn(led uint8) {
	c.pwmWriter.PwmWrite(led, 1)
}

func (c *LedController) Increase(led uint8, d time.Duration) {
	duration := d / 100
	ticker := time.NewTicker(duration)
	val := 0.0
	inc := 0.01
	c.inProc[led] = true
	defer func() { c.inProc[led] = false }()
	for {
		select {
		case <-ticker.C:
			val = val + inc
			if val > 1 {
				c.pwmWriter.PwmWrite(led, 1)
				return
			}
			c.pwmWriter.PwmWrite(led, val)
		case <-c.channels[led]:
			// fmt.Println("Cancel  on led ", led)
			return

		}
	}
}

func (c *LedController) TurnOff(led uint8, d time.Duration) {
	duration := d / 100
	ticker := time.NewTicker(duration)
	val := 1.0
	inc := 0.01
	c.inProc[led] = true
	defer func() { c.inProc[led] = false }()
	for {
		select {
		case <-ticker.C:
			val = val - inc
			if val < 0 {
				c.pwmWriter.PwmWrite(led, 0)
				return
			}
			c.pwmWriter.PwmWrite(led, val)
		case <-c.channels[led]:
			// fmt.Println("Cancel  off led ", led)
			return
		}
	}
}

func (c *LedController) Scanner() {

	for {
		// wg := sync.WaitGroup{}

		for i := 0; i < len(c.pinNums)-1; i++ {
			go func() {
				// wg.Add(1)
				// defer wg.Done()
				if c.inProc[c.pinNums[i]] {
					c.channels[c.pinNums[i]] <- true
				}
				c.TurnOn(c.pinNums[i])
				c.TurnOff(c.pinNums[i], time.Second/2)
			}()
			time.Sleep(time.Second / 6)
		}
		for i := len(c.pinNums) - 1; i >= 0; i-- {
			go func() {
				// wg.Add(1)
				// defer wg.Done()
				if c.inProc[c.pinNums[i]] {
					c.channels[c.pinNums[i]] <- true
				}
				c.TurnOn(c.pinNums[i])
				c.TurnOff(c.pinNums[i], time.Second/2)
			}()
			time.Sleep(time.Second / 6)
		}
		// wg.Wait()
		// fmt.Println("led set value ", value, " to ", pin)
	}

}
