package decks

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/httputils"
	"github.com/4kord/english-flashcards/pkg/services/decks"
	"github.com/go-chi/chi/v5"
)

type editDeckRequest struct {
	Name string `json:"name"`
}

type editDeckResponse struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	Name      string    `json:"name"`
	Amount    int32     `json:"amount"`
	IsPremade bool      `json:"is_premade"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *Controller) editDeck(w http.ResponseWriter, r *http.Request) {
	err := httputils.RequireContentType(r, "application/json")
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("unsupported_request_type")))
		return
	}

	deckID := chi.URLParam(r, "deckID")

	deckIDInt, err := strconv.ParseInt(deckID, 10, 32)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("invalid_userID_parameter")))
		return
	}

	var request editDeckRequest

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.Internal, errs.Code("json_decode_failed")))
		return
	}

	deck, err := c.DecksService.EditDeck(r.Context(), &decks.EditDeckParams{
		DeckID:  int32(deckIDInt),
		NewName: request.Name,
	})

	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, err)
		return
	}

	response := editDeckResponse{
		ID:        deck.ID,
		UserID:    deck.UserID,
		Name:      deck.Name,
		Amount:    deck.Amount,
		IsPremade: deck.IsPremade,
		CreatedAt: deck.CreatedAt,
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.Internal, errs.Code("send_request_failed")))
		return
	}
}
