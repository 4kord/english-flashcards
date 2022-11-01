package users

import (
	"net/http"
	"strconv"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/httputils"
	"github.com/go-chi/chi/v5"
)

func (c *Controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	err := httputils.RequireContentType(r, "application/json")
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("unsupported_request_type")))
		return
	}

	userID := chi.URLParam(r, "userID")

	userIDInt, err := strconv.ParseInt(userID, 10, 32)

	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("invalid_parameter")))
		return
	}

	err = c.UsersService.DeleteUser(r.Context(), int32(userIDInt))
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
