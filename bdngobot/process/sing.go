package process

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
	"time"
)

const duration = time.Second * 10

type SingProcess struct {
	DistanceAlertChannel chan ObstacleDistance
	singing              bool
	rwMutex              sync.RWMutex
	timer                *time.Timer
}

func NewSingProcess() *SingProcess {
	return &SingProcess{
		DistanceAlertChannel: make(chan ObstacleDistance),
		singing:              false,
		timer:                time.NewTimer(duration),
	}
}

func (sp *SingProcess) Start() {
	sp.timer = time.AfterFunc(duration, sp.pickSong)

	go ObstacleChannelListener(sp.DistanceAlertChannel, sp.farHandler, sp.mediumHandler, sp.closeHandler)
}

func (sp *SingProcess) pickSong() {
	sp.sing("/home/pi/mp3/ILBB.mp3")
}

func (sp *SingProcess) sing(song string) {
	if !sp.singing {
		sp.rwMutex.RLock()
		defer sp.rwMutex.RUnlock()
		sp.singing = true
		cmd := exec.Command("omxplayer", song)
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
			fmt.Println(err)
		}
		log.Printf("Waiting for command to finish...")
		err = cmd.Wait()
		log.Printf("Command finished with error: %v", err)
		sp.singing = false
	}
}

func (sp *SingProcess) Stop() {
	sp.timer.Stop()
}

func (sp *SingProcess) farHandler() {

}

func (sp *SingProcess) mediumHandler() {
	sp.timer.Reset(duration)
}

func (sp *SingProcess) closeHandler() {
	sp.timer.Reset(duration)
}
