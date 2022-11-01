package cards

import (
	"github.com/4kord/english-flashcards/pkg/services/cards"
	"go.uber.org/zap"
)

type Controller struct {
	CardsService cards.Service
	Log          *zap.Logger
}

func New(cardsService cards.Service, log *zap.Logger) *Controller {
	return &Controller{
		CardsService: cardsService,
		Log:          log,
	}
}
