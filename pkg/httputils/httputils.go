package httputils

import (
	"errors"
	"net/http"
)

func RequireContentType(r *http.Request, ct string) error {
	if r.Header.Get("Content-Type") != ct {
		return errors.New("unsupported content-type")
	}

	return nil
}
