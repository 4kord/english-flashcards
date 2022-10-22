package decks

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/maindb"
)

func (s *service) EditDeck(ctx context.Context, deck maindb.Deck) (*maindb.Deck, error) {
	return s.store.EditDeck(ctx, maindb.EditDeckParams{
		ID:   deck.ID,
		Name: deck.Name,
	})
}
