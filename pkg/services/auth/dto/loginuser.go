package dto

import "github.com/4kord/english-flashcards/pkg/maindb"

type LoginUserResult struct {
	User         *maindb.User
	AccessToken  string
	RefreshToken string
}
