package auth

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) RegisterUser(ctx context.Context, email, password string) (*maindb.User, error) {
	crypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errs.E(err, errs.Internal, errs.Code("error_crypt_password"))
	}

	user, err := s.store.CreateUser(ctx, maindb.CreateUserParams{
		Email:    email,
		Password: string(crypt),
		Role:     "user",
	})

	if err, ok := err.(*pq.Error); ok {
		if err.Constraint == "uc_users_email" {
			return nil, errs.E(err, errs.Exist, errs.Code("email_taken"))
		}

		return nil, errs.E(err, errs.Database, errs.Code("operation_create_user_failed"))
	}

	return user, nil
}
