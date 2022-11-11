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

type RefreshResult struct {
	User        *maindb.User
	Session     *maindb.Session
	AccessToken string
}

func (s *service) Refresh(ctx context.Context, refreshToken string) (*RefreshResult, error) {
	result := new(RefreshResult)

	err := s.store.ExecTx(ctx, func(q maindb.Querier) error {
		var err error

		session, err := q.GetSessionByToken(ctx, refreshToken)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return errs.E(err, errs.NotExist, errs.Code("session_not_found"))
			}
			return errs.E(err, errs.Database, errs.Code("get_session_failed"))
		}

		err = q.DeleteSession(ctx, session.ID)
		if err != nil {
			return errs.E(err, errs.Database, errs.Code("delete_session_failed"))
		}

		result.User, err = q.GetUser(ctx, session.UserID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return errs.E(err, errs.NotExist, errs.Code("user_not_found"))
			}
			return errs.E(err, errs.Database, errs.Code("get_user_failed"))
		}

		result.AccessToken, err = s.maker.CreateAccessToken(result.User.ID, result.User.Admin, accessTokenExpiration)
		if err != nil {
			return errs.E(err, errs.Internal, errs.Code("error_creating_token"))
		}

		refreshToken, err := uuid.NewRandom()
		if err != nil {
			return errs.E(err, errs.Internal, errs.Code("generate_uuid_failed"))
		}

		result.Session, err = q.CreateSession(ctx, maindb.CreateSessionParams{
			RefreshToken: refreshToken.String(),
			UserAgent:    "",
			ClientIp:     "",
			UserID:       result.User.ID,
			ExpiresAt:    time.Now().UTC().Add(refreshTokenExpiration),
		})
		if err != nil {
			return errs.E(err, errs.Database, errs.Code("create_session_failed"))
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
