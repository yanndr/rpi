package tts

import (
	"sync"

	"github.com/gokyle/gofestival"
)

type Speaker interface {
	Speak(text string)
}

type Festival struct {
	mutex    sync.Mutex
	speaking bool
}

func (f *Festival) Speak(text string) {
	if !f.speaking {
		f.mutex.Lock()
		defer f.mutex.Unlock()
		f.speaking = true
		festival.Speak(text)
		f.speaking = false
	}
}
