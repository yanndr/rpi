package main

import (
	"fmt"

	"github.com/kidoman/embd"
	"github.com/kidoman/embd/controller/pca9685"
	"github.com/yanndr/rpi/controller/led"
	"github.com/yanndr/rpi/pwm"
)

func main() {
	pca := pca9685.New(embd.NewI2CBus(0x1), 0x40)
	pca.Freq = 60

	defer pca.Close()

	pca9685 := pwm.NewPca9685(pca)

	lc := led.NewLedController(pca9685, 4, 5, 6, 7, 8, 9, 10, 11)
	//lc.BlinkAll(time.Second/8, 0, 0.8, 0.1)
	lc.SetAllOff()
	// lc.TurnOn(4, time.Second)
	// lc.TurnOff(4, time.Second*2)

	lc.Scanner()
	var response int
	for response != 'q' {
		fmt.Scanf("%c", &response)

		switch response {
		default:
			fmt.Println("Pouet")
		case 'q':
			fmt.Println("Killing ")
		}
	}

	lc.SetAllOff()
}
