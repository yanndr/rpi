package mouvement

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/yanndr/rpi/bdngobot/process"
	"github.com/yanndr/rpi/bdngobot/situation"
	"github.com/yanndr/rpi/controller"
)

type MouvmentCommand string

const (
	Stop  MouvmentCommand = "StartMoving"
	Start MouvmentCommand = "StopMoving"
)

var Started = false

const (
	cruiseSpeed = 0.8
	escapeSpeed = 1.0
	backSpeed   = 0.3
)

type MouvementProcess struct {
	process.BaseProcess
	motorsController controller.MotorsController
	mutex            sync.Mutex
}

func NewMouvementProcess(motorsController controller.MotorsController) *MouvementProcess {
	return &MouvementProcess{
		BaseProcess:      process.BaseProcess{Channel: make(chan interface{})},
		motorsController: motorsController,
	}
}

func (mp *MouvementProcess) Start() {
	go mp.eventChannelListener()
	fmt.Println("Mouvment process started.")
}

func (mp *MouvementProcess) Stop() {
	mp.motorsController.Stop()
	fmt.Println("Mouvment process stoped.")
}

func (mp *MouvementProcess) farHandler() {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()
	fmt.Println("Mouvement Far handler")
	mp.moveStraight(cruiseSpeed)
}

func (mp *MouvementProcess) mediumHandler() {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()
	fmt.Println("Mouvement medium handler")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	v := r.Intn(1)

	if v == 0 {
		mp.turnLeft(cruiseSpeed - 0.2)
	} else {
		mp.turnRight(cruiseSpeed - 0.2)
	}
}

func (mp *MouvementProcess) closeHandler() {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()
	fmt.Println("Mouvement Close handler")
	mp.motorsController.Stop()
	mp.moveStraight(-cruiseSpeed / 2)
	time.Sleep(time.Second * 1)
	mp.motorsController.Stop()
	mp.motorsController.RotateRight()
	time.Sleep(time.Second * 1)
	mp.motorsController.Stop()
}

func (mp *MouvementProcess) moveStraight(speed float64) {
	if !mp.motorsController.IsMoving() {
		mp.motorsController.Start()
	}

	mp.motorsController.SetSpeed(speed)
	mp.motorsController.SetBalance(1, 1)
}

func (mp *MouvementProcess) turnLeft(speed float64) {
	mp.motorsController.SetSpeed(speed)
	mp.motorsController.SetBalance(0.6, 1)
}

func (mp *MouvementProcess) turnRight(speed float64) {
	mp.motorsController.SetSpeed(speed)
	mp.motorsController.SetBalance(1, 0.6)
}

// func (mp *MouvementProcess) rotate(speed float64) {
// 	mp.motorsController.SetSpeed(speed)
// 	mp.motorsController.RotateRight()
// 	time.Sleep(time.Second * 1)

// }

func (mp *MouvementProcess) eventChannelListener() {
	for value := range mp.Channel {
		if Started {
			if value == situation.ObstacleFar {
				mp.farHandler()
			} else if value == situation.ObstacleMedium {
				mp.mediumHandler()
			} else if value == situation.ObstacleClose {
				mp.closeHandler()
			} else if value == Stop {
				Started = false
			}
		}
		if value == Start {
			Started = true
		}
	}
}
