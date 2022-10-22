package decks

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/maindb"
)

func (s *service) CreateDeck(ctx context.Context, deck maindb.Deck) (*maindb.Deck, error) {
	return s.store.CreateDeck(ctx, maindb.CreateDeckParams{
		UserID:    deck.UserID,
		Name:      deck.Name,
		IsPremade: deck.IsPremade,
	})
}
