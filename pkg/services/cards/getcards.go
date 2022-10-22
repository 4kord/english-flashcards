package cards

import (
	"context"
	"database/sql"
	"errors"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
)

func (s *service) GetCards(ctx context.Context, deckId int32) ([]*maindb.Card, error) {
	card, err := s.store.GetCards(ctx, deckId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.E(err, errs.NotExist, "card_not_found")
		}
		return nil, errs.E(err, errs.Database, "get_cards_failed")
	}

	return card, nil
}
