package text

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	. "github.com/ahmetalpbalkan/go-linq"
)

type Mood int

const (
	Neutral Mood = iota
	Happy
	Angry
	Annoyed
	Scarred
)

type Situation int

const (
	Any Situation = iota
)

type SituationText struct {
	text string
	mood Mood
}

type TextGenerator interface {
	Text(m Mood, s Situation) string
}

type MemoryText struct {
	source map[Situation][]SituationText
}

func NewMemoryText() *MemoryText {

	text := SituationText{text: "I like big but and I cannot lie", mood: Neutral}
	text2 := SituationText{text: "Vive les slips", mood: Neutral}
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
		return t.(SituationText).mood == m
	}).Select(func(t interface{}) interface{} {
		return t.(SituationText).text
	}).ToSlice(&result)

	l := len(result)
	if l == 0 {
		return "I have nothing to say."
	}
	rand.Seed(time.Now().UTC().UnixNano())
	return result[rand.Intn(len(result))]

}
