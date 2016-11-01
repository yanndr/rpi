package speech

import (
	"fmt"
	"github.com/yanndr/rpi/bdngobot/mood"
	"github.com/yanndr/rpi/bdngobot/process"
	"github.com/yanndr/rpi/bdngobot/process/mouvement"
	"github.com/yanndr/rpi/bdngobot/situation"
	"github.com/yanndr/rpi/bdngobot/text"
	"github.com/yanndr/rpi/tts"
)

type speechCommand string

var Muted = true

const (
	Mute   speechCommand = "Mute"
	Unmute speechCommand = "UnMute"
)

type SpeechProcess struct {
	process.BaseProcess
	speaker tts.Speaker
	textGen text.TextGenerator
}

func NewSpeechProcess(speaker tts.Speaker, tg text.TextGenerator) *SpeechProcess {
	return &SpeechProcess{
		BaseProcess: process.BaseProcess{Channel: make(chan interface{})},
		speaker:     speaker,
		textGen:     tg,
	}
}

func (sp *SpeechProcess) Start() {
	go sp.eventListener()
	fmt.Println("Speech process started.")
}

func (sp *SpeechProcess) Stop() {
	go sp.speaker.Speak(sp.textGen.Text(mood.Neutral, situation.Any))
	fmt.Println("Speech process stopped.")
}

func (sp *SpeechProcess) farHandler() {
	if mouvement.Started {
		go sp.speaker.Speak(sp.textGen.Text(mood.Neutral, situation.Any))
	} else {
		go sp.speaker.Speak(sp.textGen.Text(mood.Neutral, situation.MovingFar))
	}
}

func (sp *SpeechProcess) mediumHandler() {
	if mouvement.Started {
		go sp.speaker.Speak(sp.textGen.Text(mood.Neutral, situation.ObstacleMedium))
	} else {
		go sp.speaker.Speak(sp.textGen.Text(mood.Neutral, situation.MovingMedium))
	}
}

func (sp *SpeechProcess) closeHandler() {
	if mouvement.Started {
		go sp.speaker.Speak(sp.textGen.Text(mood.Neutral, situation.ObstacleClose))
	} else {
		go sp.speaker.Speak(sp.textGen.Text(mood.Neutral, situation.MovingClose))
	}
}

func (sp *SpeechProcess) eventListener() {
	for value := range sp.Channel {
		if !Muted {
			if value == situation.ObstacleFar {
				sp.farHandler()
			} else if value == situation.ObstacleMedium {
				sp.mediumHandler()
			} else if value == situation.ObstacleClose {
				sp.closeHandler()
			} else if value == Mute {
				Muted = true
			}
		}
		if value == Unmute {
			Muted = false
		}
	}
}
