package decks

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/httputils"
	"github.com/go-chi/chi/v5"
)

type getDecksResponseEntity struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	Name      string    `json:"name"`
	Amount    int32     `json:"amount"`
	IsPremade bool      `json:"is_premade"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *Controller) GetDecks(w http.ResponseWriter, r *http.Request) {
	err := httputils.RequireContentType(r, "application/json")
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("unsupported_request_type")))
		return
	}

	deckID := chi.URLParam(r, "userID")

	deckIDInt, err := strconv.ParseInt(deckID, 10, 32)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("invalid_userID_parameter")))
		return
	}

	decks, err := c.DeckService.GetDecks(r.Context(), int32(deckIDInt))
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, err)
		return
	}

	var response = make([]getDecksResponseEntity, len(decks))

	for i, v := range decks {
		response[i] = getDecksResponseEntity{
			ID:        v.ID,
			UserID:    v.UserID,
			Name:      v.Name,
			Amount:    v.Amount,
			IsPremade: v.IsPremade,
			CreatedAt: v.CreatedAt,
		}
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.Internal, errs.Code("send_request_failed")))
		return
	}
}
