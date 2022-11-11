package auth

import (
	"encoding/json"
	"net/http"

	"github.com/4kord/english-flashcards/api/internal/middlewares"
	"github.com/4kord/english-flashcards/pkg/errs"
)

type UserResponse struct {
	UserID int32  `json:"user_id"`
	Email  string `json:"email"`
	Admin  bool   `json:"admin"`
}

func (c *Controller) User(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value(middlewares.UserIDKey{}).(int32)
	if !ok {
		errs.HTTPErrorResponse(w, c.Log, errs.E("Lack Of User ID Value", errs.InvalidRequest, errs.Code("lack_of_user_id")))
		return
	}

	user, err := c.AuthService.User(r.Context(), id)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, err)
		return
	}

	response := UserResponse{
		UserID: user.ID,
		Email:  user.Email,
		Admin:  user.Admin,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.Internal, errs.Code("send_request_failed")))
		return
	}
}
