package cards

import (
	"net/http"
	"strconv"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/httputils"
	"github.com/go-chi/chi/v5"
)

func (c *Controller) DeleteCard(w http.ResponseWriter, r *http.Request) {
	err := httputils.RequireContentType(r, "application/json")
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("unsupported_request_type")))
		return
	}

	cardID := chi.URLParam(r, "cardID")

	cardIDInt, err := strconv.ParseInt(cardID, 10, 32)

	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("invalid_parameter")))
		return
	}

	err = c.CardsService.DeleteCard(r.Context(), int32(cardIDInt))
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
