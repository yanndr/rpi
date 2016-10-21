package process

import (
	"math/rand"
	"sync"
	"time"

	"github.com/yanndr/bdngobot/controller"
)

const (
	cruiseSpeed = 0.5
	escapeSpeed = 1.0
	backSpeed   = 0.3
)

type MouvementProcess struct {
	DistanceAlertChannel chan ObstacleDistance
	motorsController     controller.MotorsController
	mutex                sync.Mutex
}

func NewMouvementProcess(motorsController controller.MotorsController) *MouvementProcess {
	return &MouvementProcess{
		DistanceAlertChannel: make(chan ObstacleDistance),
		motorsController:     motorsController,
	}
}

func (mp *MouvementProcess) Start() {
	go ObstacleChannelListener(mp.DistanceAlertChannel, mp.farHandler, mp.mediumHandler, mp.closeHandler)
}

func (mp *MouvementProcess) Stop() {
}

func (mp *MouvementProcess) farHandler() {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()
	mp.moveStraight(cruiseSpeed)
}

func (mp *MouvementProcess) mediumHandler() {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()
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
