package decks

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/maindb"
)

type Service interface {
	GetDecks(ctx context.Context, userID int32) ([]*maindb.Deck, error)
	GetPremadeDecks(ctx context.Context) ([]*maindb.Deck, error)
	CreateDeck(ctx context.Context, deck *maindb.Deck) (*maindb.Deck, error)
	EditDeck(ctx context.Context, deck *maindb.Deck) (*maindb.Deck, error)
	DeleteDeck(ctx context.Context, deckID int32) error
}

type service struct {
	store *maindb.Store
}

func New(store *maindb.Store) Service {
	return &service{
		store: store,
	}
}
