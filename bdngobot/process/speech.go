package process

import (
	"github.com/yanndr/rpi/bdngobot/text"
	"github.com/yanndr/rpi/tts"
)

type SpeechProcess struct {
	baseProcess
	speaker tts.Speaker
	textGen text.TextGenerator
}

func NewSpeechProcess(speaker tts.Speaker, tg text.TextGenerator) *SpeechProcess {
	return &SpeechProcess{
		baseProcess: baseProcess{channel: make(chan interface{})},
		speaker:     speaker,
		textGen:     tg,
	}
}

func (sp *SpeechProcess) Start() {
	go ObstacleChannelListener(sp.channel, sp.farHandler, sp.mediumHandler, sp.closeHandler)
}

func (sp *SpeechProcess) Stop() {
	go sp.speaker.Speak("Bye!")
}

func (sp *SpeechProcess) farHandler() {
	go sp.speaker.Speak(sp.textGen.Text(text.Neutral, text.Any))
}

func (sp *SpeechProcess) mediumHandler() {
	go sp.speaker.Speak("Something is in the way.")
}

func (sp *SpeechProcess) closeHandler() {
	go sp.speaker.Speak(sp.textGen.Text(text.Neutral, text.Any))
}
