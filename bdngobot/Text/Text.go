package text

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	. "github.com/ahmetalpbalkan/go-linq"
)

type Mood string

const (
	Neutral Mood = "Neutral"
	Happy   Mood = "Happy"
	Angry   Mood = "Angry"
	Annoyed Mood = "Annoyed"
	Scarred Mood = "Scarred"
)

type Situation string

const (
	Any            Situation = "Any"
	ObstacleClose  Situation = "ObstacleClose"
	ObstacleMedium Situation = "ObstacleMedium"
)

type SituationText struct {
	Text string `json:"text"`
	Mood Mood   `json:"mood"`
}

type TextGenerator interface {
	Text(m Mood, s Situation) string
}

type MemoryText struct {
	source map[Situation][]SituationText
}

func NewMemoryText(textFile string) *MemoryText {

	source := make(map[Situation][]SituationText)

	file, _ := os.Open(textFile)
	decoder := json.NewDecoder(file)

	err := decoder.Decode(&source)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}

	return &MemoryText{
		source: source,
	}
}

func (t *MemoryText) Text(m Mood, s Situation) string {

	ts, ok := t.source[s]
	if !ok {
		return "I have nothing to say."
	}

	result := []string{}
	From(ts).Where(func(t interface{}) bool {
		return t.(SituationText).Mood == m
	}).Select(func(t interface{}) interface{} {
		return t.(SituationText).Text
	}).ToSlice(&result)

	l := len(result)
	if l == 0 {
		return "I have nothing to say."
	}
	rand.Seed(time.Now().UTC().UnixNano())
	return result[rand.Intn(len(result))]

}
