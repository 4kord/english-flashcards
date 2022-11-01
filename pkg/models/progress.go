package models

import "time"

type Progress struct {
	ID           int
	CardID       int
	UserID       int
	Learned      bool
	Selected     bool
	Correct      int
	Incorrect    int
	LastViewedAt time.Time
}
