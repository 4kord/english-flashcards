package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/httputils"
)

func (c *Controller) User(w http.ResponseWriter, r *http.Request) {
	err := httputils.RequireContentType(r, "application/json")
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("unsupported_request_type")))
		return
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("no_session_cookie")))
			return
		}

		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.Internal, errs.Code("unable_to_get_session_cookie")))

		return
	}

	user, session, err := c.AuthService.User(r.Context(), cookie.Value)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, err)
		return
	}

	response := LoginUserResponse{
		UserID:    user.ID,
		Email:     user.Email,
		Role:      user.Role,
		Session:   session.Session,
		ExpiresAt: session.ExpiresAt,
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.Internal, errs.Code("send_response_failed")))
		return
	}
}
