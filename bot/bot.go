package bot

import (
	"2010s/internal/ws"
	"database/sql"
)

type Bot struct {
	hub *ws.Hub
	db  *sql.DB

	botState BotState
	botMood  Mood
}
