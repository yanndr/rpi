package pwm

import (
	"sync"

	"github.com/kidoman/embd/controller/pca9685"
	_ "github.com/kidoman/embd/host/rpi" // This loads the RPi driver
)

func NewPca9685(pca *pca9685.PCA9685) *Pca9685 {
	return &Pca9685{pca: pca}
}

type Pca9685 struct {
	pca   *pca9685.PCA9685
	mutex sync.Mutex
}

func (p *Pca9685) PwmWrite(pin uint8, val float64) (err error) {
	// p.mutex.Lock()
	// defer p.mutex.Unlock()

	off := 4095 - int(2047*(1-val)) //1 -> 4095; 0.5 -> 3071
	on := 2047 - int(2047*val)      // 1 -> 0;  0.5 ->
	// fmt.Println("val:", val, " -> set chan", pin, " to on:", on, " off:", off)
	return p.pca.SetPwm(int(pin), on, off)

}
