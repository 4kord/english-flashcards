package google

import (
	"encoding/json"
	"net/http"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/go-chi/chi/v5"
)

type fetchAudioResponse struct {
	URL string `json:"url"`
}

func (c *Controller) FetchAudio(w http.ResponseWriter, r *http.Request) {
	word := chi.URLParam(r, "word")

	url, err := c.GoogleService.FetchAudio(r.Context(), word)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, err)
		return
	}

	var response fetchAudioResponse

	response.URL = url

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.Internal, errs.Code("send_request_failed")))
		return
	}
}
