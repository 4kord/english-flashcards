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
	"github.com/4kord/english-flashcards/pkg/httputils"
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

func (c *cardsController) GetCards(w http.ResponseWriter, r *http.Request) {
	err := httputils.RequireContentType(r, "application/json")
	if err != nil {
		errs.HTTPErrorResponse(w, c.log, errs.E(err, errs.InvalidRequest, errs.Code("unsupported_request_type")))
		return
	}

	var request getCardsRequest

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errs.HTTPErrorResponse(w, c.log, errs.E(err, errs.InvalidRequest, "decode_body_failed"))
		return
	}
	defer r.Body.Close()

	cardsResult, err := c.cardService.GetCards(r.Context(), request.DeckID)
	if err != nil {
		errs.HTTPErrorResponse(w, c.log, err)
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
	Association   null.String             `form:"association"`
	Example       null.String             `form:"example"`
	Transcription null.String             `form:"transcription"`
	Image         []*multipart.FileHeader `form:"image"`
	Audio         []*multipart.FileHeader `form:"audio"`
}

type createCardResponse struct {
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

func (c *cardsController) CreateCard(w http.ResponseWriter, r *http.Request) {
	err := httputils.RequireContentType(r, "multipart/form-data")
	if err != nil {
		errs.HTTPErrorResponse(w, c.log, errs.E(err, errs.InvalidRequest, errs.Code("unsupported_request_type")))
		return
	}

	var request createCardRequest

	deckID := chi.URLParam(r, "deckID")

	deckIDInt, err := strconv.ParseInt(deckID, 10, 32)

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

	card, err := c.cardService.CreateCard(r.Context(), &maindb.Card{
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
		Association:   null.String(card.Association),
		Example:       null.String(card.Example),
		Transcription: null.String(card.Transcription),
		Image:         null.String(card.Image),
		ImageURL:      null.String(card.ImageUrl),
		Audio:         null.String(card.Audio),
		AudioURL:      null.String(card.AudioUrl),
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
