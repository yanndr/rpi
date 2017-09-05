/*
export GOOS=linux
export GOARCH=arm
*/
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/yanndr/rpi/controller/led"
	"github.com/yanndr/rpi/pwm"

	"github.com/kidoman/embd"
	"github.com/kidoman/embd/controller/pca9685"
	"github.com/yanndr/rpi/bdngobot/config"
	"github.com/yanndr/rpi/bdngobot/process"
	"github.com/yanndr/rpi/bdngobot/process/decision"
	"github.com/yanndr/rpi/bdngobot/process/mouvement"
	"github.com/yanndr/rpi/bdngobot/process/speech"
	"github.com/yanndr/rpi/bdngobot/text"
	"github.com/yanndr/rpi/controller"
	"github.com/yanndr/rpi/event"
	"github.com/yanndr/rpi/gpio"
	"github.com/yanndr/rpi/media"
	"github.com/yanndr/rpi/sensor"
	"github.com/yanndr/rpi/tts"
)

func main() {

	var motorsController controller.MotorsController
	var ultrasoundSensor sensor.DistanceSensor

	processes := make(map[string]process.Process)

	bdnConfig, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(bdnConfig)

	// Open and map memory to access gpio, check for errors
	err = gpio.Open()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer gpio.Close()

	pca := pca9685.New(embd.NewI2CBus(0x1), 0x40)
	pca.Freq = 60

	defer pca.Close()

	pca9685 := pwm.NewPca9685(pca)

	motorsController, err = controller.NewDualMotorController(
		controller.NewDRV833MotorController("left motor", bdnConfig.MotorController.Motor1.Pin1, bdnConfig.MotorController.Motor1.Pin2, pca9685),
		controller.NewDRV833MotorController("right motor", bdnConfig.MotorController.Motor2.Pin1, bdnConfig.MotorController.Motor2.Pin2, pca9685),
		bdnConfig.MotorController.LeftCorrection, bdnConfig.MotorController.RightCorrection, bdnConfig.MotorController.MaxSpeed, bdnConfig.MotorController.MinSpeed)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ultrasoundSensor = sensor.NewHCSRO4Sensor(bdnConfig.UltraSoundSensor.Trigger, bdnConfig.UltraSoundSensor.Echo)

	ed := event.NewEventDispatcher()

	processes["mouvment"] = mouvement.NewMouvementProcess(motorsController)
	processes["speech"] = speech.NewSpeechProcess(&tts.Festival{}, text.NewMemoryText("text.json"))
	processes["player"] = process.NewPlayerProcess(&media.OmxPlayer{})
	processes["obstacle"] = process.NewObstacleDetectorProcess(ultrasoundSensor, ed, 30.0, 60.0)

	for name, p := range processes {
		fmt.Println("start process ", name)
		ed.Subscribe(name, p.Chan())
		p.Start()
	}
	processes["decision"] = decision.NewDecisionProcess(ed)
	ed.Subscribe("decision", processes["decision"].Chan())
	processes["decision"].Start()

	lc := led.NewLedController(pca9685, 4, 5, 6, 7, 8, 9, 10, 11)
	lc.BlinkAll(time.Second/8, 0, 0.8, 0.1)

	fmt.Println("Q to kill Bdnbot")
	var response int
	for response != 'q' {
		fmt.Scanf("%c", &response)

		switch response {
		default:
			fmt.Println("Pouet")
		case 'q':
			fmt.Println("Killing BdnBot")
		}
	}

	for _, p := range processes {
		p.Stop()
	}
	// lc.SetAllOff()
}
