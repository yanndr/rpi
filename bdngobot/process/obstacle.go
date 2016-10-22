package process

import (
	"fmt"
	"time"

	"github.com/yanndr/rpi/event"
	"github.com/yanndr/rpi/sensor"
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
	baseProcess
	ultrasoundSensor sensor.DistanceSensor
	alertDistance    ObstacleDistance
	alertChannels    map[string]chan<- ObstacleDistance
	warningDistance  ObstacleDistance
	lastAlert        ObstacleDistance
	ticker           *time.Ticker
	alerter          event.Alerter
}

func NewObstacleDetectorProcess(ultrasoundSensor sensor.DistanceSensor, alerter event.Alerter, alertDistance, warningDistance ObstacleDistance) *ObstacleDetectorProcess {

	odp := &ObstacleDetectorProcess{
		ultrasoundSensor: ultrasoundSensor,
		alertDistance:    alertDistance,
		warningDistance:  warningDistance,
		lastAlert:        None,
		alerter:          alerter,
		baseProcess:      baseProcess{channel: make(chan interface{})},
	}

	return odp
}

func (odp *ObstacleDetectorProcess) Start() {

	go ObstacleChannelListener(odp.channel, func() {}, func() {}, func() {})
	fmt.Println("Obstacle detector process started.")
	odp.ticker = time.NewTicker(time.Second / 4)
	go func() {
		for range odp.ticker.C {

			d, err := odp.ultrasoundSensor.Distance()
			if err != nil {
				fmt.Println(err)
			}

			if float64(odp.alertDistance) > d {
				if odp.lastAlert != Close {
					odp.alerter.PostAlert(Close)
					odp.lastAlert = Close
				}
			} else if float64(odp.warningDistance) > d {
				if odp.lastAlert != Medium {
					odp.alerter.PostAlert(Medium)
					odp.lastAlert = Medium
				}
			} else {
				if odp.lastAlert != Far {
					odp.alerter.PostAlert(Far)
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
	fmt.Println("Obstacle detector process stopped.")
}

func ObstacleChannelListener(channel chan interface{}, far, medium, close func()) {
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
