package process

import (
	"github.com/yanndr/rpi/tts"
)

type SpeechProcess struct {
	baseProcess
	speaker tts.Speaker
}

func NewSpeechProcess(speaker tts.Speaker) *SpeechProcess {
	return &SpeechProcess{
		baseProcess: baseProcess{channel: make(chan interface{})},
		speaker:     speaker,
	}
}

func (sp *SpeechProcess) Start() {
	go ObstacleChannelListener(sp.channel, sp.farHandler, sp.mediumHandler, sp.closeHandler)
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
