package cards

import (
	"github.com/4kord/english-flashcards/pkg/services/cards"
	"go.uber.org/zap"
)

type Controller struct {
	CardService cards.Service
	Log         *zap.Logger
}
