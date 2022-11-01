package models

import "time"

type Session struct {
	ID        int
	Session   string
	UserID    int
	IP        string
	ExpiresAt time.Time
	CreatedAt time.Time
}
