package decks

import (
	"net/http"
	"strconv"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/httputils"
	"github.com/go-chi/chi/v5"
)

func (c *Controller) deleteDeck(w http.ResponseWriter, r *http.Request) {
	err := httputils.RequireContentType(r, "application/json")
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("unsupported_request_type")))
		return
	}

	deckID := chi.URLParam(r, "deckID")

	deckIDInt, err := strconv.ParseInt(deckID, 10, 32)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("invalid_deckID")))
		return
	}

	err = c.DecksService.DeleteDeck(r.Context(), int32(deckIDInt))
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
