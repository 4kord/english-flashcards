package cards

import (
	"context"
	"database/sql"
	"errors"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
)

func (s *service) DeleteCard(ctx context.Context, cardID int32) error {
	err := s.store.ExecTx(ctx, func(q maindb.Querier) (bool, error) {
		var err error

		err = q.DeckAmountDownByCard(ctx, cardID)
		if err != nil {
			return false, errs.E(err, errs.Database, errs.Code("deck_amount_down_failed"))
		}

		err = q.DeleteCard(ctx, cardID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return false, errs.E(err, errs.NotExist, errs.Code("card_not_found"))
			}

			return false, errs.E(err, errs.Database, errs.Code("delete_card_failed"))
		}

		return true, nil
	})

	return err
}
