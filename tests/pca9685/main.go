package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/kidoman/embd"
	"github.com/kidoman/embd/controller/pca9685"
	_ "github.com/kidoman/embd/host/rpi" // This loads the RPi driver
)

func main() {

	pca := pca9685.New(embd.NewI2CBus(0x1), 0x40)
	pca.Freq = 50

	defer pca.Close()

	var val int64
	var err error
	var chanel int64
	if len(os.Args) < 3 {
		val = 50
		chanel = 1
	} else {
		val, err = strconv.ParseInt(os.Args[2], 10, 32)
		if err != nil {
			fmt.Println("Error after val parse")
			fmt.Println(err)
			os.Exit(-1)
		}

		chanel, err = strconv.ParseInt(os.Args[1], 10, 32)
		if err != nil {
			fmt.Println("Error at Channel pars")
			fmt.Println(err)
			os.Exit(-1)
		}
	}
	err = pca.SetPwm(int(chanel), 0, int(val))

	if err != nil {
		fmt.Println("Error after SetPwm")
		fmt.Println(err)
		os.Exit(-1)
	}

	time.Sleep(time.Second * 5)

	// for i := 0; i < 4000; i += 50 {
	// 	setPwm(i, 100, pca)
	// }
	// for i := 4000; i > 0; i -= 50 {
	// 	setPwm(i, 100, pca)
	// }

}

func setPwm(val, sec int, pca *pca9685.PCA9685) {
	err := pca.SetPwm(0, 0, val)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	time.Sleep(time.Millisecond * time.Duration(sec))
}
