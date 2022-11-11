package cards

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
)

func (s *service) InsertCards(ctx context.Context, deckID int32, cardIDs []int32) error {
	err := s.store.ExecTx(ctx, func(q maindb.Querier) (bool, error) {
		for _, c := range cardIDs {
			err := q.CopyCard(ctx, maindb.CopyCardParams{
				ID:     c,
				DeckID: deckID,
			})
			if err != nil {
				return false, err
			}
		}

		return true, nil
	})

	if err != nil {
		return errs.E(err, errs.Database, errs.Code("insert_cards_failed"))
	}

	return nil
}
