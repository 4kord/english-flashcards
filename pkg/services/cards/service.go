package cards

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/cld"
	"github.com/4kord/english-flashcards/pkg/maindb"
	"github.com/4kord/english-flashcards/pkg/services/cards/dto"
)

type Service interface {
	GetCards(ctx context.Context, deckID int32) ([]*maindb.Card, error)
	CreateCard(ctx context.Context, arg *dto.CreateCardParams) (*maindb.Card, error)
	EditCard(ctx context.Context, arg *dto.EditCardParams) (*maindb.Card, error)
	DeleteCard(ctx context.Context, cardID int32) error
	InsertCards(ctx context.Context, arg *dto.InsertCardsParams) error
}

type service struct {
	store *maindb.Store
	cld   cld.Cld
}

func New(store *maindb.Store, c cld.Cld) Service {
	return &service{
		store: store,
		cld:   c,
	}
}
