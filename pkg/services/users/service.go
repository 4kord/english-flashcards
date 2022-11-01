package users

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/maindb"
)

type Service interface {
	GetUsers(ctx context.Context) ([]*maindb.User, error)
	DeleteUser(ctx context.Context, userID int32) error
}

type service struct {
	store maindb.Store
}

func New(store maindb.Store) Service {
	return &service{
		store: store,
	}
}
