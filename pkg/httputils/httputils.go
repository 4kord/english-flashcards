package httputils

import (
	"errors"
	"net/http"
	"strings"
)

func RequireContentType(r *http.Request, ct string) error {
	if !strings.Contains(r.Header.Get("Content-Type"), ct) {
		return errors.New("unsupported content-type")
	}

	return nil
}
