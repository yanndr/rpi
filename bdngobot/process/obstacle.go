package process

import (
	"fmt"
	"sync"
	"time"

	"github.com/yanndr/bdngobot/sensor"
)

//ObstacleDistance distance type for the alertChannel
type ObstacleDistance int

//enum for ObstacleDistance
const (
	None ObstacleDistance = iota
	Far
	Medium
	Close
)

//ObstacleDetectorProcess process that check the obstacle with an ultrasoundSensor
type ObstacleDetectorProcess struct {
	ultrasoundSensor sensor.DistanceSensor
	alertDistance    ObstacleDistance
	alertChannels    map[string]chan<- ObstacleDistance
	warningDistance  ObstacleDistance
	lastAlert        ObstacleDistance
	ticker           *time.Ticker
}

func NewObstacleDetectorProcess(ultrasoundSensor sensor.DistanceSensor, alertDistance, warningDistance ObstacleDistance) *ObstacleDetectorProcess {

	odp := &ObstacleDetectorProcess{
		ultrasoundSensor: ultrasoundSensor,
		//alertChannel:     alertChannel,
		alertDistance:   alertDistance,
		warningDistance: warningDistance,
		lastAlert:       None,
		alertChannels:   make(map[string]chan<- ObstacleDistance),
	}

	return odp
}

func (odp *ObstacleDetectorProcess) Start() {

	odp.ticker = time.NewTicker(time.Second / 4)
	go func() {
		for range odp.ticker.C {

			d, err := odp.ultrasoundSensor.Distance()

			if err != nil {
				fmt.Println(err)
			}

			if float64(odp.alertDistance) > d {
				if odp.lastAlert != Close {
					odp.postAlert(Close)
					odp.lastAlert = Close
				}
			} else if float64(odp.warningDistance) > d {
				if odp.lastAlert != Medium {
					odp.postAlert(Medium)
					odp.lastAlert = Medium
				}
			} else {
				if odp.lastAlert != Far {
					odp.postAlert(Far)
					odp.lastAlert = Far
				}
			}
		}
		fmt.Println("Obstacle detector process exited")
	}()
	return
}

func (odp *ObstacleDetectorProcess) Stop() {
	odp.ticker.Stop()
	time.Sleep(time.Second)
	odp.closeChannels()
}

var rwMutex sync.RWMutex

func (odp *ObstacleDetectorProcess) postAlert(data ObstacleDistance) {
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	for _, outputChan := range odp.alertChannels {
		outputChan <- data
	}
}

func (odp *ObstacleDetectorProcess) closeChannels() {
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	for _, outputChan := range odp.alertChannels {
		close(outputChan)
	}
}

func (odp *ObstacleDetectorProcess) Subscribe(name string, channel chan<- ObstacleDistance) {
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	odp.alertChannels[name] = channel
}

func ObstacleChannelListener(channel chan ObstacleDistance, far, medium, close func()) {
	for value := range channel {
		if value == Far {
			far()
		} else if value == Medium {
			medium()
		} else {
			close()
		}
	}
}
