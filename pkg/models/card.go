package models

import (
	"database/sql"
	"time"
)

type Card struct {
	ID            int
	DeckID        int
	English       string
	Russian       string
	Association   sql.NullString
	Example       sql.NullString
	Transcription sql.NullString
	Image         sql.NullString
	ImageURL      sql.NullString
	Audio         sql.NullString
	AudioURL      sql.NullString
	CreatedAt     time.Time
}
