package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kidoman/embd"
	"github.com/kidoman/embd/controller/pca9685"
	_ "github.com/kidoman/embd/host/rpi" // This loads the RPi driver
)

func main() {

	pca := pca9685.New(embd.NewI2CBus(0x1), 0x40)
	pca.Freq = 60

	defer pca.Close()

	for i := 0; i < 4000; i += 50 {
		setPwm(i, 100, pca)
	}
	for i := 4000; i > 0; i -= 50 {
		setPwm(i, 100, pca)
	}

}

func setPwm(val, sec int, pca *pca9685.PCA9685) {
	err := pca.SetPwm(0, 0, val)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	time.Sleep(time.Millisecond * time.Duration(sec))
}
