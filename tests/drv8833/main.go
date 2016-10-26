package main

import (
	"fmt"
	"github.com/yanndr/rpi/controller"
	"github.com/yanndr/rpi/gpio"
	"os"
	"time"
)

func main() {

	// Open and map memory to access gpio, check for errors
	err := gpio.Open()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c := controller.NewDRV833MotorController("test", 23, 24)

	c.SetSpeed(-1)
	time.Sleep(time.Second * 4)
	c.SetSpeed(1)
	time.Sleep(time.Second * 4)
	c.SetSpeed(-.5)
	time.Sleep(time.Second * 4)
	c.SetSpeed(.5)
	time.Sleep(time.Second * 4)
	c.Stop()
}
