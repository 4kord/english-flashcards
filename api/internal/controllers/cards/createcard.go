package cards

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/formdata"
	"github.com/4kord/english-flashcards/pkg/httputils"
	"github.com/4kord/english-flashcards/pkg/null"
	"github.com/4kord/english-flashcards/pkg/services/cards"
	"github.com/go-chi/chi/v5"
	"github.com/guregu/null/zero"
)

type test struct {
	Test1 zero.String `json:"test1,omitempty"`
}

type createCardRequest struct {
	English       string                  `form:"english"`
	Russian       string                  `form:"russian"`
	Association   null.String             `form:"association"`
	Example       null.String             `form:"example"`
	Transcription null.String             `form:"transcription"`
	Image         []*multipart.FileHeader `form:"image"`
	ImageURL      null.String             `form:"image_url"`
	Audio         []*multipart.FileHeader `form:"audio"`
	AudioURL      null.String             `form:"audio_url"`
}

type createCardResponse struct {
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

func (c *Controller) CreateCard(w http.ResponseWriter, r *http.Request) {
	err := httputils.RequireContentType(r, "multipart/form-data")
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("unsupported_request_type")))
		return
	}

	deckID := chi.URLParam(r, "deckID")

	deckIDInt, err := strconv.ParseInt(deckID, 10, 32)

	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("invalid_parameter")))
		return
	}

	var request createCardRequest

	err = formdata.Decode(r, &request)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.Internal, errs.Code("formdata_decode_failed")))
		return
	}

	var image *multipart.FileHeader

	if len(request.Image) > 0 {
		image = request.Image[0]
	}

	var audio *multipart.FileHeader

	if len(request.Audio) > 0 {
		audio = request.Audio[0]
	}

	card, err := c.CardsService.CreateCard(r.Context(), &cards.CreateCardParams{
		DeckID:        int32(deckIDInt),
		English:       request.English,
		Russian:       request.Russian,
		Association:   request.Association,
		Example:       request.Example,
		Transcription: request.Transcription,
		Image:         image,
		ImageURL:      request.ImageURL,
		Audio:         audio,
		AudioURL:      request.AudioURL,
	})

	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, err)
		return
	}

	response := createCardResponse{
		ID:            card.ID,
		DeckID:        card.DeckID,
		English:       card.English,
		Russian:       card.Russian,
		Association:   card.Association,
		Example:       card.Example,
		Transcription: card.Transcription,
		Image:         card.Image,
		ImageURL:      card.ImageUrl,
		Audio:         card.Audio,
		AudioURL:      card.AudioUrl,
		CreatedAt:     card.CreatedAt,
	}

	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.Internal, errs.Code("send_request_failed")))
		return
	}
}
