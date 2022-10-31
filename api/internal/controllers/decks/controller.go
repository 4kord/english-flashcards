package decks

import (
	"github.com/4kord/english-flashcards/pkg/services/decks"
	"go.uber.org/zap"
)

type Controller struct {
	DecksService decks.Service
	Log          *zap.Logger
}

func New(decksServices decks.Service, log *zap.Logger) *Controller {
	return &Controller{
		DecksService: decksServices,
		Log:          log,
	}
}
