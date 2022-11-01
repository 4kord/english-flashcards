package models

import "time"

type Session struct {
	ID        int32
	Session   string
	UserID    int32
	Ip        string
	ExpiresAt time.Time
	CreatedAt time.Time
}
