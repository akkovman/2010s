package bot

import (
	"2010s/internal/ws"
	"database/sql"
)

type Bot struct {
	hub *ws.Hub
	db  *sql.DB

	BotState  BotState
	BotMood   Mood
	SendState chan InState
}

func NewBot(
	h *ws.Hub,
	db *sql.DB,
	initialState BotState) *Bot {
	return &Bot{
		hub: h,
		db:  db,

		BotState:  initialState,
		SendState: make(chan InState),
	}
}
