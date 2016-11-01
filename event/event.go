package event

import (
	"fmt"
	"sync"
)

type Alerter interface {
	PostAlert(data interface{})
}

type Suscriber interface {
	Subscribe(name string, channel chan<- interface{})
	Unsubscribe(name string)
}

type EventDispatcher struct {
	mutex         sync.RWMutex
	eventChannels map[string]chan<- interface{}
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		eventChannels: make(map[string]chan<- interface{}),
	}
}

func (e *EventDispatcher) PostAlert(data interface{}) {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	fmt.Println("Post Alert ", data)
	for _, outputChan := range e.eventChannels {
		outputChan <- data
	}
}

func (e *EventDispatcher) CloseChannels() {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	for _, outputChan := range e.eventChannels {
		close(outputChan)
	}
}

func (e *EventDispatcher) Subscribe(name string, channel chan<- interface{}) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.eventChannels[name] = channel
}

func (e *EventDispatcher) Unsubscribe(name string) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	delete(e.eventChannels, name)
}
