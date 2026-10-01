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
	moods [moodCount]int
}
