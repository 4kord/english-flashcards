package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/httputils"
)

type RefreshResponse struct {
	UserID       int32     `json:"user_id"`
	Email        string    `json:"email"`
	Admin        bool      `json:"admin"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func (c *Controller) Refresh(w http.ResponseWriter, r *http.Request) {
	err := httputils.RequireContentType(r, "application/json")
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("unsupported_request_type")))
		return
	}

	sessionCookie, err := r.Cookie("session")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("lack_of_session_cookie")))
			return
		}

		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("error_getting_cookie")))

		return
	}

	result, err := c.AuthService.Refresh(r.Context(), sessionCookie.Value)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    result.Session.RefreshToken,
		Path:     "/",
		HttpOnly: true,
		Expires:  result.Session.ExpiresAt,
	})

	response := RefreshResponse{
		UserID:       result.User.ID,
		Email:        result.User.Email,
		Admin:        result.User.Admin,
		AccessToken:  result.AccessToken,
		RefreshToken: result.Session.RefreshToken,
		ExpiresAt:    result.Session.ExpiresAt,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.Internal, "send_request_failed"))
		return
	}
}
