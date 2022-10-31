package decks

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/httputils"
)

type getPremadeDecksResponseEntity struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	Name      string    `json:"name"`
	Amount    int32     `json:"amount"`
	IsPremade bool      `json:"is_premade"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *Controller) GetPremadeDecks(w http.ResponseWriter, r *http.Request) {
	err := httputils.RequireContentType(r, "application/json")
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("unsupported_request_type")))
		return
	}

	decks, err := c.DecksService.GetPremadeDecks(r.Context())
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, err)
		return
	}

	var response = make([]getPremadeDecksResponseEntity, len(decks))

	for i, v := range decks {
		response[i] = getPremadeDecksResponseEntity{
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
