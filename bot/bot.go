package bot

import (
	"2010s/internal/ws"
	"database/sql"
	"time"
)

type Bot struct {
	hub *ws.Hub
	db  *sql.DB

	BotState BotState
	BotMood  Mood
}

func NewBot(
	h *ws.Hub,
	db *sql.DB,
	initialState BotState) *Bot {
	return &Bot{
		hub: h,
		db:  db,

		BotState: initialState,
	}
}

func (b *Bot) Run() {
	emotionTicker := time.NewTicker(15 * time.Second)
	defer func() {
		emotionTicker.Stop()
	}()

	for {
		select {
		case <-emotionTicker.C:
			b.decayEmotions()
		}
	}
}

/*
Function take current value of mood and, adds the delta,
and ensures the value stays within 0 and 100
*/
func (b *Bot) changeEmotion(mood Mood, delta int) {
	currentValue := b.BotState.moods[mood]
	newValue := max(0, min(100, currentValue+delta))

	b.updateState(mood, newValue)
}

func (b *Bot) decayEmotions() {

}
