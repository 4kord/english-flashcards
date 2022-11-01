package cards

import (
	"context"
	"database/sql"
	"errors"

	"github.com/4kord/english-flashcards/pkg/errs"
)

func (s *service) DeleteCard(ctx context.Context, cardID int32) error {
	err := s.store.DeleteCard(ctx, cardID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errs.E(err, errs.NotExist, "card_not_found")
		}

		return errs.E(err, errs.Database, "delete_card_failed")
	}

	return nil
}
