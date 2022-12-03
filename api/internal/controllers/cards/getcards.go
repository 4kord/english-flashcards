package cards

import (
	"encoding/json"
	"fmt"
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
	Association   null.String `json:"association,omitempty"`
	Example       null.String `json:"example,omitempty"`
	Transcription null.String `json:"transcription,omitempty"`
	Image         null.String `json:"image,omitempty"`
	ImageURL      null.String `json:"image_url,omitempty"`
	Audio         null.String `json:"audio,omitempty"`
	AudioURL      null.String `json:"audio_url,omitempty"`
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
			Association:   cardsResult[i].Association,
			Example:       cardsResult[i].Example,
			Transcription: cardsResult[i].Transcription,
			Image:         cardsResult[i].Image,
			ImageURL:      cardsResult[i].ImageUrl,
			Audio:         cardsResult[i].Audio,
			AudioURL:      cardsResult[i].AudioUrl,
			CreatedAt:     cardsResult[i].CreatedAt,
		}
	}

	w.WriteHeader(http.StatusOK)

	b, _ := json.Marshal(response)
	fmt.Println(string(b))

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.Internal, errs.Code("send_request_failed")))
		return
	}
}
