package auth

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/maindb"
)

type Service interface {
	RegisterUser(ctx context.Context, email, password string) error
	LoginUser(ctx context.Context, email, password string) (*maindb.User, *maindb.Session, error)
	LogoutUser(ctx context.Context, session string) error
	User(ctx context.Context, session string) (*maindb.User, *maindb.Session, error)
}

type service struct {
	store *maindb.Store
}

func New(store *maindb.Store) Service {
	return &service{
		store: store,
	}
}
