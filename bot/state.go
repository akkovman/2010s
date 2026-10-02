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

	maxVolume := -1
	dominantMood := Joy

	for moodIndex, volume := range b.BotState.Moods {
		if volume > maxVolume {
			maxVolume = volume
			dominantMood = Mood(moodIndex)
		}
	}

	b.BotMood = dominantMood
}

type InState struct {
	In    Mood
	Value int
}
