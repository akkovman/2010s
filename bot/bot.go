package bot

import (
	"2010s/internal/ws"
	"database/sql"
	"time"
)

type Bot struct {
	hub *ws.Hub
	db  *sql.DB

	BotState  BotState
	BotMood   Mood
	SendEvent chan BotEvent
}

func NewBot(
	h *ws.Hub,
	db *sql.DB,
	initialState BotState) *Bot {
	return &Bot{
		hub: h,
		db:  db,

		BotState:  initialState,
		SendEvent: make(chan BotEvent),
	}
}

func (b *Bot) Run() {
	emotionTicker := time.NewTicker(15 * time.Second)
	defer func() {
		emotionTicker.Stop()
	}()

	for {
		select {
		case event := <-b.SendEvent:
			b.handleEvent(event)
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
	// Every 15 seconds
	b.changeEmotion(Joy, -5)
	b.changeEmotion(Sadness, 2)
	b.changeEmotion(Anger, -10)
	b.changeEmotion(Fear, -3)
	b.changeEmotion(Disgust, -1)
}

func (b *Bot) handleEvent(event BotEvent) {

}
