package decision

import (
	"fmt"
	"github.com/yanndr/rpi/bdngobot/process"
	"github.com/yanndr/rpi/bdngobot/process/mouvement"
	"github.com/yanndr/rpi/bdngobot/process/speech"
	"github.com/yanndr/rpi/bdngobot/situation"
	"github.com/yanndr/rpi/event"
	"sync"
)

type DecisionProcess struct {
	process.BaseProcess
	mutex   sync.Mutex
	alerter event.Alerter
}

func NewDecisionProcess(alerter event.Alerter) *DecisionProcess {
	return &DecisionProcess{
		BaseProcess: process.BaseProcess{Channel: make(chan interface{})},
		alerter:     alerter,
	}
}

func (p *DecisionProcess) Start() {
	go p.eventChannelListener()
	fmt.Println("DecisionProcess process started.")
	p.alerter.PostAlert(mouvement.Start)
	p.alerter.PostAlert(speech.Unmute)
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

}

func (p *DecisionProcess) mediumHandler() {

}

func (p *DecisionProcess) closeHandler() {

}
