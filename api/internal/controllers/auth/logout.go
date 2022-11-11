package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/httputils"
)

func (c *Controller) Logout(w http.ResponseWriter, r *http.Request) {
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

	err = c.AuthService.LogoutUser(r.Context(), cookie.Value)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(-100 * time.Hour),
	})

	w.WriteHeader(http.StatusOK)
}
