package decks

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
)

func (s *service) CreateDeck(ctx context.Context, deck *maindb.Deck) (*maindb.Deck, error) {
	d, err := s.store.CreateDeck(ctx, maindb.CreateDeckParams{
		UserID:    deck.UserID,
		Name:      deck.Name,
		IsPremade: deck.IsPremade,
	})
	if err != nil {
		return nil, errs.E(err, errs.Database, errs.Code("create_deck_failed"))
	}

	return d, nil
}
