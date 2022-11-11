package auth

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
	"github.com/google/uuid"
)

type RefreshParams struct {
	RefreshToken string
	UserAgent    string
}

type RefreshResult struct {
	Session     *maindb.Session
	AccessToken string
}

func (s *service) Refresh(ctx context.Context, arg *RefreshParams) (*RefreshResult, error) {
	result := new(RefreshResult)

	err := s.store.ExecTx(ctx, func(q maindb.Querier) (bool, error) {
		var err error

		session, err := q.GetSessionByToken(ctx, arg.RefreshToken)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return false, errs.E(err, errs.NotExist, errs.Code("session_not_found"))
			}
			return false, errs.E(err, errs.Database, errs.Code("get_session_failed"))
		}

		err = q.DeleteSession(ctx, session.ID)
		if err != nil {
			return false, errs.E(err, errs.Database, errs.Code("delete_session_failed"))
		}

		if session.ExpiresAt.Before(time.Now()) {
			return true, errs.E("Session expired", errs.Unauthenticated, errs.Code("session_expired"))
		}

		user, err := q.GetUser(ctx, session.UserID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return false, errs.E(err, errs.NotExist, errs.Code("user_not_found"))
			}
			return false, errs.E(err, errs.Database, errs.Code("get_user_failed"))
		}

		result.AccessToken, err = s.maker.CreateAccessToken(user.ID, user.Admin, accessTokenExpiration)
		if err != nil {
			return false, errs.E(err, errs.Internal, errs.Code("error_creating_token"))
		}

		refreshToken, err := uuid.NewRandom()
		if err != nil {
			return false, errs.E(err, errs.Internal, errs.Code("generate_uuid_failed"))
		}

		result.Session, err = q.CreateSession(ctx, maindb.CreateSessionParams{
			RefreshToken: refreshToken.String(),
			UserAgent:    arg.UserAgent,
			UserID:       user.ID,
			ExpiresAt:    time.Now().UTC().Add(refreshTokenExpiration),
		})
		if err != nil {
			return false, errs.E(err, errs.Database, errs.Code("create_session_failed"))
		}

		return true, nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
