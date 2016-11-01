package decision

import (
	"fmt"
	"github.com/yanndr/rpi/bdngobot/process"
	"github.com/yanndr/rpi/bdngobot/process/mouvement"
	"github.com/yanndr/rpi/bdngobot/process/speech"
	"github.com/yanndr/rpi/bdngobot/situation"
	"github.com/yanndr/rpi/event"
	"sync"
	"time"
)

const duration = time.Second * 15

type DecisionProcess struct {
	process.BaseProcess
	mutex   sync.Mutex
	alerter event.Alerter
	timer   *time.Timer
}

func NewDecisionProcess(alerter event.Alerter) *DecisionProcess {
	return &DecisionProcess{
		BaseProcess: process.BaseProcess{Channel: make(chan interface{})},
		alerter:     alerter,
		timer:       time.NewTimer(duration),
	}
}

func (p *DecisionProcess) Start() {
	go p.eventChannelListener()
	fmt.Println("DecisionProcess process started.")
	// p.alerter.PostAlert(mouvement.Start)
	p.alerter.PostAlert(speech.Unmute)

	p.timer = time.AfterFunc(duration, func() {
		p.alerter.PostAlert(mouvement.Start)
		time.Sleep(time.Second * 7)
		p.alerter.PostAlert(mouvement.Stop)
	})
}

func (p *DecisionProcess) Stop() {
	fmt.Println("DecisionProcess stopped.")
}

func (p *DecisionProcess) eventChannelListener() {
	for value := range p.Channel {
		if value == situation.ObstacleFar {
			p.farHandler()
		} else if value == situation.ObstacleMedium {
			p.mediumHandler()
		} else if value == situation.ObstacleClose {
			p.closeHandler()
		}
	}
}

func (p *DecisionProcess) farHandler() {
	p.timer.Reset(duration)
}

func (p *DecisionProcess) mediumHandler() {
	p.timer.Reset(duration)
}

func (p *DecisionProcess) closeHandler() {

}
