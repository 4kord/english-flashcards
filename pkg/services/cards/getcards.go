package cards

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
)

func (s *service) GetCards(ctx context.Context, deckID int32) ([]*maindb.Card, error) {
	card, err := s.store.GetCards(ctx, deckID)
	if err != nil {
		return nil, errs.E(err, errs.Database, errs.Code("get_cards_failed"))
	}

	return card, nil
}
