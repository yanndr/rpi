package process

import (
	"fmt"
	"github.com/yanndr/rpi/media"
	"time"
)

const duration = time.Second * 20

type PlayerProcess struct {
	BaseProcess
	timer  *time.Timer
	player media.Player
}

func NewPlayerProcess(player media.Player) *PlayerProcess {
	return &PlayerProcess{
		BaseProcess: BaseProcess{Channel: make(chan interface{})},
		timer:       time.NewTimer(duration),
		player:      player,
	}
}

func (sp *PlayerProcess) Start() {
	sp.timer = time.AfterFunc(duration, func() { sp.player.Play(sp.pickFile()) })

	go ObstacleChannelListener(sp.Channel, sp.farHandler, sp.mediumHandler, sp.closeHandler)
	fmt.Println("Player process started.")
}

func (sp *PlayerProcess) pickFile() string {
	return "/home/pi/Music/halloween.mp3"
}

func (sp *PlayerProcess) Stop() {
	sp.timer.Stop()
	fmt.Println("Player process stopped.")
}

func (sp *PlayerProcess) farHandler() {

}

func (sp *PlayerProcess) mediumHandler() {
	sp.timer.Reset(duration)
}

func (sp *PlayerProcess) closeHandler() {
	sp.timer.Reset(duration)
}
