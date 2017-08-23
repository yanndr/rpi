package config

import (
	"encoding/json"
	"fmt"
	"os"
)

//BotConfiguration Represent a configuration of the Bot
type BotConfiguration struct {
	MotorController  dualMotorController
	UltraSoundSensor ultraSoundSensor
}

//DualMotorController Configuration of the dualmotor controller
type dualMotorController struct {
	Motor1, Motor2  motorController
	LeftCorrection  float64
	RightCorrection float64
	MaxSpeed        float64
	MinSpeed        float64
}

//MotorController Configuration of a motor
type motorController struct {
	Pin1, Pin2 uint8
}

type ultraSoundSensor struct {
	Echo, Trigger uint8
}

//LoadConfig Retrun the configuration of the Bot
func LoadConfig() (BotConfiguration, error) {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := BotConfiguration{}

	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
		return configuration, err
	}

	return configuration, nil
}
