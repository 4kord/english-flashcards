package services

import (
	"github.com/4kord/english-flashcards/pkg/cld"
	"github.com/4kord/english-flashcards/pkg/maindb"
	"github.com/4kord/english-flashcards/pkg/maker"
	"github.com/4kord/english-flashcards/pkg/services/auth"
	"github.com/4kord/english-flashcards/pkg/services/cards"
	"github.com/4kord/english-flashcards/pkg/services/decks"
	"github.com/4kord/english-flashcards/pkg/services/google"
	"github.com/4kord/english-flashcards/pkg/services/users"
)

type Config struct {
	Store *maindb.Store
	Cld   cld.Cld
	Maker maker.Maker
}

type Services struct {
	Auth   auth.Service
	Users  users.Service
	Decks  decks.Service
	Cards  cards.Service
	Google google.Service
}

func New(cfg *Config) *Services {
	return &Services{
		Auth:   auth.New(cfg.Store, cfg.Maker),
		Users:  users.New(cfg.Store),
		Decks:  decks.New(cfg.Store),
		Cards:  cards.New(cfg.Store, cfg.Cld),
		Google: google.New(),
	}
}
