package users

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/4kord/english-flashcards/pkg/errs"
)

type getUsersResponseEntity struct {
	ID        int32     `json:"id"`
	Email     string    `json:"email"`
	Admin     bool      `json:"admin"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *Controller) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.UsersService.GetUsers(r.Context())
	if err != nil {
		errs.HTTPErrorResponse(w, c.Log, err)
		return
	}

	var response = make([]getUsersResponseEntity, len(users))

	for i, v := range users {
		response[i] = getUsersResponseEntity{
			ID:        v.ID,
			Email:     v.Email,
			Admin:     v.Admin,
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
