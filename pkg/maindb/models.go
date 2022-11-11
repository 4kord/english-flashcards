// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package maindb

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

type Deck struct {
	ID        int32
	UserID    int32
	Name      string
	Amount    int32
	IsPremade bool
	CreatedAt time.Time
}

type Progress struct {
	ID           int32
	CardID       int32
	UserID       int32
	Learned      bool
	Selected     bool
	Correct      int32
	Incorrect    int32
	LastViewedAt time.Time
}

type Session struct {
	ID           int32
	RefreshToken string
	UserAgent    string
	UserID       int32
	ExpiresAt    time.Time
	CreatedAt    time.Time
}

type User struct {
	ID        int32
	Email     string
	Password  string
	Admin     bool
	CreatedAt time.Time
}
