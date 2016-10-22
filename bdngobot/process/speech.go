package process

import (
	"fmt"
	"github.com/yanndr/rpi/bdngobot/mood"
	"github.com/yanndr/rpi/bdngobot/situation"
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
	fmt.Println("Speech process started.")
}

func (sp *SpeechProcess) Stop() {
	go sp.speaker.Speak(sp.textGen.Text(mood.Neutral, situation.Any))
	fmt.Println("Speech process stopped.")
}

func (sp *SpeechProcess) farHandler() {
	go sp.speaker.Speak(sp.textGen.Text(mood.Neutral, situation.Any))
}

func (sp *SpeechProcess) mediumHandler() {
	go sp.speaker.Speak(sp.textGen.Text(mood.Neutral, situation.ObstacleMedium))
}

func (sp *SpeechProcess) closeHandler() {
	go sp.speaker.Speak(sp.textGen.Text(mood.Neutral, situation.ObstacleClose))
}
