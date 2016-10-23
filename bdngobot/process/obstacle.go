package process

import (
	"fmt"
	"time"

	"github.com/yanndr/rpi/bdngobot/situation"
	"github.com/yanndr/rpi/event"
	"github.com/yanndr/rpi/sensor"
)

//ObstacleDetectorProcess process that check the obstacle with an ultrasoundSensor
type ObstacleDetectorProcess struct {
	BaseProcess
	ultrasoundSensor sensor.DistanceSensor
	alertDistance    int
	alertChannels    map[string]chan<- situation.Situation
	warningDistance  int
	lastAlert        situation.Situation
	ticker           *time.Ticker
	alerter          event.Alerter
}

func NewObstacleDetectorProcess(ultrasoundSensor sensor.DistanceSensor, alerter event.Alerter, alertDistance, warningDistance int) *ObstacleDetectorProcess {

	odp := &ObstacleDetectorProcess{
		ultrasoundSensor: ultrasoundSensor,
		alertDistance:    alertDistance,
		warningDistance:  warningDistance,
		lastAlert:        situation.Any,
		alerter:          alerter,
		BaseProcess:      BaseProcess{Channel: make(chan interface{})},
	}

	return odp
}

func (odp *ObstacleDetectorProcess) Start() {

	go ObstacleChannelListener(odp.Channel, func() {}, func() {}, func() {})
	fmt.Println("Obstacle detector process started.")
	odp.ticker = time.NewTicker(time.Second / 4)
	go func() {
		for range odp.ticker.C {

			d, err := odp.ultrasoundSensor.Distance()
			if err != nil {
				fmt.Println(err)
			}

			if float64(odp.alertDistance) > d {
				if odp.lastAlert != situation.ObstacleClose {
					odp.alerter.PostAlert(situation.ObstacleClose)
					odp.lastAlert = situation.ObstacleClose
				}
			} else if float64(odp.warningDistance) > d {
				if odp.lastAlert != situation.ObstacleMedium {
					odp.alerter.PostAlert(situation.ObstacleMedium)
					odp.lastAlert = situation.ObstacleMedium
				}
			} else {
				if odp.lastAlert != situation.ObstacleFar {
					odp.alerter.PostAlert(situation.ObstacleFar)
					odp.lastAlert = situation.ObstacleFar
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
	fmt.Println("Obstacle detector process stopped.")
}

func ObstacleChannelListener(channel chan interface{}, far, medium, close func()) {
	for value := range channel {
		if value == situation.ObstacleFar {
			far()
		} else if value == situation.ObstacleMedium {
			medium()
		} else {
			close()
		}
	}
}
