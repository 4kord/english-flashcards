package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/services/auth"
	"go.uber.org/zap"
)

type Auth struct {
	AuthService  auth.Service
	Log          *zap.Logger
	AllowedRoles []string
}

func (a *Auth) Handler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				errs.HTTPErrorResponse(w, a.Log, errs.E(err, errs.InvalidRequest, errs.Code("no_session_cookie")))
				return
			}

			errs.HTTPErrorResponse(w, a.Log, errs.E(err, errs.Internal, errs.Code("unable_to_get_session_cookie")))

			return
		}

		user, session, err := a.AuthService.User(r.Context(), cookie.Value)
		if err != nil {
			errs.HTTPErrorResponse(w, a.Log, err)
			return
		}

		contextWithUserID := context.WithValue(r.Context(), "userID", user.ID)

		contextWithRole := context.WithValue(contextWithUserID, "role", user.Role)

		request := r.WithContext(contextWithRole)

		// if allowed roles is empty, allow any role to access endpoint
		if len(a.AllowedRoles) == 0 {
			next.ServeHTTP(w, request)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    session.Session,
			Expires:  session.ExpiresAt,
			Path:     "/",
			HttpOnly: true,
		})
	}

	return http.HandlerFunc(fn)
}

func New(authService auth.Service, log *zap.Logger, allowedRoles []string) *Auth {
	return &Auth{
		AuthService:  authService,
		Log:          log,
		AllowedRoles: allowedRoles,
	}
}
