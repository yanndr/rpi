package process

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/yanndr/bdngobot/controller"
)

const (
	cruiseSpeed = 0.7
	escapeSpeed = 1.0
	backSpeed   = 0.3
)

type MouvementProcess struct {
	baseProcess
	motorsController controller.MotorsController
	mutex            sync.Mutex
}

func NewMouvementProcess(motorsController controller.MotorsController) *MouvementProcess {
	return &MouvementProcess{
		baseProcess:      baseProcess{channel: make(chan interface{})},
		motorsController: motorsController,
	}
}

func (mp *MouvementProcess) Start() {
	go ObstacleChannelListener(mp.channel, mp.farHandler, mp.mediumHandler, mp.closeHandler)
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
		mp.turnLeft(cruiseSpeed / 2)
	} else {
		mp.turnRight(cruiseSpeed / 2)
	}
}

func (mp *MouvementProcess) closeHandler() {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()
	fmt.Println("Mouvement Close handler")
	mp.motorsController.Stop()
	mp.moveStraight(-cruiseSpeed / 2)
	time.Sleep(time.Second * 1)
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
