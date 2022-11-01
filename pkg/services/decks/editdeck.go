package decks

import (
	"context"
	"database/sql"
	"errors"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
)

type EditDeckParams struct {
	DeckID  int32
	NewName string
}

func (s *service) EditDeck(ctx context.Context, arg *EditDeckParams) (*maindb.Deck, error) {
	d, err := s.store.EditDeck(ctx, maindb.EditDeckParams{
		ID:   arg.DeckID,
		Name: arg.NewName,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.E(err, errs.NotExist, errs.Code("deck_not_found"))
		}

		return nil, errs.E(err, errs.Database, errs.Code("edit_deck_failed"))
	}

	return d, err
}
