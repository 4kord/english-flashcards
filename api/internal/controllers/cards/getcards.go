package cards

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/null"
	"github.com/go-chi/chi/v5"
)

type getCardsResponseEntity struct {
	ID            int32       `json:"id"`
	DeckID        int32       `json:"deck_id"`
	English       string      `json:"english"`
	Russian       string      `json:"russian"`
	Association   null.String `json:"association"`
	Example       null.String `json:"example"`
	Transcription null.String `json:"transcription"`
	Image         null.String `json:"image"`
	ImageURL      null.String `json:"image_url"`
	Audio         null.String `json:"audio"`
	AudioURL      null.String `json:"audio_url"`
	CreatedAt     time.Time   `json:"created_at"`
}

func (c *Controller) GetCards(w http.ResponseWriter, r *http.Request) {
	deckID := chi.URLParam(r, "deckID")

	deckIDInt, err := strconv.ParseInt(deckID, 10, 32)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("invalid_parameter")))
		return
	}

	cardsResult, err := c.CardsService.GetCards(r.Context(), int32(deckIDInt))
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, err)
		return
	}

	response := make([]*getCardsResponseEntity, len(cardsResult))
	for i := 0; i < len(response); i++ {
		response[i] = &getCardsResponseEntity{
			ID:            cardsResult[i].ID,
			DeckID:        cardsResult[i].DeckID,
			English:       cardsResult[i].English,
			Russian:       cardsResult[i].Russian,
			Association:   null.String(cardsResult[i].Association),
			Example:       null.String(cardsResult[i].Example),
			Transcription: null.String(cardsResult[i].Transcription),
			Image:         null.String(cardsResult[i].Image),
			ImageURL:      null.String(cardsResult[i].ImageUrl),
			Audio:         null.String(cardsResult[i].Audio),
			AudioURL:      null.String(cardsResult[i].AudioUrl),
			CreatedAt:     cardsResult[i].CreatedAt,
		}
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.Internal, errs.Code("send_request_failed")))
		return
	}
}
