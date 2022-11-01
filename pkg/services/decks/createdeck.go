package decks

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
)

type CreateDeckParams struct {
	UserID    int32
	Name      string
	IsPremade bool
}

func (s *service) CreateDeck(ctx context.Context, arg *CreateDeckParams) (*maindb.Deck, error) {
	d, err := s.store.CreateDeck(ctx, maindb.CreateDeckParams{
		UserID:    arg.UserID,
		Name:      arg.Name,
		IsPremade: arg.IsPremade,
	})
	if err != nil {
		return nil, errs.E(err, errs.Database, errs.Code("create_deck_failed"))
	}

	return d, nil
}
