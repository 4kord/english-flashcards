package decks

import (
	"context"
	"database/sql"
	"errors"

	"github.com/4kord/english-flashcards/pkg/errs"
)

func (s *service) DeleteDeck(ctx context.Context, deckID int32) error {
	err := s.store.DeleteDeck(ctx, deckID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errs.E(err, errs.NotExist, errs.Code("deck_not_found"))
		}

		return errs.E(err, errs.Database, errs.Code("delete_deck_failed"))
	}

	return err
}
