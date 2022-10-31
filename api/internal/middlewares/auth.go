package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maker"
	"go.uber.org/zap"
)

type UserIDKey struct{}
type IsAdminKey struct{}

type Auth struct {
	Maker     maker.Maker
	Log       *zap.Logger
	AdminOnly bool
}

func (a *Auth) Handler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if header == "" {
			errs.HTTPErrorResponse(w, a.Log, errs.E("authorization header must not be empty", errs.Unauthenticated, errs.Code("invalid_auth_header")))
			return
		}

		if len(strings.Split(header, " ")) != 2 {
			errs.HTTPErrorResponse(w, a.Log, errs.E("authorization header must contain two parts", errs.Unauthenticated, errs.Code("invalid_auth_header")))
			return
		}

		auth := strings.Split(header, " ")[0]

		if auth != "Bearer" {
			errs.HTTPErrorResponse(w, a.Log, errs.E("authorization header is not bearer", errs.Unauthenticated, errs.Code("invalid_auth_header")))
			return
		}

		token := strings.Split(header, " ")[1]

		payload, err := a.Maker.VerifyAccessToken(token)
		if err != nil {
			errs.HTTPErrorResponse(w, a.Log, errs.E(err, errs.Unauthenticated, errs.Code("invalid_auth_header")))
			return
		}

		if a.AdminOnly && !payload.Admin {
			errs.HTTPErrorResponse(w, a.Log, errs.E("Lack of admin privileges", errs.Unauthorized, errs.Code("access_denied")))
			return
		}

		ctx := r.Context()

		ctxWithPayload := context.WithValue(context.WithValue(ctx, UserIDKey{}, payload.Sub), IsAdminKey{}, payload.Admin)

		next.ServeHTTP(w, r.WithContext(ctxWithPayload))
	}

	return http.HandlerFunc(fn)
}

func NewAuth(m maker.Maker, log *zap.Logger, adminOnly bool) *Auth {
	return &Auth{
		Maker:     m,
		Log:       log,
		AdminOnly: adminOnly,
	}
}
