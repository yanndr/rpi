/*
export GOOS=linux
export GOARCH=arm
*/
package main

import (
	"fmt"
	"os"

	"github.com/yanndr/rpi/bdngobot/config"
	"github.com/yanndr/rpi/bdngobot/process"
	"github.com/yanndr/rpi/controller"
	"github.com/yanndr/rpi/gpio"
	"github.com/yanndr/rpi/media"
	"github.com/yanndr/rpi/sensor"
	"github.com/yanndr/rpi/tts"
)

func main() {

	var motorsController controller.MotorsController
	var ultrasoundSensor sensor.DistanceSensor
	var obstacleDetector *process.ObstacleDetectorProcess
	var speechProcess *process.SpeechProcess
	var mouvementProcess *process.MouvementProcess
	var playerProcess *process.PlayerProcess
	//var err error

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

	motorsController, err = controller.NewDualMotorController(
		controller.NewL298NMotorController("left motor", bdnConfig.MotorController.Motor1.Pin1, bdnConfig.MotorController.Motor1.Pin2, bdnConfig.MotorController.Motor1.SpeedPin),
		controller.NewL298NMotorController("right motor", bdnConfig.MotorController.Motor2.Pin1, bdnConfig.MotorController.Motor2.Pin2, bdnConfig.MotorController.Motor2.SpeedPin),
		bdnConfig.MotorController.LeftCorrection, bdnConfig.MotorController.RightCorrection, bdnConfig.MotorController.MaxSpeed, bdnConfig.MotorController.MinSpeed)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ultrasoundSensor = sensor.NewHCSRO4Sensor(bdnConfig.UltraSoundSensor.Trigger, bdnConfig.UltraSoundSensor.Echo)
	mouvementProcess = process.NewMouvementProcess(motorsController)
	speechProcess = process.NewSpeechProcess(&tts.Festival{})
	playerProcess = process.NewPlayerProcess(&media.OmxPlayer{})

	obstacleDetector = process.NewObstacleDetectorProcess(ultrasoundSensor, 30.0, 60.0)

	obstacleDetector.Subscribe("mouvement", mouvementProcess.DistanceAlertChannel)
	obstacleDetector.Subscribe("speak", speechProcess.DistanceAlertChannel)
	obstacleDetector.Subscribe("sing", playerProcess.DistanceAlertChannel)

	obstacleDetector.Start()
	mouvementProcess.Start()
	speechProcess.Start()
	playerProcess.Start()

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

	speechProcess.Stop()
	obstacleDetector.Stop()
	motorsController.Stop()
	playerProcess.Stop()

}
