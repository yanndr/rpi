package process

import (
	"github.com/yanndr/rpi/tts"
)

type SpeechProcess struct {
	DistanceAlertChannel chan ObstacleDistance
	speaker              tts.Speaker
}

func NewSpeechProcess(speaker tts.Speaker) *SpeechProcess {
	return &SpeechProcess{
		DistanceAlertChannel: make(chan ObstacleDistance),
		speaker:              speaker,
	}
}

func (sp *SpeechProcess) Start() {
	go ObstacleChannelListener(sp.DistanceAlertChannel, sp.farHandler, sp.mediumHandler, sp.closeHandler)
}

func (sp *SpeechProcess) Stop() {
	go sp.speaker.Speak("Bye!")
}

func (sp *SpeechProcess) farHandler() {
	go sp.speaker.Speak("Yay! everything is fine!")
}

func (sp *SpeechProcess) mediumHandler() {
	go sp.speaker.Speak("Something is in the way.")
}

func (sp *SpeechProcess) closeHandler() {
	go sp.speaker.Speak("Ho no I am stuck")
}
