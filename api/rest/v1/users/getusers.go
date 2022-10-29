package users

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/httputils"
)

type getUsersResponseEntity struct {
	ID        int32     `json:"id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *Controller) GetUsers(w http.ResponseWriter, r *http.Request) {
	err := httputils.RequireContentType(r, "application/json")
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.InvalidRequest, errs.Code("unsupported_request_type")))
		return
	}

	users, err := c.UserService.GetUsers(r.Context())
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, err)
		return
	}

	var response = make([]getUsersResponseEntity, len(users))

	for i, v := range users {
		response[i] = getUsersResponseEntity{
			ID:        v.ID,
			Email:     v.Email,
			Role:      v.Role,
			CreatedAt: v.CreatedAt,
		}
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, errs.E(err, errs.Internal, errs.Code("send_request_failed")))
		return
	}
}
