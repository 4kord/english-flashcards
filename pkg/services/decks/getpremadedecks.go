package decks

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
)

func (s *service) GetPremadeDecks(ctx context.Context) ([]*maindb.Deck, error) {
	d, err := s.store.GetPremadeDecks(ctx)
	if err != nil {
		return nil, errs.E(err, errs.Database, "get_premade_decks_failed")
	}

	return d, err
}
