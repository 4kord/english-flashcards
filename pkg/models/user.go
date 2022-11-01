package models

import "time"

type User struct {
	ID        int32
	Email     string
	Password  string
	Admin     bool
	CreatedAt time.Time
}
