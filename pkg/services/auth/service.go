package auth

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/maindb"
	"github.com/4kord/english-flashcards/pkg/maker"
	"github.com/4kord/english-flashcards/pkg/services/auth/dto"
)

type Service interface {
	RegisterUser(ctx context.Context, email, password string) error
	LoginUser(ctx context.Context, email, password string) (*dto.LoginUserResult, error)
	LogoutUser(ctx context.Context, session string) error
	Refresh(ctx context.Context, refreshToken string) error
}

type service struct {
	store *maindb.Store
	maker maker.Maker
}

func New(store *maindb.Store, m maker.Maker) Service {
	return &service{
		store: store,
		maker: m,
	}
}
