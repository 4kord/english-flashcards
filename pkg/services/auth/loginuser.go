package auth

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
	"github.com/4kord/english-flashcards/pkg/services/auth/dto"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) LoginUser(ctx context.Context, email, password string) (*dto.LoginUserResult, error) {
	var result = new(dto.LoginUserResult)

	err := s.store.ExecTx(ctx, func(q *maindb.Queries) error {
		var err error

		result.User, err = q.GetUserByEmail(ctx, email)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return errs.E(err, errs.NotExist, errs.Code("email_not_found"))
			}

			return errs.E(err, errs.Database, errs.Code("operation_get_user_failed"))
		}

		err = bcrypt.CompareHashAndPassword([]byte(result.User.Password), []byte(password))
		if err != nil {
			if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
				return errs.E(err, errs.Unauthenticated, errs.Code("incorrect_password"))
			}

			return errs.E(err, errs.Database, errs.Code("compare_hash_and_password_failed"))
		}

		result.AccessToken, err = s.maker.CreateAccessToken(result.User.ID, result.User.Admin, time.Minute*5)
		if err != nil {
			return errs.E(err, errs.Internal, errs.Code("error_creating_token"))
		}

		result.RefreshToken, err = s.maker.CreateRefreshToken(result.User.ID, time.Hour*24*30)
		if err != nil {
			return errs.E(err, errs.Internal, errs.Code("error_creating_token"))
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, err
}
