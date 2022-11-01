package cards

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/cld"
	"github.com/4kord/english-flashcards/pkg/maindb"
)

type Service interface {
	GetCards(ctx context.Context, deckID int32) ([]*maindb.Card, error)
	CreateCard(ctx context.Context, arg *CreateCardParams) (*maindb.Card, error)
	EditCard(ctx context.Context, arg *EditCardParams) (*maindb.Card, error)
	DeleteCard(ctx context.Context, cardID int32) error
	InsertCards(ctx context.Context, deckID int32, cardIds []int32) error
}

type service struct {
	store maindb.Store
	cld   cld.Cld
}

func New(store maindb.Store, c cld.Cld) Service {
	return &service{
		store: store,
		cld:   c,
	}
}
