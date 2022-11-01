package models

import "time"

type Deck struct {
	ID        int
	UserID    int
	Name      string
	Amount    int
	IsPremade bool
	CreatedAt time.Time
}
