package auth

import (
	"context"
	"database/sql"
	"errors"
	"math/rand"
	"time"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) LoginUser(ctx context.Context, email, password string) (*maindb.User, *maindb.Session, error) {
	var User *maindb.User

	var Session *maindb.Session

	err := s.store.ExecTx(ctx, func(q *maindb.Queries) error {
		user, err := q.GetUserByEmail(ctx, email)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return errs.E(err, errs.NotExist, errs.Code("email_not_found"))
			}

			return errs.E(err, errs.Database, errs.Code("operation_get_user_failed"))
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
				return errs.E(err, errs.Unauthenticated, errs.Code("incorrect_password"))
			}

			return errs.E(err, errs.Database, errs.Code("compare_hash_and_password_failed"))
		}

		amount, err := s.store.CountSessions(ctx, user.ID)
		if err != nil {
			return errs.E(err, errs.Database, errs.Code("count_sessions_failed"))
		}

		if amount >= 5 {
			err := s.store.DeleteOldestSession(ctx)
			if err != nil {
				return errs.E(err, errs.Database, errs.Code("delete_oldest_session_failed"))
			}
		}

		rand.Seed(time.Now().UnixNano())

		var symbols = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
		b := make([]byte, 20)
		for i := range b {
			b[i] = symbols[rand.Intn(len(symbols))]
		}

		generatedSession := string(b)

		expiresAt := time.Now().UTC().Add(30 * 24 * time.Hour)

		session, err := q.CreateSession(ctx, maindb.CreateSessionParams{
			Session:   generatedSession,
			UserID:    user.ID,
			Ip:        "0.0.0.0",
			ExpiresAt: expiresAt,
		})

		if err != nil {
			errs.E(err, errs.Database, "create_session_failed")
		}

		User = user
		Session = session

		return nil
	})

	return User, Session, err
}
