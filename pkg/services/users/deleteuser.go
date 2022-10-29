package users

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/errs"
)

func (s *service) DeleteUser(ctx context.Context, userID int32) error {
	err := s.store.DeleteUser(ctx, userID)
	if err != nil {
		return errs.E(err, errs.Database, errs.Code("delete_user_failed"))
	}

	return nil
}
