package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/httputils"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	UserID    int32     `json:"userID"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Session   string    `json:"session"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	err := httputils.RequireContentType(r, "application/json")
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("unsupported_request_type")))
		return
	}

	var request LoginRequest

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, "decode_body_failed"))
		return
	}
	defer r.Body.Close()

	user, session, err := c.AuthService.LoginUser(r.Context(), request.Email, request.Password)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    session.Session,
		Expires:  session.ExpiresAt,
		Path:     "/",
		HttpOnly: true,
	})

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
