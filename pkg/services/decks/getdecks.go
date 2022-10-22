package decks

import (
	"context"

	"github.com/4kord/english-flashcards/pkg/maindb"
)

func (s *service) GetDecks(ctx context.Context, userId int32) ([]*maindb.Deck, error) {
	return s.store.GetDecks(ctx, userId)
}
