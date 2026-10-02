package bot

const moodCount = 5

// Enum of emotions
type Mood int

const (
	Joy Mood = iota
	Sadness
	Anger
	Fear
	Disgust
)

// BotState containing bot's emotional state
type BotState struct {
	Moods [moodCount]int
}

func (b *Bot) UpdateState(in Mood, value int) {
	if !(in >= 0 && in < moodCount) {
		return
	}

	// 0 <= moodValue <= 100
	if !(value >= 0 && value <= 100) {
		return
	}

	b.BotState.Moods[in] = value
}

type InState struct {
	In    Mood
	Value int
}
