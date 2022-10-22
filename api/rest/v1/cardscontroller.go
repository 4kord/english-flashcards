package v1

import (
	"database/sql"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/formdata"
	"github.com/4kord/english-flashcards/pkg/maindb"
	"github.com/4kord/english-flashcards/pkg/null"
	"github.com/4kord/english-flashcards/pkg/services/cards"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type cardsController struct {
	cardService cards.Service
	log         *zap.Logger
}

type getCardsRequest struct {
	DeckID int32 `json:"deck_id"`
}

type getCardsResponseEntity struct {
	ID            int32           `json:"id"`
	DeckID        int32           `json:"deck_id"`
	English       string          `json:"english"`
	Russian       string          `json:"russian"`
	Association   null.NullString `json:"association"`
	Example       null.NullString `json:"example"`
	Transcription null.NullString `json:"transcription"`
	Image         null.NullString `json:"image"`
	ImageUrl      null.NullString `json:"image_url"`
	Audio         null.NullString `json:"audio"`
	AudioUrl      null.NullString `json:"audio_url"`
	CreatedAt     time.Time       `json:"created_at"`
}

func (c *cardsController) GetCards(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		errs.HTTPErrorResponse(w, c.log, errs.E("Unsupported request type", errs.InvalidRequest, errs.Code("unsupported_request_type")))
		return
	}

	var request getCardsRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errs.HTTPErrorResponse(w, c.log, errs.E(err, errs.InvalidRequest, "decode_body_failed"))
		return
	}
	defer r.Body.Close()

	cards, err := c.cardService.GetCards(r.Context(), request.DeckID)
	if err != nil {
		errs.HTTPErrorResponse(w, c.log, err)
		return
	}

	response := make([]*getCardsResponseEntity, len(cards))
	for i := 0; i < len(response); i++ {
		response[i] = &getCardsResponseEntity{
			ID:            cards[i].ID,
			DeckID:        cards[i].DeckID,
			English:       cards[i].English,
			Russian:       cards[i].Russian,
			Association:   null.NullString(cards[i].Association),
			Example:       null.NullString(cards[i].Example),
			Transcription: null.NullString(cards[i].Transcription),
			Image:         null.NullString(cards[i].Image),
			ImageUrl:      null.NullString(cards[i].ImageUrl),
			Audio:         null.NullString(cards[i].Audio),
			AudioUrl:      null.NullString(cards[i].AudioUrl),
			CreatedAt:     cards[i].CreatedAt,
		}
	}

	b, err := json.Marshal(response)
	if err != nil {
		errs.HTTPErrorResponse(w, c.log, errs.E(err, errs.Internal, errs.Code("cards_marshal_failed")))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		errs.HTTPErrorResponse(w, c.log, errs.E(err, errs.Internal, errs.Code("sending_request_failed")))
		return
	}
}

type createCardRequest struct {
	English       string                  `form:"english"`
	Russian       string                  `form:"russian"`
	Association   null.NullString         `form:"association"`
	Example       null.NullString         `form:"example"`
	Transcription null.NullString         `form:"transcription"`
	Image         []*multipart.FileHeader `form:"image"`
	Audio         []*multipart.FileHeader `form:"audio"`
}

type createCardResponse struct {
	ID            int32           `json:"id"`
	DeckID        int32           `json:"deck_id"`
	English       string          `json:"english"`
	Russian       string          `json:"russian"`
	Association   null.NullString `json:"association"`
	Example       null.NullString `json:"example"`
	Transcription null.NullString `json:"transcription"`
	Image         null.NullString `json:"image"`
	ImageUrl      null.NullString `json:"image_url"`
	Audio         null.NullString `json:"audio"`
	AudioUrl      null.NullString `json:"audio_url"`
	CreatedAt     time.Time       `json:"created_at"`
}

func (c *cardsController) CreateCard(w http.ResponseWriter, r *http.Request) {
	var request createCardRequest

	deckID := chi.URLParam(r, "deckID")

	deckIDInt, err := strconv.Atoi(deckID)
	if err != nil {
		errs.HTTPErrorResponse(w, c.log, errs.E(err, errs.InvalidRequest, errs.Code("bad_url_parameter")))
		return
	}

	err = formdata.Decode(r, &request)
	if err != nil {
		errs.HTTPErrorResponse(w, c.log, errs.E(err, errs.Internal, errs.Code("formdata_decode_failed")))
		return
	}

	var image *multipart.FileHeader
	if len(request.Image) > 0 {
		image = request.Image[0]
	}

	var audio *multipart.FileHeader
	if len(request.Audio) > 0 {
		image = request.Audio[0]
	}

	card, err := c.cardService.CreateCard(r.Context(), maindb.Card{
		DeckID:        int32(deckIDInt),
		English:       request.English,
		Russian:       request.Russian,
		Association:   sql.NullString(request.Association),
		Example:       sql.NullString(request.Example),
		Transcription: sql.NullString(request.Transcription),
	}, image, audio)

	if err != nil {
		errs.HTTPErrorResponse(w, c.log, err)
		return
	}

	response := createCardResponse{
		ID:            card.ID,
		DeckID:        card.DeckID,
		English:       card.English,
		Russian:       card.Russian,
		Association:   null.NullString(card.Association),
		Example:       null.NullString(card.Example),
		Transcription: null.NullString(card.Transcription),
		Image:         null.NullString(card.Image),
		ImageUrl:      null.NullString(card.ImageUrl),
		Audio:         null.NullString(card.Audio),
		AudioUrl:      null.NullString(card.AudioUrl),
		CreatedAt:     card.CreatedAt,
	}

	b, err := json.Marshal(response)
	if err != nil {
		errs.HTTPErrorResponse(w, c.log, errs.E(err, errs.Internal, errs.Code("card_marshal_failed")))
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(b)
	if err != nil {
		errs.HTTPErrorResponse(w, c.log, errs.E(err, errs.Internal, errs.Code("sending_request_failed")))
		return
	}
}
