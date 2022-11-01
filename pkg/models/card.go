package models

import (
	"database/sql"
	"time"
)

type Card struct {
	ID            int32
	DeckID        int32
	English       string
	Russian       string
	Association   sql.NullString
	Example       sql.NullString
	Transcription sql.NullString
	Image         sql.NullString
	ImageUrl      sql.NullString
	Audio         sql.NullString
	AudioUrl      sql.NullString
	CreatedAt     time.Time
}
