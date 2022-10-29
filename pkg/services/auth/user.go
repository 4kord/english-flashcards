package auth

import (
	"context"
	"database/sql"
	"errors"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
)

func (s *service) User(ctx context.Context, session string) (*maindb.User, *maindb.Session, error) {
	returnedSession, err := s.store.GetSessionBySession(ctx, session)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, errs.E(err, errs.NotExist, errs.Code("session_not_found"))
		}

		return nil, nil, errs.E(err, errs.Database, "get_session_failed")
	}

	returnedUser, err := s.store.GetUser(ctx, returnedSession.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, errs.E(err, errs.NotExist, errs.Code("user_not_found"))
		}

		return nil, nil, errs.E(err, errs.Database, "get_user_failed")
	}

	return returnedUser, returnedSession, nil
}
