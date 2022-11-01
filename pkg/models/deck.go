package models

import "time"

type Deck struct {
	ID        int32
	UserID    int32
	Name      string
	Amount    int32
	IsPremade bool
	CreatedAt time.Time
}
