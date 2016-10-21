package process

import (
	"github.com/yanndr/rpi/media"
	"time"
)

const duration = time.Second * 10

type PlayerProcess struct {
	DistanceAlertChannel chan ObstacleDistance
	timer                *time.Timer
	player               media.Player
}

func NewPlayerProcess(player media.Player) *PlayerProcess {
	return &PlayerProcess{
		DistanceAlertChannel: make(chan ObstacleDistance),
		timer:                time.NewTimer(duration),
		player:               player,
	}
}

func (sp *PlayerProcess) Start() {
	sp.timer = time.AfterFunc(duration, func() { sp.player.Play(sp.pickFile()) })

	go ObstacleChannelListener(sp.DistanceAlertChannel, sp.farHandler, sp.mediumHandler, sp.closeHandler)
}

func (sp *PlayerProcess) pickFile() string {
	return "/home/pi/mp3/ILBB.mp3"
}

func (sp *PlayerProcess) Stop() {
	sp.timer.Stop()
}

func (sp *PlayerProcess) farHandler() {

}

func (sp *PlayerProcess) mediumHandler() {
	sp.timer.Reset(duration)
}

func (sp *PlayerProcess) closeHandler() {
	sp.timer.Reset(duration)
}
