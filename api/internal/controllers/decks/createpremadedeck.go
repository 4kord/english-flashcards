package decks

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/httputils"
	"github.com/4kord/english-flashcards/pkg/services/decks/dto"
	"github.com/go-chi/chi/v5"
)

type createPremadeDeckRequest struct {
	Name string `json:"name"`
}

type createPremadeDeckResponse struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	Name      string    `json:"name"`
	Amount    int32     `json:"amount"`
	IsPremade bool      `json:"is_premade"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *Controller) CreatePremadeDeck(w http.ResponseWriter, r *http.Request) {
	err := httputils.RequireContentType(r, "application/json")
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("unsupported_request_type")))
		return
	}

	userID := chi.URLParam(r, "userID")

	userIDInt, err := strconv.ParseInt(userID, 10, 32)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("invalid_userID_parameter")))
		return
	}

	var request createDeckRequest

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.Internal, errs.Code("json_decode_failed")))
		return
	}

	deck, err := c.DecksService.CreateDeck(r.Context(), &dto.CreateDeckParams{
		UserID:    int32(userIDInt),
		Name:      request.Name,
		IsPremade: true,
	})

	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, err)
		return
	}

	response := createPremadeDeckResponse{
		ID:        deck.ID,
		UserID:    deck.UserID,
		Name:      deck.Name,
		Amount:    deck.Amount,
		IsPremade: deck.IsPremade,
		CreatedAt: deck.CreatedAt,
	}

	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.Internal, errs.Code("send_request_failed")))
		return
	}
}
