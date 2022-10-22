package services

import (
	"github.com/4kord/english-flashcards/pkg/cld"
	"github.com/4kord/english-flashcards/pkg/maindb"
	"github.com/4kord/english-flashcards/pkg/services/auth"
	"github.com/4kord/english-flashcards/pkg/services/cards"
	"github.com/4kord/english-flashcards/pkg/services/decks"
)

type Config struct {
	Store *maindb.Store
	Cld   *cld.Cld
}

type Services struct {
	Auth  auth.Service
	Decks decks.Service
	Cards cards.Service
}

func New(cfg *Config) *Services {
	return &Services{
		Auth:  auth.New(cfg.Store),
		Decks: decks.New(cfg.Store),
		Cards: cards.New(cfg.Store, cfg.Cld),
	}
}
