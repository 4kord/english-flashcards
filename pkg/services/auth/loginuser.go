package auth

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginUserParams struct {
	Email     string
	Password  string
	UserAgent string
}

type LoginUserResult struct {
	User        *maindb.User
	Session     *maindb.Session
	AccessToken string
}

func (s *service) LoginUser(ctx context.Context, arg *LoginUserParams) (*LoginUserResult, error) {
	var result = new(LoginUserResult)

	err := s.store.ExecTx(ctx, func(q maindb.Querier) (bool, error) {
		var err error

		result.User, err = q.GetUserByEmail(ctx, arg.Email)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return false, errs.E(err, errs.NotExist, errs.Code("email_not_found"))
			}

			return false, errs.E(err, errs.Database, errs.Code("operation_get_user_failed"))
		}

		err = bcrypt.CompareHashAndPassword([]byte(result.User.Password), []byte(arg.Password))
		if err != nil {
			if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
				return false, errs.E(err, errs.Unauthenticated, errs.Code("incorrect_password"))
			}

			return false, errs.E(err, errs.Database, errs.Code("compare_hash_and_password_failed"))
		}

		result.AccessToken, err = s.maker.CreateAccessToken(result.User.ID, result.User.Admin, time.Minute*5)
		if err != nil {
			return false, errs.E(err, errs.Internal, errs.Code("error_creating_token"))
		}

		refreshToken, err := uuid.NewRandom()
		if err != nil {
			return false, errs.E(err, errs.Internal, errs.Code("generate_uuid_failed"))
		}

		amount, err := q.CountSessions(ctx, result.User.ID)
		if err != nil {
			return false, errs.E(err, errs.Database, "count_esssions_failed")
		}

		if amount > 4 {
			err = q.DeleteOldestSession(ctx)
			if err != nil {
				return false, errs.E(err, errs.Database, "delete_oldest_session_failed")
			}
		}

		result.Session, err = q.CreateSession(ctx, maindb.CreateSessionParams{
			RefreshToken: refreshToken.String(),
			UserAgent:    arg.UserAgent,
			UserID:       result.User.ID,
			ExpiresAt:    time.Now().UTC().Add(refreshTokenExpiration),
		})
		if err != nil {
			return false, errs.E(err, errs.Database, "create_session_failed")
		}

		return true, nil
	})

	if err != nil {
		return nil, err
	}

	return result, err
}
