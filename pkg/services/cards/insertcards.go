package cards

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
	"github.com/4kord/english-flashcards/pkg/services/cards/dto"
)

func (s *service) InsertCards(ctx context.Context, arg *dto.InsertCardsParams) error {
	err := s.store.ExecTx(ctx, func(q *maindb.Queries) error {
		for _, c := range arg.CardIDs {
			err := q.CopyCard(ctx, maindb.CopyCardParams{
				ID:     c,
				DeckID: arg.DeckID,
			})
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return errs.E(err, errs.Database, errs.Code("insert_cards_failed"))
	}

	return nil
}
