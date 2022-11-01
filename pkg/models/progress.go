package models

import "time"

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
