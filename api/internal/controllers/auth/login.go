package auth

import (
	"encoding/json"
	"net/http"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/httputils"
	"github.com/4kord/english-flashcards/pkg/services/auth"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	UserID       int32  `json:"user_id"`
	Email        string `json:"email"`
	Admin        bool   `json:"admin"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
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

	userAgent := r.UserAgent()

	loginResult, err := c.AuthService.LoginUser(r.Context(), &auth.LoginUserParams{
		Email:     request.Email,
		Password:  request.Password,
		UserAgent: userAgent,
	})
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    loginResult.Session.RefreshToken,
		Path:     "/",
		HttpOnly: true,
		Expires:  loginResult.Session.ExpiresAt,
	})

	response := LoginUserResponse{
		UserID:       loginResult.User.ID,
		Email:        loginResult.User.Email,
		Admin:        loginResult.User.Admin,
		AccessToken:  loginResult.AccessToken,
		RefreshToken: loginResult.Session.RefreshToken,
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.Internal, errs.Code("send_response_failed")))
		return
	}
}
