package middlewares

import (
	"net/http"
	"strconv"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Require struct {
	Log *zap.Logger
}

func NewRequire(log *zap.Logger) *Require {
	return &Require{
		Log: log,
	}
}

func (m *Require) SameUserID(paramName string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, err := strconv.ParseInt(chi.URLParam(r, paramName), 10, 32)
			if err != nil {
				errs.HTTPErrorResponse(w, m.Log, errs.E(err, errs.Internal, errs.Code("parse_userID_failed")))
				return
			}

			userIDCtx, ok := r.Context().Value(UserIDKey{}).(int32)
			if !ok {
				errs.HTTPErrorResponse(w, m.Log, errs.E("User ID context is not int32", errs.Internal, errs.Code("userIDCtx_not_int32")))
				return
			}

			if userID != int64(userIDCtx) {
				errs.HTTPErrorResponse(w, m.Log, errs.E("You don't have access to this resource", errs.Unauthorized, errs.Code("lack_of_access")))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
