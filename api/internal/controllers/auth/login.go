package auth

import (
	"encoding/json"
	"net/http"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/httputils"
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

	loginResult, err := c.AuthService.LoginUser(r.Context(), request.Email, request.Password)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, err)
	}

	response := LoginUserResponse{
		UserID:       loginResult.User.ID,
		Email:        loginResult.User.Email,
		Admin:        loginResult.User.Admin,
		AccessToken:  loginResult.AccessToken,
		RefreshToken: loginResult.RefreshToken,
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.Internal, errs.Code("send_response_failed")))
		return
	}
}
