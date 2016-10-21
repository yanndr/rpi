package process

import (
	"sync"

	"github.com/gokyle/gofestival"
)

type SpeechProcess struct {
	DistanceAlertChannel chan ObstacleDistance
	speaking             bool
	rwMutex              sync.RWMutex
}

func NewSpeechProcess() *SpeechProcess {
	return &SpeechProcess{
		DistanceAlertChannel: make(chan ObstacleDistance),
		speaking:             false,
	}
}

func (sp *SpeechProcess) Start() {
	go ObstacleChannelListener(sp.DistanceAlertChannel, sp.farHandler, sp.mediumHandler, sp.closeHandler)
}

func (sp *SpeechProcess) Speak(text string) {
	if !sp.speaking {
		sp.rwMutex.RLock()
		defer sp.rwMutex.RUnlock()
		sp.speaking = true
		festival.Speak(text)
		sp.speaking = false
	}
}

func (sp *SpeechProcess) Stop() {
}

func (sp *SpeechProcess) farHandler() {
	go sp.Speak("Yay! everything is fine!")
}

func (sp *SpeechProcess) mediumHandler() {
	go sp.Speak("Something is in the way.")
}

func (sp *SpeechProcess) closeHandler() {
	go sp.Speak("Ho no I am stuck")
}
