package auth

import (
	"context"
	"time"

	"github.com/4kord/english-flashcards/pkg/maindb"
	"github.com/4kord/english-flashcards/pkg/maker"
)

const (
	accessTokenExpiration  = time.Minute * 2
	refreshTokenExpiration = time.Hour * 24 * 30
)

type Service interface {
	RegisterUser(ctx context.Context, email, password string) error
	LoginUser(ctx context.Context, arg *LoginUserParams) (*LoginUserResult, error)
	LogoutUser(ctx context.Context, session string) error
	Refresh(ctx context.Context, arg *RefreshParams) (*RefreshResult, error)
	User(ctx context.Context, userID int32) (*maindb.User, error)
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
