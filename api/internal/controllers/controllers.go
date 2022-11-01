package controllers

import (
	"github.com/4kord/english-flashcards/api/internal/controllers/auth"
	"github.com/4kord/english-flashcards/api/internal/controllers/cards"
	"github.com/4kord/english-flashcards/api/internal/controllers/decks"
	"github.com/4kord/english-flashcards/api/internal/controllers/google"
	"github.com/4kord/english-flashcards/api/internal/controllers/users"
	"github.com/4kord/english-flashcards/pkg/services"
	"go.uber.org/zap"
)

type Config struct {
	Services *services.Services
	Logger   *zap.Logger
}

type Controllers struct {
	Auth   *auth.Controller
	Users  *users.Controller
	Cards  *cards.Controller
	Decks  *decks.Controller
	Google *google.Controller
}

func New(cfg *Config) *Controllers {
	return &Controllers{
		Auth:   auth.New(cfg.Services.Auth, cfg.Logger),
		Users:  users.New(cfg.Services.Users, cfg.Logger),
		Cards:  cards.New(cfg.Services.Cards, cfg.Logger),
		Decks:  decks.New(cfg.Services.Decks, cfg.Logger),
		Google: google.New(cfg.Services.Google, cfg.Logger),
	}
}
