package auth

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
)

func (s *service) User(ctx context.Context, userID int32) (*maindb.User, error) {
	user, err := s.store.GetUser(ctx, userID)
	if err != nil {
		return nil, errs.E(err, errs.Database, errs.Code("get_user_failed"))
	}

	return user, nil
}
