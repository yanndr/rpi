package text

import (
	"encoding/json"
	"fmt"
	"github.com/yanndr/rpi/bdngobot/mood"
	"github.com/yanndr/rpi/bdngobot/situation"
	"math/rand"
	"os"
	"time"
)

type TextGenerator interface {
	Text(m mood.Mood, s situation.Situation) string
}

type MemoryText struct {
	source map[situation.Situation]map[mood.Mood][]string
}

func NewMemoryText(textFile string) *MemoryText {

	source := make(map[situation.Situation]map[mood.Mood][]string)

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

func (t *MemoryText) Text(m mood.Mood, s situation.Situation) string {

	moods, ok := t.source[s]
	if !ok {
		return "I have nothing to say."
	}

	sentences := moods[m]

	l := len(sentences)
	if l == 0 {
		return "I have nothing to say."
	}
	rand.Seed(time.Now().UTC().UnixNano())
	return sentences[rand.Intn(len(sentences))]

}
