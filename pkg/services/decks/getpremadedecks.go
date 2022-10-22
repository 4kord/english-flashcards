package decks

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/maindb"
)

func (s *service) GetPremadeDecks(ctx context.Context) ([]*maindb.Deck, error) {
	return s.store.GetPremadeDecks(ctx)
}
