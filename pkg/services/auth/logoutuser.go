package auth

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/errs"
)

func (s *service) LogoutUser(ctx context.Context, session string) error {
	err := s.store.DeleteSessionByToken(ctx, session)
	if err != nil {
		return errs.E(err, errs.Database, errs.Code("logout_user_failed"))
	}

	return nil
}
