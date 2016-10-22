package text

import (
	"encoding/json"
	"fmt"
	"math/rand"
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
	Any Situation = "Any"
)

type SituationText struct {
	Text string
	Mood Mood
}

type TextGenerator interface {
	Text(m Mood, s Situation) string
}

type MemoryText struct {
	source map[Situation][]SituationText
}

func NewMemoryText() *MemoryText {

	text := SituationText{Text: "I like big but and I cannot lie", Mood: Neutral}
	text2 := SituationText{Text: "Vive les slips", Mood: Neutral}
	source := make(map[Situation][]SituationText)

	source[Any] = []SituationText{text, text2}

	b, err := json.Marshal(source)
	if err == nil {
		s := string(b[:])
		fmt.Println(s)
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
