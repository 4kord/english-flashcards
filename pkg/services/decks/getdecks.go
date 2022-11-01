package decks

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
)

func (s *service) GetDecks(ctx context.Context, userID int32) ([]*maindb.Deck, error) {
	d, err := s.store.GetDecks(ctx, userID)
	if err != nil {
		return nil, errs.E(err, errs.Database, errs.Code("get_decks_failed"))
	}

	return d, err
}
