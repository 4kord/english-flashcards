package users

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
)

func (s *service) GetUsers(ctx context.Context) ([]*maindb.User, error) {
	users, err := s.store.GetUsers(ctx)
	if err != nil {
		return nil, errs.E(err, errs.Database, errs.Code("get_users_failed"))
	}

	return users, err
}
