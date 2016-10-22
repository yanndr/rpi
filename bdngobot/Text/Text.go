package Text

type Mood int

const (
	Neutral Mood = iota
	Angry
	Annoyed
	Scarred
)

type Situation int

const (
	Any Situation = iota
)

type TextGenerator interface {
	Text(m Mood, s Situation) string
}

type MemoryText struct {
	source map[Situation][]string
}

func NewMemoryText() *MemoryText {
	return &MemoryText{
		source: make(map[Situation][]string),
	}
}

func (t *MemoryText) Text(m Mood, s Situation) string {

}
