package v1

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/httputils"
	"github.com/4kord/english-flashcards/pkg/services/auth"
	"go.uber.org/zap"
)

type authController struct {
	authService auth.Service
	log         *zap.Logger
}

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerResponse struct {
	ID        int32     `json:"id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *authController) Register(w http.ResponseWriter, r *http.Request) {
	err := httputils.RequireContentType(r, "application/json")
	if err != nil {
		errs.HTTPErrorResponse(w, c.log, errs.E(err, errs.InvalidRequest, errs.Code("unsupported_request_type")))
		return
	}

	var request registerRequest

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errs.HTTPErrorResponse(w, c.log, errs.E(err, errs.InvalidRequest, "decode_body_failed"))
		return
	}
	defer r.Body.Close()

	created, err := c.authService.RegisterUser(r.Context(), request.Email, request.Password)
	if err != nil {
		errs.HTTPErrorResponse(w, c.log, err)
		return
	}

	response := registerResponse{
		ID:        created.ID,
		Email:     created.Email,
		Role:      created.Role,
		CreatedAt: created.CreatedAt,
	}

	b, err := json.Marshal(response)
	if err != nil {
		errs.HTTPErrorResponse(w, c.log, errs.E(err, errs.Internal, errs.Code("user_marshal_failed")))
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(b)

	if err != nil {
		errs.HTTPErrorResponse(w, c.log, errs.E(err, errs.Internal, errs.Code("sending_request_failed")))
		return
	}
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	UserID    int32     `json:"userID"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Session   string    `json:"session"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (c *authController) Login(w http.ResponseWriter, r *http.Request) {
	err := httputils.RequireContentType(r, "application/json")
	if err != nil {
		errs.HTTPErrorResponse(w, c.log, errs.E(err, errs.InvalidRequest, errs.Code("unsupported_request_type")))
		return
	}

	var request loginRequest

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errs.HTTPErrorResponse(w, c.log, errs.E(err, errs.InvalidRequest, "decode_body_failed"))
		return
	}
	defer r.Body.Close()

	user, session, err := c.authService.LoginUser(r.Context(), request.Email, request.Password)
	if err != nil {
		errs.HTTPErrorResponse(w, c.log, err)
	}

	response := loginResponse{
		UserID:    user.ID,
		Email:     user.Email,
		Role:      user.Role,
		Session:   session.Session,
		ExpiresAt: session.ExpiresAt,
	}

	b, err := json.Marshal(response)
	if err != nil {
		errs.HTTPErrorResponse(w, c.log, errs.E(err, errs.Internal, errs.Code("user_marshal_failed")))
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(b)

	if err != nil {
		errs.HTTPErrorResponse(w, c.log, errs.E(err, errs.Internal, errs.Code("sending_request_failed")))
		return
	}
}
