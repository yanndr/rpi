package process

import (
	"fmt"
	"github.com/yanndr/rpi/bdngobot/situation"
	"github.com/yanndr/rpi/event"
	"sync"
)

type DecisionProcess struct {
	baseProcess
	mutex   sync.Mutex
	alerter event.Alerter
}

func NewDecisionProcess(alerter event.Alerter) *DecisionProcess {
	return &DecisionProcess{
		baseProcess: baseProcess{channel: make(chan interface{})},
		alerter:     alerter,
	}
}

func (p *DecisionProcess) Start() {
	go p.eventChannelListener()
	fmt.Println("DecisionProcess process started.")
}

func (p *DecisionProcess) Stop() {
	fmt.Println("DecisionProcess stopped.")
}

func (p *DecisionProcess) eventChannelListener() {
	for value := range p.channel {
		if value == situation.ObstacleFar {
			p.farHandler()
		} else if value == situation.ObstacleMedium {
			p.mediumHandler()
		} else if value == situation.ObstacleClose {
			p.closeHandler()
		} else {
			fmt.Println("No handler for ", value)
		}
	}
}

func (p *DecisionProcess) farHandler() {

}

func (p *DecisionProcess) mediumHandler() {

}

func (p *DecisionProcess) closeHandler() {

}
