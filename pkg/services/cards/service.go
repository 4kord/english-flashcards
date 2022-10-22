package cards

import (
	"context"
	"mime/multipart"

	"github.com/4kord/english-flashcards/pkg/cld"
	"github.com/4kord/english-flashcards/pkg/maindb"
)

type Service interface {
	GetCards(ctx context.Context, deckId int32) ([]*maindb.Card, error)
	CreateCard(ctx context.Context, card maindb.Card, image *multipart.FileHeader, audio *multipart.FileHeader) (*maindb.Card, error)
	// EditCard(ctx context.Context, card maindb.Card) (*maindb.Card, error)
	// DeleteCard(ctx context.Context, cardId int32) error
	// CopyCards(ctx context.Context, cardIds []int32) error
}

type service struct {
	store *maindb.Store
	cld   *cld.Cld
}

func New(store *maindb.Store, cld *cld.Cld) Service {
	return &service{
		store: store,
		cld:   cld,
	}
}
