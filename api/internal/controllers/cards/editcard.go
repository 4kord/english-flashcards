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
)

type editCardRequest struct {
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

type editCardResponse struct {
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

func (c *Controller) EditCard(w http.ResponseWriter, r *http.Request) {
	err := httputils.RequireContentType(r, "multipart/form-data")
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("unsupported_request_type")))
		return
	}

	var request editCardRequest

	cardID := chi.URLParam(r, "cardID")

	cardIDInt, err := strconv.ParseInt(cardID, 10, 32)

	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("invalid_parameter")))
		return
	}

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
		image = request.Audio[0]
	}

	card, err := c.CardsService.EditCard(r.Context(), &cards.EditCardParams{
		CardID:        int32(cardIDInt),
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
	}

	response := editCardResponse{
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

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.Internal, errs.Code("send_request_failed")))
		return
	}
}
