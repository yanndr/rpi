package pwm

import (
	"fmt"
	"os"
)

type PiBlaster struct{

}

func (*PiBlaster) PwmWrite(pin uint8, val float64) (err error) {
	return piBlaster(fmt.Sprintf("%v=%v\n", pin, val))
}

func piBlaster(data string) (err error) {
	fi, err := os.OpenFile("/dev/pi-blaster", os.O_WRONLY|os.O_APPEND, 0644)
	defer fi.Close()

	if err != nil {
		return err
	}

	_, err = fi.WriteString(data)
	return
}
