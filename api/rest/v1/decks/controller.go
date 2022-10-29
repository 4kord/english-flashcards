package decks

import (
	"github.com/4kord/english-flashcards/pkg/services/decks"
	"go.uber.org/zap"
)

type Controller struct {
	DeckService decks.Service
	Log         *zap.Logger
}
