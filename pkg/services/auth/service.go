package auth

import (
	"context"
	"time"

	"github.com/4kord/english-flashcards/pkg/maindb"
	"github.com/4kord/english-flashcards/pkg/maker"
)

const (
	accessTokenExpiration  = time.Minute * 15
	refreshTokenExpiration = time.Hour * 24 * 30
)

type Service interface {
	RegisterUser(ctx context.Context, email, password string) error
	LoginUser(ctx context.Context, email, password string) (*LoginUserResult, error)
	LogoutUser(ctx context.Context, session string) error
	Refresh(ctx context.Context, refreshToken string) (*RefreshResult, error)
}

type service struct {
	store maindb.Store
	maker maker.Maker
}

func New(store maindb.Store, m maker.Maker) Service {
	return &service{
		store: store,
		maker: m,
	}
}
