package cards

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/httputils"
	"github.com/4kord/english-flashcards/pkg/services/cards/dto"
	"github.com/go-chi/chi/v5"
)

type insertCardsRequest struct {
	CardIDs []int32 `json:"card_ids"`
}

func (c *Controller) InsertCards(w http.ResponseWriter, r *http.Request) {
	err := httputils.RequireContentType(r, "application/json")
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("unsupported_request_type")))
		return
	}

	var request insertCardsRequest

	deckID := chi.URLParam(r, "deckID")

	deckIDInt, err := strconv.ParseInt(deckID, 10, 32)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("invalid_parameter")))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, "decode_body_failed"))
		return
	}
	defer r.Body.Close()

	err = c.CardsService.InsertCards(r.Context(), &dto.InsertCardsParams{
		DeckID:  int32(deckIDInt),
		CardIDs: request.CardIDs,
	})

	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
