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
			continue
		}
	}
}

func (b *Bot) decayEmotions() {

}
